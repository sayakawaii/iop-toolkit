package omciSchema

import (
	"fmt"
)

type SchemaDef struct {
	schemaList      SchemaListType
	meClassIdList   MeClassIdListType
	omciMsgTypeList OmciMsgTypeListType
	Path            string
}

var schemaDef *SchemaDef
var schemaStandard *SchemaDef

func (def *SchemaDef) MeSchemaList() SchemaListType {
	return def.schemaList
}

func (def *SchemaDef) OmciMsgTypeList() OmciMsgTypeListType {
	return def.omciMsgTypeList
}

func (def *SchemaDef) MeSchema(key SchemaKey) (Schema, error) {
	meSchema, found := def.schemaList[key]
	if !found {
		return Schema{}, fmt.Errorf("ME [%v] do not exist", key)
	}
	return *meSchema, nil
}

func (def *SchemaDef) MeClassIdList() MeClassIdListType {
	return def.meClassIdList
}

func (def *SchemaDef) MeSchemaByClassId(id ClassIdType) (Schema, error) {
	key, found := def.meClassIdList[id]
	if !found {
		return Schema{}, fmt.Errorf("ME Class Id [%v] not found", id)
	}
	return def.MeSchema(*key)
}

func DefaultSchemaDef() *SchemaDef {
	return schemaDef
}

func StandardSchemaDef() *SchemaDef {
	return schemaStandard
}

type OnuInfo struct {
	schemaDef *SchemaDef
}

func (onuInfo *OnuInfo) SchemaDef() *SchemaDef {
	return onuInfo.schemaDef
}

func init() {
	LoadSchema()
}
