package msgType

import (
	"omciAnalyzer/controller/omci/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

var (
	TEST uint8 = 18
)

type Test struct {
	//private fields if needed
}

func (m *Test) HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error) {
	if msg.Header.AR == 1 {
		// handle test ar
		contents := omciSchema.OmciTestArContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		r = orderedmap.New()
		r.Set("Payload", contents.Payload)
		r.Set("TestRequestAttrs", contents.TestRequestAttrs)
		return r, nil
	} else {
		// handle test ak
		contents := omciSchema.OmciTestAkContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		r = orderedmap.New()
		r.Set("Result", contents.Result)
		return r, nil
	}
}

func init() {
	api.RegisterTypeHandler(TEST, &Test{})
}
