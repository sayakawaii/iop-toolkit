package api

import (
	"fmt"
	"omciAnalyzer/service/omcianalyzer/omciSchema"
)

func Dispatch(msg *omciSchema.OmciContext) error {
	var handled bool

	if h := GetTypeHandler(msg.Header.MsgType); h != nil {
		if _, err := h.HandleMessage(msg); err != nil {
			return err
		}
		handled = true
	}

	if h := GetEntityHandler(msg.Header.MeClass); h != nil {
		if _, err := h.HandleEntity(msg); err != nil {
			return err
		}
		handled = true
	}

	if !handled {
		return fmt.Errorf("no handler for type=%d entity=%d", msg.Header.MsgType, msg.Header.MeClass)
	}

	return nil
}
