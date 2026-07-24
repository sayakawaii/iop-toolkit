package msgType

import (
	"fmt"
	"omciAnalyzer/controller/omci/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

var (
	MIBUPLOADNEXT uint8 = 14
)

type MibUploadNext struct {
	//private fields if needed
}

func (m *MibUploadNext) HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error) {
	// handle mibupload next
	if msg.Header.AK == 1 {
		contents := omciSchema.OmciMibuploadNextAkContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		attrsList := omciSchema.SliceMapConvToMap(contents.Attrs)
		r = orderedmap.New()
		r.Set("UploadMe", contents.UploadMeName)
		r.Set("UploadMeInstance", fmt.Sprintf("%d", contents.UploadMeInstance)+` (0x`+fmt.Sprintf("%X", contents.UploadMeInstance)+`)`)
		r.Set("Attrs", api.OmciContentAttrsToJSON(attrsList))
	} else {
		contents := omciSchema.OmciMibuploadNextArContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		r = orderedmap.New()
		r.Set("Command Sequence Number", contents.CommandSequenceNum)
	}
	return r, nil
}

func init() {
	api.RegisterTypeHandler(MIBUPLOADNEXT, &MibUploadNext{})
}
