/*
# ------------------------------------------------------------
# -- mt2HtmlShape.go
# --
# -- Huang Minghe
# -- 2022-8-16
# ------------------------------------------------------------
*/

package omciShape

import (
	"bufio"
	"fmt"
	"html"
	"omciAnalyzer/models"
	"omciAnalyzer/utils"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type mt2HtmlShaper struct {
	outputPath      string
	onus            map[string]string
	omciMetaDataMap map[string][]OmciMetaData
}

var (
	mt2HtmlRowRegexp = regexp.MustCompile(`(?is)<tr[^>]*onclick="drawInfo\([0-9]+\)"[^>]*>.*?</tr>`)
	mt2HtmlTdRegexp  = regexp.MustCompile(`(?is)<td[^>]*>(.*?)</td>`)
	mt2HtmlTagRegexp = regexp.MustCompile(`(?is)<[^>]+>`)
)

func parseMt2HtmlTimestamp(timestampText string) time.Time {
	timestampText = strings.TrimSpace(timestampText)
	if timestampText == "" {
		return time.Time{}
	}

	parts := strings.Split(timestampText, ":")
	if len(parts) != 5 {
		return time.Time{}
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return time.Time{}
	}
	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return time.Time{}
	}
	second, err := strconv.Atoi(parts[2])
	if err != nil {
		return time.Time{}
	}
	msPart := fmt.Sprintf("%03s", strings.TrimSpace(parts[3]))
	fracPart := fmt.Sprintf("%06s", strings.TrimSpace(parts[4]))
	nanoPart, err := strconv.Atoi(msPart + fracPart)
	if err != nil {
		return time.Time{}
	}

	return time.Date(1970, time.January, 1, hour, minute, second, nanoPart, time.UTC)
}

func cleanMt2HtmlText(s string) string {
	withoutTag := mt2HtmlTagRegexp.ReplaceAllString(s, "")
	return strings.TrimSpace(html.UnescapeString(withoutTag))
}

func parseMt2HtmlRows(data string) []string {
	rows := mt2HtmlRowRegexp.FindAllString(data, -1)
	if len(rows) == 0 {
		return nil
	}
	return rows
}

func (shaper *mt2HtmlShaper) appendOmciRow(row string) {
	columns := mt2HtmlTdRegexp.FindAllStringSubmatch(row, -1)
	if len(columns) < 12 {
		return
	}

	timestamp := parseMt2HtmlTimestamp(cleanMt2HtmlText(columns[1][1]))
	serial := cleanMt2HtmlText(columns[7][1])
	if serial == "" {
		serial = "MT2-Onu"
	}

	rawText := models.StringStrip(cleanMt2HtmlText(columns[11][1]))
	if len(rawText) < 80 {
		return
	}
	rawData := rawText[:80]

	if _, exists := shaper.onus[serial]; !exists {
		fileName := shaper.outputPath + serial
		shaper.onus[serial] = fileName
		shaper.omciMetaDataMap[serial] = []OmciMetaData{}
	}

	shaper.omciMetaDataMap[serial] = append(shaper.omciMetaDataMap[serial], OmciMetaData{
		Timestamp: timestamp,
		RawData:   rawData,
	})
}

func (shaper *mt2HtmlShaper) Shape(logPath string, outputPath string) map[string]string {
	utils.Log("MT2 HTML OMCI shape processing started")

	shaper.outputPath = outputPath
	shaper.onus = map[string]string{}
	shaper.omciMetaDataMap = map[string][]OmciMetaData{}

	file, err := os.Open(logPath)
	if err != nil {
		utils.Log("open file failed: " + fmt.Sprintf("%s", err))
		return nil
	}
	defer file.Close()

	var htmlBuilder strings.Builder
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 1024), 10*1024*1024)
	for scanner.Scan() {
		htmlBuilder.WriteString(scanner.Text())
		htmlBuilder.WriteByte('\n')
	}
	if err := scanner.Err(); err != nil {
		utils.Log("read file failed: " + fmt.Sprintf("%s", err))
		return nil
	}

	rows := parseMt2HtmlRows(htmlBuilder.String())
	for _, row := range rows {
		shaper.appendOmciRow(row)
	}

	flushOmciDataToFile(shaper.onus, shaper.omciMetaDataMap)
	utils.Log("MT2 HTML OMCI shape processing completed, processed " + fmt.Sprintf("%d", len(shaper.onus)) + " ONUs")

	return shaper.onus
}

func init() {
	shaper := new(mt2HtmlShaper)
	shaperRegist(Mt2Html, shaper)
}
