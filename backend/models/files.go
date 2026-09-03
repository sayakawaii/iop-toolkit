package models

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"omciAnalyzer/utils"
	"os"
	"path/filepath"
	"strings"
)

func GetReader(file *os.File) *bufio.Reader {
	text, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//for old mac file, CR as newline, so replace it with LF
	text = bytes.ReplaceAll(text, []byte("\r"), []byte("\n"))
	return bufio.NewReader(strings.NewReader(string(text)))
}

func WriteFile(fileName string, line string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		utils.Log("open file failed")
	} else {
		// utils.Log("open file success")
		defer file.Close()
		file.WriteString(line)
	}
}

// Function to read directory and build HTML tree structure
func BuildDownloadsTree(path string) string {
	var result strings.Builder
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	result.WriteString("<ul>\n") // Start the tree root

	for _, file := range files {
		if file.IsDir() {
			// If it's a directory, recursively read its contents
			result.WriteString(fmt.Sprintf("<li class=\"isFolder\">\n%s\n", file.Name()))
			result.WriteString(BuildDownloadsTree(filepath.Join(path, file.Name())))
			result.WriteString("</li>\n")
		} else {
			// If it's a file, create a link to download it
			filePath := filepath.Join(path, file.Name())
			result.WriteString(fmt.Sprintf("<li><a href=\"/%s\" class=\"file-download\">%s</a></li>\n", filePath, file.Name()))
		}
	}

	result.WriteString("</ul>\n") // Close the tree root
	return result.String()
}
