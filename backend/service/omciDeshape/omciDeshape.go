package omciDeshape

import (
	"errors"
	"io"
	"omciAnalyzer/utils"
	"os"
)

type omciDeshape interface {
	Translate(io.Reader, io.Writer) error
}

var (
	allDeshapes = make(map[string]omciDeshape)
)

func newDeshape(k string, d omciDeshape) {
	allDeshapes[k] = d
}

func OmciDeshapeTrans(k string, inFile string, outFile string) error {
	s, exist := allDeshapes[k]
	if !exist {
		return errors.New("the intented omciShape not existing")
	}
	iFile, err := os.Open(inFile)
	if err != nil {
		utils.Log("open input file failed")
		return errors.New("open input file failed")
	}
	defer iFile.Close()

	oFile, err := os.Create(outFile)
	if err != nil {
		utils.Log("open output file failed")
		return errors.New("open output file failed")
	}
	defer oFile.Close()
	return s.Translate(iFile, oFile)
}
