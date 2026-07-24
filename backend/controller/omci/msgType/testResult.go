package msgType

import (
	"omciAnalyzer/controller/omci/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

var (
	TESTRESULT uint8 = 27
)

type TestResult struct {
	//private fields if needed
}

func (m *TestResult) HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error) {
	// handle test result
	contents := omciSchema.OmciTestResultContents{}
	err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
	if err != nil {
		return nil, err
	}
	r = orderedmap.New()
	r.Set("UploadMeClass", contents.UploadMeClass)
	r.Set("UploadMeInstance", contents.UploadMeInstance)
	r.Set("TestResults", contents.TestResults)
	r.Set("Contents", contents.Contents)
	return r, nil
}

func init() {
	api.RegisterTypeHandler(TESTRESULT, &TestResult{})
}
