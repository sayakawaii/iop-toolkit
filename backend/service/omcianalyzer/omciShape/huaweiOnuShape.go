/*
# ------------------------------------------------------------
# -- huaweiOnuShape.go
# --
# -- Huang Minghe
# -- 2024-4-3
# ------------------------------------------------------------
*/

package omciShape

import (
	"encoding/hex"
	"fmt"
	"io"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type huaweiOnuShaper struct {
	cachedRegex     *regexp.Regexp
	cachedLayout    string
	outputPath      string
	state           OmciState
	onus            map[string]string
	ontid           string
	omciBuff        string
	timeStamp       time.Time
	omciMetaDataMap map[string][]OmciMetaData
	deviceSN        string
}

const (
	HuaweiLineInit OmciState = iota
	HuaweiLineData0
	HuaweiLineData1
	HuaweiLineData2
	HuaweiLineData3
)

func (shaper *huaweiOnuShaper) getTimestamp(s string) time.Time {
	// example: OLT--------->ONT: 1970/1/1 0:6:42:562995
	// example: ONT--------->OLT: 1970/1/1 0:6:42:745773
	// Format: YYYY/M/D H:M:S:UUUUUU (where U is microseconds)

	re := regexp.MustCompile(`\d{4}/\d{1,2}/\d{1,2}\s+\d{1,2}:\d{1,2}:\d{1,2}:\d{6}`)
	match := re.FindString(s)
	if match == "" {
		return time.Time{}
	}

	// Split by space to get date and time parts
	parts := strings.Split(match, " ")
	if len(parts) != 2 {
		return time.Time{}
	}

	datePart := parts[0] // "1970/1/1"
	timePart := parts[1] // "0:6:42:562995"

	// Parse date
	dateFields := strings.Split(datePart, "/")
	if len(dateFields) != 3 {
		return time.Time{}
	}

	year, err1 := strconv.Atoi(dateFields[0])
	month, err2 := strconv.Atoi(dateFields[1])
	day, err3 := strconv.Atoi(dateFields[2])
	if err1 != nil || err2 != nil || err3 != nil {
		return time.Time{}
	}

	// Parse time
	timeFields := strings.Split(timePart, ":")
	if len(timeFields) != 4 {
		return time.Time{}
	}

	hour, err4 := strconv.Atoi(timeFields[0])
	minute, err5 := strconv.Atoi(timeFields[1])
	second, err6 := strconv.Atoi(timeFields[2])
	microsecond, err7 := strconv.Atoi(timeFields[3])
	if err4 != nil || err5 != nil || err6 != nil || err7 != nil {
		return time.Time{}
	}

	t := time.Date(year, time.Month(month), day, hour, minute, second, microsecond*1000, time.UTC)
	return t
}

func (shaper *huaweiOnuShaper) extractSN(s string) string {
	// example: SN.:48575443512a1e15
	if !strings.Contains(s, "SN.:") && !strings.Contains(s, "SN:") {
		return ""
	}

	snPrefix := "SN.:"
	if !strings.Contains(s, snPrefix) {
		snPrefix = "SN:"
	}

	idx := strings.Index(s, snPrefix)
	if idx == -1 {
		return ""
	}

	snValue := strings.TrimSpace(s[idx+len(snPrefix):])
	// Remove any trailing characters after SN value
	if spaceIdx := strings.Index(snValue, " "); spaceIdx != -1 {
		snValue = snValue[:spaceIdx]
	}

	return snValue
}

func (shaper *huaweiOnuShaper) snToOnuID(sn string) string {
	// SN format: 48575443512a1e15
	// Convert to: HWTC + readable ASCII or hex
	if len(sn) < 8 {
		return "HuaweiOnu"
	}

	// Decode hex string to bytes
	snBytes, err := hex.DecodeString(sn)
	if err != nil {
		return "HuaweiOnu"
	}

	if len(snBytes) < 4 {
		return "HuaweiOnu"
	}

	// First 4 bytes are vendor ID (e.g., "HWTC")
	vendorID := string(snBytes[0:4])

	// Remaining bytes: convert to ASCII if printable, otherwise use hex
	var result strings.Builder
	result.WriteString(vendorID)

	for _, b := range snBytes[4:] {
		if unicode.IsPrint(rune(b)) && b >= 0x20 && b < 0x7F {
			result.WriteByte(b)
		} else {
			result.WriteString(fmt.Sprintf("%02x", b))
		}
	}

	return result.String()
}

func (shaper *huaweiOnuShaper) getOnuId() string {
	if shaper.deviceSN == "" {
		return "HuaweiOnu"
	}
	return shaper.snToOnuID(shaper.deviceSN)
}

// fixTCID corrects the Huawei ONU byte order issues in AR messages
// Analysis: Huawei ONU logs show that AR (request) messages have reversed byte order
// for both TCID and Managed Entity Class ID, compared to their AK (acknowledge) responses.
//
// Examples:
//   AR message: c681 48 0a 0c01 ... (TCID=c681, ME_Class=0c01)
//   AK message: 81c6 28 0a 010c ... (TCID=81c6, ME_Class=010c)
//
// This function reverses both fields for AR=1 messages to ensure AR/AK pairs
// have consistent values.
//
// OMCI message format (first 6 bytes):
// Byte 0-1: TCID (Transaction Correlation Identifier)
// Byte 2: Message Type (bits 0-5) + AR flag (bit 6, 0x40) + AK flag (bit 5, 0x20)
// Byte 3: Device Identifier
// Byte 4-5: Managed Entity Class ID
//
// This behavior is likely due to:
// 1. Huawei's internal implementation details
// 2. Logging system byte order handling
// Rather than intentional obfuscation, as the pattern is consistent and easily detected.
func (shaper *huaweiOnuShaper) fixTCID(omciData string) string {
	if len(omciData) < 12 {
		return omciData
	}

	// Parse the third byte to check AR flag
	// Byte index 2 in the hex string is at position 4-5 (each byte is 2 hex chars)
	thirdByteStr := omciData[4:6]
	thirdByte, err := hex.DecodeString(thirdByteStr)
	if err != nil || len(thirdByte) == 0 {
		return omciData
	}

	// Check if AR bit (bit 6, 0x40) is set
	// In OMCI protocol, byte 2 contains the message type in lower 6 bits
	// and AR flag in bit 6 (0x40). AR=1 means request, AR=0 means response.
	isAR := (thirdByte[0] & 0x40) != 0

	// If AR=1, reverse both TCID (bytes 0-1) and ME Class ID (bytes 4-5)
	if isAR {
		// TCID is at position 0-3 (2 bytes = 4 hex chars)
		byte0 := omciData[0:2] // TCID byte 0
		byte1 := omciData[2:4] // TCID byte 1

		// ME Class ID is at position 8-11 (2 bytes = 4 hex chars)
		byte4 := omciData[8:10]  // ME Class byte 0
		byte5 := omciData[10:12] // ME Class byte 1

		// Reconstruct: reversed_TCID + byte2 + byte3 + reversed_ME_Class + rest
		return byte1 + byte0 + omciData[4:8] + byte5 + byte4 + omciData[12:]
	}

	return omciData
}

func (shaper *huaweiOnuShaper) matchLineInit(s string) bool {
	return strings.Contains(s, "OLT--------->ONT:") ||
		strings.Contains(s, "ONT--------->OLT:")
}

func (shaper *huaweiOnuShaper) processLine(s string) {
	// Check for SN in the line
	if sn := shaper.extractSN(s); sn != "" {
		shaper.deviceSN = sn
		onuID := shaper.getOnuId()
		fileName := shaper.outputPath + onuID
		if _, exists := shaper.onus[onuID]; !exists {
			shaper.onus[onuID] = fileName
			shaper.omciMetaDataMap[onuID] = []OmciMetaData{}
		}
		shaper.ontid = onuID
		return
	}

	switch shaper.state {
	case HuaweiLineInit:
		shaper.handleInit(s)
	case HuaweiLineData0:
		shaper.handleData0(s)
	case HuaweiLineData1:
		shaper.handleData(s, HuaweiLineData2)
	case HuaweiLineData2:
		shaper.handleData(s, HuaweiLineData3)
	case HuaweiLineData3:
		shaper.handleData(s, HuaweiLineInit)
	}
}

func (shaper *huaweiOnuShaper) handleInit(s string) {
	if !shaper.matchLineInit(s) {
		return
	}

	shaper.timeStamp = shaper.getTimestamp(s)
	shaper.state = HuaweiLineData0
}

func (shaper *huaweiOnuShaper) handleData0(s string) {
	// Look for separator line, skip other lines
	if strings.Contains(s, "------------------------------------------------") {
		shaper.state = HuaweiLineData1
	}
	// Stay in Data0 state to wait for separator line
}

func (shaper *huaweiOnuShaper) handleData(s string, nextState OmciState) {
	// Extract hex data from line
	// Format: "fe ff 4f 0a 02 00 00 00 00 00 00 00 00 00 00 00 "
	data := models.StringStrip(s)

	// Remove spaces and validate it's hex
	hexData := strings.ReplaceAll(data, " ", "")

	// Check if this looks like hex data (has at least some hex chars)
	if len(hexData) == 0 {
		shaper.state = HuaweiLineInit
		shaper.omciBuff = ""
		return
	}

	shaper.omciBuff += hexData

	if nextState == HuaweiLineInit {
		if len(shaper.omciBuff) == 80 {
			// Use current ONU ID or default
			currentOnuID := shaper.ontid
			if currentOnuID == "" {
				currentOnuID = shaper.getOnuId()
				fileName := shaper.outputPath + currentOnuID
				if _, exists := shaper.onus[currentOnuID]; !exists {
					shaper.onus[currentOnuID] = fileName
					shaper.omciMetaDataMap[currentOnuID] = []OmciMetaData{}
				}
				shaper.ontid = currentOnuID
			}

			// Fix TCID byte order for AR messages
			correctedData := shaper.fixTCID(shaper.omciBuff)

			shaper.omciMetaDataMap[currentOnuID] = append(shaper.omciMetaDataMap[currentOnuID], OmciMetaData{
				Timestamp: shaper.timeStamp,
				RawData:   correctedData,
			})
		}
		shaper.omciBuff = ""
	}

	shaper.state = nextState
}

func (shaper *huaweiOnuShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("Huawei onu omci shape")
	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}
	shaper.state = HuaweiLineInit

	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}

	defer file.Close()
	reader := models.GetReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			utils.Log("end of read")
			break
		}
		if len(line) == 0 {
			continue
		}
		shaper.processLine(string(line))
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)

	return shaper.onus
}

func init() {
	shaper := new(huaweiOnuShaper)
	shaperRegist(HuaweiOnu, shaper)
}
