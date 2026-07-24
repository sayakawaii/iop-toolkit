package msgType

import (
	"omciAnalyzer/controller/omci/api"
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

var (
	ALARM uint8 = 16
)

type Alarm struct {
	//private fields if needed
}

func (m *Alarm) HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error) {
	// handle alarm
	contents := omciSchema.OmciAlarmContents{}
	err := omciSchema.GetContentFromInterface(msg.Contents, &contents)
	if err != nil {
		return nil, err
	}
	r = orderedmap.New()
	r.Set("AlarmSeqNum", contents.AlarmSeqNum)
	r.Set("AlarmName", contents.AlarmName)
	r.Set("AlarmDescription", contents.AlarmDescription)
	return r, nil
}

func init() {
	api.RegisterTypeHandler(ALARM, &Alarm{})
}
