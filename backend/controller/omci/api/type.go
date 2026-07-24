package api

import (
	"omciAnalyzer/service/omcianalyzer/omciSchema"

	"github.com/iancoleman/orderedmap"
)

type TypeHandler interface {
	HandleMessage(msg *omciSchema.OmciContext) (r *orderedmap.OrderedMap, e error)
}

var (
	typeRegistry = make(map[uint8]TypeHandler)
)

func RegisterTypeHandler(typeID uint8, h TypeHandler) {
	typeRegistry[typeID] = h
}

func GetTypeHandler(typeID uint8) TypeHandler {
	if h, ok := typeRegistry[typeID]; ok {
		return h
	} else {
		return typeRegistry[0] // default handler
	}
}
