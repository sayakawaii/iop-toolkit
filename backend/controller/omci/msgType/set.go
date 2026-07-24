package msgType

import (
	"fmt"
	"omciAnalyzer/controller/omci/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

var (
	SET uint8 = 8
)

type Set struct {
	//private fields if needed
}

func (m *Set) HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error) {
	if msg.Header.AK == 1 {
		// handle service
		contents := omciSchema.OmciSetAkContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		r = orderedmap.New()
		if contents.Result == api.AttrFailOrUnknown {
			r.Set("Result", contents.Result)
			r.Set("OptAttrMask", contents.OptAttrMask)
			r.Set("AttrExecMask", fmt.Sprintf("0x%X", contents.OptAttrMask))
		} else {
			r.Set("Result", contents.Result)
		}
	} else {
		contents := omciSchema.OmciArContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		attrsList := omciSchema.SliceMapConvToMap(contents.Attrs)
		r = orderedmap.New()
		r.Set("AttrMask", fmt.Sprintf("0x%X", contents.AttrMask))
		r.Set("Attrs", api.OmciContentAttrsToJSON(attrsList))
	}
	return r, nil
}

func init() {
	api.RegisterTypeHandler(SET, &Set{})
}
