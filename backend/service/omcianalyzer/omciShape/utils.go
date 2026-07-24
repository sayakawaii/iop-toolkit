package omciShape

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type OmciState int

type OmciMetaData struct {
	Timestamp time.Time `json:"Timestamp"`
	RawData   string    `json:"RawData"`
}

func flushOmciDataToFile(onus map[string]string, omciMetaDataMap map[string][]OmciMetaData) {
	for onu, metaDataList := range omciMetaDataMap {
		omciMetaDataList := metaDataList
		jsonData, err := json.MarshalIndent(omciMetaDataList, "", "  ")
		if err != nil {
			utils.Log("json marshal failed: " + err.Error())
			continue
		}
		models.WriteFile(onus[onu], string(jsonData))
	}
}

func isLikelyValidXML(data []byte) bool {
	open := bytes.Count(data, []byte("<"))
	close := bytes.Count(data, []byte(">"))
	return open == close && bytes.HasSuffix(bytes.TrimSpace(data), []byte(">"))
}

// SanitizeXML cleans illegal XML characters in attributes and text content.
// It escapes &, <, >, ", and ' where appropriate, without breaking tag structure.
func sanitizeXML(data []byte) []byte {
	s := string(data)

	// --- Step 0: 移除非法字符引用（例如 &#x0; &#xB; &#x1F;）---
	// XML 1.0 不允许 0x0–0x8、0xB、0xC、0xE–0x1F
	reIllegalEntity := regexp.MustCompile(`&#x(?:0?[0-8bBcCeEfF]|1[0-9a-fA-F]);`)
	s = reIllegalEntity.ReplaceAllString(s, ".")

	// --- Step 1: 先处理属性值内的非法字符 ---
	// 匹配形如 key="value"
	reAttr := regexp.MustCompile(`="([^"]*)"`)
	s = reAttr.ReplaceAllStringFunc(s, func(attr string) string {
		// 提取属性值
		val := attr[2 : len(attr)-1]
		val = xmlEscape(val)
		return `="` + val + `"`
	})

	// --- Step 2: 再处理标签外（纯文本节点）里的非法字符 ---
	// 匹配标签外部内容
	reText := regexp.MustCompile(`>([^<]+)<`)
	s = reText.ReplaceAllStringFunc(s, func(txt string) string {
		content := txt[1 : len(txt)-1]
		content = xmlEscape(content)
		return ">" + content + "<"
	})

	return []byte(s)
}

// xmlEscape replaces the 5 reserved XML chars with entities.
func xmlEscape(val string) string {
	var buf bytes.Buffer
	for _, r := range val {
		switch r {
		case '&':
			buf.WriteString("&amp;")
		case '<':
			buf.WriteString("&lt;")
		case '>':
			buf.WriteString("&gt;")
		case '"':
			buf.WriteString("&quot;")
		case '\'':
			buf.WriteString("&apos;")
		default:
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

func xmlDataClean(data []byte) []byte {
	data = bytes.ToValidUTF8(data, []byte{})
	data = sanitizeXML(data)
	data = bytes.ToValidUTF8(data, []byte{})
	return data
}

func decodeSerialPartialVisible(fieldValue string) string {
	fieldValue = strings.TrimSpace(fieldValue)
	if !strings.Contains(fieldValue, "SerialNumber") {
		return fieldValue
	}

	var vendor, serial string
	if idx := strings.Index(fieldValue, "VendorId:"); idx >= 0 {
		part := fieldValue[idx+len("VendorId:"):]
		part = strings.TrimSpace(part)
		if cidx := strings.Index(part, ","); cidx > 0 {
			part = part[:cidx]
		}
		vendor = strings.Trim(part, " {}")
	}
	if idx := strings.Index(fieldValue, "SerialNumber:"); idx >= 0 {
		part := fieldValue[idx+len("SerialNumber:"):]
		part = strings.TrimSpace(part)
		part = strings.Trim(part, "{} ")
		re := regexp.MustCompile(`0x([0-9a-fA-F]{2})`)
		matches := re.FindAllStringSubmatch(part, -1)
		for _, m := range matches {
			serial += m[1]
		}
	}

	if serial == "" {
		return vendor
	}
	return vendor + strings.ToUpper(serial)
}

func sstWriteFile(fileName string, line string, direction string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		utils.Log("open file failed")
	} else {
		// utils.Log("open file success")
		defer file.Close()
		file.WriteString("[")
		file.WriteString(strconv.FormatInt(time.Now().Unix(), 10) + "." + strconv.FormatUint(rand.Uint64(), 10)[0:10] + ",")
		file.WriteString("\"" + direction + "\"" + ",")
		file.WriteString("\"" + line + "\"")
		file.WriteString("],\n")
	}
}

func locateValidOmciData(s string, startIndex int) int {
	//locate valide OMCI data
	for k, c := range s[startIndex:] {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			return startIndex + k
		}
	}
	return startIndex
}

func extractSegment(strIn, marker string, length int) (string, bool) {
	s := strings.TrimSpace(strIn)
	idx := strings.LastIndex(s, marker)
	if idx < 0 {
		return "", false
	}

	idx += len(marker)
	idx = locateValidOmciData(s, idx)

	if idx+length > len(s) {
		return "", false
	}

	return models.StringStrip(s[idx : idx+length]), true
}
