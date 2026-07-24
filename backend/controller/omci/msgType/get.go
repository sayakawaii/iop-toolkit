package msgType

import (
	"fmt"
	"omciAnalyzer/controller/omci/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

var (
	GET uint8 = 9
)

type Get struct {
	//private fields if needed
}

func (m *Get) HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error) {
	if msg.Header.AK == 1 {
		// handle service
		contents := omciSchema.OmciGetAkContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		attrsList := omciSchema.SliceMapConvToMap(contents.Attrs)
		r = orderedmap.New()
		if contents.Result == api.AttrFailOrUnknown {
			r.Set("Result", contents.Result)
			r.Set("AttrMask", fmt.Sprintf("0x%X", contents.AttrMask))
			r.Set("OptAttrMask", fmt.Sprintf("0x%X", contents.OptAttrMask))
			r.Set("AttrExecMask", fmt.Sprintf("0x%X", contents.AttrExecMask))
		} else {
			r.Set("Result", contents.Result)
			r.Set("AttrMask", fmt.Sprintf("0x%X", contents.AttrMask))
			r.Set("Attrs", api.OmciContentAttrsToJSON(attrsList))
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
	api.RegisterTypeHandler(GET, &Get{})
}
