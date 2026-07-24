package msgType

import (
	"omciAnalyzer/controller/omci/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

var (
	MIBUPLOAD uint8 = 13
)

type MibUpload struct {
	//private fields if needed
}

func (m *MibUpload) HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error) {
	// handle mibupload
	if msg.Header.AK == 1 {
		contents := omciSchema.OmciMibuploadAkContents{}
		err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
		if err != nil {
			return nil, err
		}
		r = orderedmap.New()
		r.Set("Command Number", contents.NumMIBUploadCommands)
	} else {
		r = orderedmap.New()
	}
	return r, nil
}

func init() {
	api.RegisterTypeHandler(MIBUPLOAD, &MibUpload{})
}
