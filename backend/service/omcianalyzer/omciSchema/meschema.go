package omciSchema

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"omciAnalyzer/utils"
	"os"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
)

var jsonMeDirPath string = "./resource/me/schema/"
var jsonMeDirPathTest string = "../../../resource/me/schema/"

//easyjson:json
type Schema struct {
	Me struct {
		ClassId  uint16   `json:"id"`
		Name     string   `json:"name"`
		Access   string   `json:"access"`
		Type     string   `json:"type"`
		Action   []string `json:"action"`
		Instance struct {
			Type  string `json:"type"`
			Value uint16 `json:"value"`
		} `json:"instance"`
		Attributes struct {
			Attribute []AttributeSchema `json:"attribute"`
			Mask      uint16
			SbcMask   uint16
		} `json:"attributes"`
		ActionMap ActionMapType
		Alarms    struct {
			Alarm []AlarmSchema `json:"alarm"`
		} `json:"alarms"`
		TestRequestAttributes struct {
			TestRequestAttribute []AttributeSchema `json:"attribute"`
		} `json:"testRequestAttrs"`
		TestResultAttributes struct {
			TestResultAttribute []AttributeSchema `json:"attribute"`
		} `json:"testResultAttrs"`
	} `json:"Me"`
}

//easyjson:json
type AttributeSchema struct {
	Name     string             `json:"name"`
	Position uint16             `json:"position"`
	Size     uint16             `json:"size"`
	Format   string             `json:"format"`
	IsTable  bool               `json:"table"`
	Access   []string           `json:"access"`
	Category string             `json:"category"`
	Table    AttributeStructure `json:"structure"`
}

//easyjson:json
type AttributeStructure struct {
	Fields []AttributeStructureField `json:"fields"`
}

//easyjson:json
type AttributeStructureField struct {
	Name          string             `json:"name"`
	Size          uint16             `json:"size"`
	Format        string             `json:"format"`
	ExpectedValue uint16             `json:"expectedValue"`
	Table         AttributeStructure `json:"structure"`
}

type AttributeStructureFieldValue struct {
	Field AttributeStructureField
	Value interface{}
}

// type ShriGaneshOmci struct {
type ActionMapType map[string]string

//easyjson:json
type AlarmSchema struct {
	Number      uint8  `json:"number"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// type SchemaKey uint16
const turboSectionSize uint16 = 1440 //section size 4 SWDLD Turbo
type SchemaKey string
type SchemaListType map[SchemaKey]*Schema
type ClassIdType uint16
type MeClassIdListType map[ClassIdType]*SchemaKey

type SectionPayloadType []byte

func LoadSchema() {
	path := ""
	if os.Getenv("RUN_MODE") == "test" {
		path = jsonMeDirPathTest
	} else {
		path = jsonMeDirPath
	}
	files, err := os.ReadDir(path)
	if err != nil {
		utils.Log("LoadSchema")
	}
	schemaDef = &SchemaDef{
		schemaList:    make(SchemaListType),
		meClassIdList: make(MeClassIdListType),
	}
	for _, file := range files {
		readSchema(path+file.Name(), schemaDef.schemaList, schemaDef.meClassIdList, false)
	}
	schemaStandard = &SchemaDef{
		schemaList:    make(SchemaListType),
		meClassIdList: make(MeClassIdListType),
	}
	for _, file := range files {
		readSchema(path+file.Name(), schemaStandard.schemaList, schemaStandard.meClassIdList, true)
	}
}

func NewSchema() *Schema {
	return new(Schema).init()
}

func (schema *Schema) Name() string {
	return schema.Me.Name
}

func (schema *Schema) ClassId() uint16 {
	return schema.Me.ClassId
}

func (schema *Schema) ParseTestAttrs(attributeSchema []AttributeSchema, msg *OmciRespMsg) (attrs []Attribute) {
	for _, attr := range attributeSchema {
		value, conversionCompleted := attr.value(msg.content.payload[attr.Position:])
		if conversionCompleted {
			attrs = append(attrs, Attribute{Name: attr.Name, Value: value})
		}
	}
	return
}

func (schema *Schema) ParseTxOmci(msg *OmciMsg) (attrs []Attribute) {
	var loc uint16
	loc = 0
	mask := msg.content.AttrMask
	allAttrMask := schema.Me.Attributes.Mask
	if MT(msg.header.MsgType) == Create {
		mask = schema.Me.Attributes.SbcMask
	}
	if mask > allAttrMask {
		utils.Log("Invalid OMCI attribute mask")
	}
	contentPayloadlength := uint16(len(msg.content.payload))

	for n := 15; n > (15-len(schema.Me.Attributes.Attribute)) && mask != 0; n-- {
		m := uint16(1 << uint(n))
		if mask&m != 0 {
			a := schema.Me.Attributes.Attribute[15-n]
			tmp := loc + a.Size
			if tmp > contentPayloadlength && !a.IsTable {
				utils.Log("Invalid mask: contains more attributes than the payload can have.")
				return
			}
			value, _ := a.value(msg.content.payload[loc:])
			attrs = append(attrs, Attribute{Name: a.Name, Value: value})
			loc += a.Size
			mask &= ^m
		}
	}
	return
}

func (schema *Schema) Attributes(mask uint16, payload []byte) (attrs []Attribute, e error) {
	e = nil
	loc := uint16(0)
	allAttrMask := schema.Me.Attributes.Mask
	if mask > allAttrMask {
		e = errors.New(fmt.Sprintf("Invalid OMCI attribute mask [%v] for ME [%v] Class Id [%v] max possible mask [%v]", mask, schema.Name(), schema.ClassId(), allAttrMask))
	}
	contentPayloadlength := uint16(len(payload))
	for n := 15; n > (15-len(schema.Me.Attributes.Attribute)) && mask != 0; n-- {
		m := uint16(1 << uint(n))
		if mask&m != 0 {
			a := schema.Me.Attributes.Attribute[15-n]
			tmp := loc + a.Size
			if tmp > contentPayloadlength && !a.IsTable {
				utils.Log("Invalid mask : contains more attributes than the payload can have.")
				return
			}
			orderInSchema := uint8(15 - n)
			if a.isTableAttr() {
				// tableAttrSize := bytesUInteger(payload, 4)
				// attrs = append(attrs, Attribute{Name: a.Name, Value: tableAttrSize, orderInMeSchema: orderInSchema})
				/*[Huang Minghe]update this logical to adapt MOP DACL table parser*/
				if a.Size <= loc {
					fmt.Println("invalid OMCI: ", payload, " size: ", a.Size, " loc: ", loc)
					fmt.Println("schema: ", schema.Me)
					tableAttrSize := bytesUInteger(payload, 4)
					attrs = append(attrs, Attribute{Name: a.Name, Value: tableAttrSize, orderInMeSchema: orderInSchema})
				} else {
					// Clamp the upper bound to the payload length. Table attributes
					// bypass the length guard above, so a schema whose declared size
					// exceeds the received payload (e.g. proprietary MEs only present
					// in the default schema) would otherwise slice out of range.
					end := a.Size
					if end > contentPayloadlength {
						end = contentPayloadlength
					}
					value, _ := a.value(payload[loc:end])
					attrs = append(attrs, Attribute{Name: a.Name, Value: value, orderInMeSchema: orderInSchema})
				}
			} else {
				value, _ := a.value(payload[loc:])
				attrs = append(attrs, Attribute{Name: a.Name, Value: value, orderInMeSchema: orderInSchema})
			}
			loc += a.Size
			mask &= ^m
		}
	}
	return
}

func (schema *Schema) getMeAttributeName(attrOrder uint8) (attrStr string) {
	attrStr = ""
	attr := schema.Me.Attributes.Attribute[attrOrder]
	if strings.Compare(attr.Name, "") != 0 {
		attrStr = attr.Name
	}
	return attrStr
}

func (schema *Schema) getMeAttributeOrder(attrName string) (attrOrder uint8, e error) {
	for itr, attribute := range schema.Me.Attributes.Attribute {
		if strings.Compare(attribute.Name, attrName) == 0 {
			return uint8(itr), nil
		}
	}
	return 0, errors.New("attribute does not exist")
}

func (schema *Schema) getAlarmListLength() (alarmListLength uint8) {
	maxBit := uint8(0)
	alarmListLength = 0
	for _, alarm := range schema.Me.Alarms.Alarm {
		if alarm.Number > maxBit {
			maxBit = alarm.Number
		}
	}
	if len(schema.Me.Alarms.Alarm) != 0 {
		if (maxBit+1)%8 != 0 {
			alarmListLength = (maxBit+1)/8 + 1
		} else {
			alarmListLength = (maxBit + 1) / 8
		}
	}
	return
}

/*
   *************************
   Internal function section
   *************************
*/

func readSchema(fileName string, schemaList SchemaListType, meClassIdList MeClassIdListType, standardOnly bool) (err error) {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		utils.Log("readSchema:")
		return
	}

	var schema Schema
	schema.init()
	//schema = NewSchema()
	err = easyjson.Unmarshal(raw, &schema)
	if err != nil {
		utils.Log("readSchema Unmarshal")
		return
	}
	if standardOnly && ((349 < schema.Me.ClassId && schema.Me.ClassId < 400) || schema.Me.ClassId > 65279) {
		//skip unsupported Property ME
		return
	}
	//Fast Mapping
	for _, action := range schema.Me.Action {
		schema.Me.ActionMap[action] = action
	}
	schema.Me.Attributes.Mask = schema.allAttrMask()
	schema.Me.Attributes.SbcMask = schema.sbcAttrMask()
	//schemaList[SchemaKey(schema.Me.ClassId)] = schema
	schemaList[SchemaKey(schema.Me.Name)] = &schema
	k := SchemaKey(schema.Me.Name)
	meClassIdList[ClassIdType(schema.Me.ClassId)] = &k
	return
}

func processKeyBasedParser(elementSizeInBits uint16, field AttributeStructureField, bytesElementFullString string) (r map[string]interface{}) {
	structureMap := make(map[string]interface{})
	//currBitLoc := schema.Size * 8 //the LSB at the far right
	currBitLoc := uint16(0)
	//previousBitLoc := schema.Size * 8
	previousBitLoc := uint16(0)
	currBitLoc = currBitLoc + field.Size
	if currBitLoc > elementSizeInBits {
		return nil
	}
	//schema.Size   //bytes //the whole element
	//field.Size    //bits  //only the particular field
	fieldValueString := bytesElementFullString[previousBitLoc:currBitLoc]
	//TODO: Maybe extract to function in util.go
	if field.Format == "UnsignedInteger" {
		fieldValueUint, _ := strconv.ParseUint(fieldValueString, 2, 64)
		if field.Size <= 8 {
			structureMap[field.Name] = uint8(fieldValueUint)
		} else if field.Size <= 16 {
			structureMap[field.Name] = uint16(fieldValueUint)
		} else if field.Size <= 32 {
			structureMap[field.Name] = uint32(fieldValueUint)
		} else if field.Size <= 64 {
			structureMap[field.Name] = fieldValueUint
		}
	}
	if field.Format == "SignedInteger" {
		fieldValueInt, _ := strconv.ParseInt(fieldValueString, 2, 64)
		if field.Size <= 8 {
			structureMap[field.Name] = int8(fieldValueInt)
		} else if field.Size <= 16 {
			structureMap[field.Name] = int16(fieldValueInt)
		} else if field.Size <= 32 {
			structureMap[field.Name] = int32(fieldValueInt)
		} else if field.Size <= 64 {
			structureMap[field.Name] = fieldValueInt
		}
	}
	if field.Format == "String" {
		bytesOfString := []byte{}                 //string as byte array
		numOfBytes := field.Size / 8              //we assume that size is a multiply of 8
		for i := uint16(0); i < numOfBytes; i++ { //TODO:Unfortunately this is a triple-for-loop
			elementBinaryString := fieldValueString[i*8 : (i+1)*8]           //binary string per byte
			elementUint64, _ := strconv.ParseUint(elementBinaryString, 2, 8) // uint64 per byte
			bytesOfString = append(bytesOfString, byte(elementUint64))
		}
		structureMap[field.Name] = string(bytesOfString) //convert byte array to string
	}
	if field.Format == "Bytes" {
		ri := []uint16{}
		numOfBytes := field.Size / 8 //we assume that size is a multiply of 8
		/* Array Structure Byte attributes are constructed LSB->MSB in the OMCI */
		for i := uint16(0); i < numOfBytes; i++ {
			elementBinaryString := fieldValueString[i*8 : (i+1)*8]
			//for i := numOfBytes; i > uint16(0); i-- { //TODO:Unfortunately this is a triple-for-loop
			//elementBinaryString := fieldValueString[(i-1)*8 : i*8]           //binary string per byte
			elementUint64, _ := strconv.ParseUint(elementBinaryString, 2, 8) // uint64 per byte
			ri = append(ri, uint16(elementUint64))
		}
		structureMap[field.Name] = ri
	}
	if field.Format == "Boolean" {
		fieldValueUint, _ := strconv.ParseUint(fieldValueString, 2, 64)
		if field.Size <= 8 {
			structureMap[field.Name] = uint8(fieldValueUint)
		} else if field.Size <= 16 {
			structureMap[field.Name] = uint16(fieldValueUint)
		} else if field.Size <= 32 {
			structureMap[field.Name] = uint32(fieldValueUint)
		} else if field.Size <= 64 {
			structureMap[field.Name] = fieldValueUint
		}
	}
	if field.Format == "KeyBasedParser" {
		innerStructMap := processKeyBasedParser(field.Size, field.Table.Fields[0], fieldValueString)
		structureMap[field.Name] = innerStructMap
	}
	//Array Structure (but actually just structure)
	if strings.Contains(field.Format, "Structure") {
		innerStructMap := processStructure(field.Size, field.Table, fieldValueString)
		structureMap[field.Name] = innerStructMap
	}
	previousBitLoc = currBitLoc
	r = structureMap
	return
}

var myKey uint8

func processStructure(elementSizeInBits uint16, table AttributeStructure, bytesElementFullString string) (r map[string]interface{}) {
	structureMap := make(map[string]interface{})
	//currBitLoc := schema.Size * 8 //the LSB at the far right
	currBitLoc := uint16(0)
	//previousBitLoc := schema.Size * 8
	previousBitLoc := uint16(0)
	//for _, field := range schema.Table.Fields {
	for _, field := range table.Fields {
		currBitLoc = currBitLoc + field.Size
		if currBitLoc > elementSizeInBits {
			break
		}
		//schema.Size   //bytes //the whole element
		//field.Size    //bits  //only the particular field
		fieldValueString := bytesElementFullString[previousBitLoc:currBitLoc]
		//TODO: Maybe extract to function in util.go
		if field.Format == "UnsignedInteger" {
			fieldValueUint, _ := strconv.ParseUint(fieldValueString, 2, 64)
			if field.Size <= 8 {
				structureMap[field.Name] = uint8(fieldValueUint)
			} else if field.Size <= 16 {
				structureMap[field.Name] = uint16(fieldValueUint)
			} else if field.Size <= 32 {
				structureMap[field.Name] = uint32(fieldValueUint)
			} else if field.Size <= 64 {
				structureMap[field.Name] = fieldValueUint
			}
		}
		if field.Format == "SignedInteger" {
			fieldValueInt, _ := strconv.ParseInt(fieldValueString, 2, 64)
			if field.Size <= 8 {
				structureMap[field.Name] = int8(fieldValueInt)
			} else if field.Size <= 16 {
				structureMap[field.Name] = int16(fieldValueInt)
			} else if field.Size <= 32 {
				structureMap[field.Name] = int32(fieldValueInt)
			} else if field.Size <= 64 {
				structureMap[field.Name] = fieldValueInt
			}
		}
		if field.Format == "String" {
			bytesOfString := []byte{}                 //string as byte array
			numOfBytes := field.Size / 8              //we assume that size is a multiply of 8
			for i := uint16(0); i < numOfBytes; i++ { //TODO:Unfortunately this is a triple-for-loop
				elementBinaryString := fieldValueString[i*8 : (i+1)*8]           //binary string per byte
				elementUint64, _ := strconv.ParseUint(elementBinaryString, 2, 8) // uint64 per byte
				bytesOfString = append(bytesOfString, byte(elementUint64))
			}
			structureMap[field.Name] = string(bytesOfString) //convert byte array to string
		}
		if field.Format == "Bytes" {
			ri := []uint16{}
			numOfBytes := field.Size / 8 //we assume that size is a multiply of 8
			/* Array Structure Byte attributes are constructed LSB->MSB in the OMCI */
			for i := uint16(0); i < numOfBytes; i++ {
				elementBinaryString := fieldValueString[i*8 : (i+1)*8]
				//for i := numOfBytes; i > uint16(0); i-- { //TODO:Unfortunately this is a triple-for-loop
				//elementBinaryString := fieldValueString[(i-1)*8 : i*8]           //binary string per byte
				elementUint64, _ := strconv.ParseUint(elementBinaryString, 2, 8) // uint64 per byte
				ri = append(ri, uint16(elementUint64))
			}
			if field.Size == 96 && (field.Name == "Leading bytes of IPv6 source address" || field.Name == "Leading bytes of IPv6 destination address") {
				uint16ArrayToIPV6String := func(arr []uint16) string {
					hexString := ""
					count := 0
					for _, value := range arr {
						hexString += fmt.Sprintf("%02X", value)
						count = count + 1
						if count%2 == 0 {
							hexString += ":"
						}
					}
					hexString += ":"
					return hexString
				}
				structureMap[field.Name] = uint16ArrayToIPV6String(ri)
			} else if field.Size == 32 && (field.Name == "Source IP address" ||
				field.Name == "Destination IP address start of range" ||
				field.Name == "Destination IP address end of range") {
				uint16ArrayToIPV4String := func(arr []uint16) string {
					hexString := ""
					for _, value := range arr {
						hexString += fmt.Sprintf("%d.", value)
					}
					return hexString[:len(hexString)-1]
				}
				structureMap[field.Name] = uint16ArrayToIPV4String(ri)
			} else {
				structureMap[field.Name] = ri
			}
		}
		if field.Format == "Boolean" {
			fieldValueUint, _ := strconv.ParseUint(fieldValueString, 2, 64)
			if field.Size <= 8 {
				structureMap[field.Name] = uint8(fieldValueUint)
			} else if field.Size <= 16 {
				structureMap[field.Name] = uint16(fieldValueUint)
			} else if field.Size <= 32 {
				structureMap[field.Name] = uint32(fieldValueUint)
			} else if field.Size <= 64 {
				structureMap[field.Name] = fieldValueUint
			}
		}
		if field.Name == "Row part ID" {
			myKey = structureMap[field.Name].(uint8)
		}
		if field.Format == "KeyBasedParser" {
			innerStructMap := processKeyBasedParser(field.Size, field.Table.Fields[myKey], fieldValueString)
			structureMap[field.Name] = innerStructMap
		}
		//Array Structure (but actually just structure)
		if strings.Contains(field.Format, "Structure") {
			innerStructMap := processStructure(field.Size, field.Table, fieldValueString)
			structureMap[field.Name] = innerStructMap
		}
		previousBitLoc = currBitLoc
	}
	r = structureMap
	return
}

func (schema *AttributeSchema) value(byteArray []byte) (r interface{}, conversionCompleted bool) {
	conversionCompleted = true
	if schema.Format == "Boolean" {
		r = bytesUInteger(byteArray, schema.Size)
	}
	if schema.Format == "SignedInteger" {
		r = bytesInteger(byteArray, schema.Size)
	}
	if schema.Format == "UnsignedInteger" {
		r = bytesUInteger(byteArray, schema.Size)
	}
	if schema.Format == "String" {
		// FNMS-68031 Begin
		firstZeroByte := (bytes.IndexByte(byteArray, byte(0)))
		if firstZeroByte != -1 && uint16(firstZeroByte) < schema.Size {
			r = string(byteArray[:firstZeroByte])
		} else {
			r = string(byteArray[:schema.Size])
		}
		//FNMS-68031 End
	}
	if schema.Format == "Bytes" {
		ri := make([]uint16, schema.Size)
		for i, b := range byteArray {
			if i < int(schema.Size) {
				ri[i] = uint16(b)
			} else { /*  */
				break
			}
		}
		if schema.Size >= 16 && schema.Name != "DSCP to p-bit mapping" {
			uint16ArrayToHexString := func(arr []uint16) string {
				hexString := ""
				for _, value := range arr {
					hexString += fmt.Sprintf("%02X ", value)
				}
				return hexString
			}
			r = uint16ArrayToHexString(ri)
		} else {
			r = ri
		}
	}
	if schema.Format == "TestStructure" {
		var fields []AttributeStructureField
		fields = schema.Table.Fields
		if len(fields) == 2 {
			expectedValueFromOmciMsg := bytesUInteger(byteArray, (fields[0].Size)/8)
			expectedValueFromSpec := testResultExpectedValue(uint64(fields[0].ExpectedValue), fields[0].Size)
			if expectedValueFromSpec == expectedValueFromOmciMsg {
				firstFieldSizeInBytes := (fields[0].Size) / 8
				secondFieldSizeInBytes := (fields[1].Size) / 8
				if fields[1].Format == "UnsignedInteger" {
					r = bytesUInteger(byteArray[firstFieldSizeInBytes:], secondFieldSizeInBytes)
				} else if fields[1].Format == "SignedInteger" {
					r = bytesInteger(byteArray[firstFieldSizeInBytes:], secondFieldSizeInBytes)
				}
			} else {
				conversionCompleted = false
			}
		} else {
			utils.Log("TestStructure has no 2 fields, can't compute the value of the attribute")
		}
	}
	if strings.Contains(schema.Format, "Array") {
		format := strings.Split(schema.Format, " ")
		if format[1] == "UnsignedInteger" {
			var ret []interface{}
			if schema.IsTable {
				numberOfArrayElements := uint16(len(byteArray)) / schema.Size
				if numberOfArrayElements == 0 {
					return
				}
				for i := uint16(0); i < numberOfArrayElements; i++ {
					bytesPerElement := byteArray[i*schema.Size : (i+1)*schema.Size]
					ri := bytesUInteger(bytesPerElement, schema.Size)
					ret = append(ret, ri)
				}
			} else {
				if uint16(len(byteArray)) < schema.Size {
					return
				}
				bitSizeAttr, _ := stringToInt64(format[2])
				if bitSizeAttr < 8 {
					ret = append(ret, 0)
					return
				}
				byteSizeAttr := uint16(bitSizeAttr) / 8
				tableAttrSize := schema.Size / byteSizeAttr

				for i := uint16(0); i < tableAttrSize; i++ {
					bytesSize := byteArray[i*byteSizeAttr : (i+1)*byteSizeAttr]
					ri := bytesUInteger(bytesSize, byteSizeAttr)
					ret = append(ret, ri)
				}
			}
			r = ret
		} else if format[1] == "Structure" {
			structureMapArray := []map[string]interface{}{}
			if uint16(len(byteArray)) < schema.Size {
				return
			}
			bytesElement := byteArray[0:schema.Size]
			bytesElementStrings := []string{}
			for _, v := range bytesElement {
				str := strconv.FormatUint(uint64(v), 2)
				pad := strings.Repeat("0", 8-len(str))
				str = pad + str
				bytesElementStrings = append(bytesElementStrings, str)
			}
			bytesElementFullString := strings.Join(bytesElementStrings, "")
			structureMap := processStructure(schema.Size*8, schema.Table, bytesElementFullString)
			structureMapArray = append(structureMapArray, structureMap)

			r = structureMapArray
		}
	}
	return
}

func testResultExpectedValue(expectedValue uint64, size uint16) interface{} {
	if size == 8 {
		return uint8(expectedValue)
	}
	if size == 16 {
		return uint16(expectedValue)
	}
	if size == 32 {
		return uint32(expectedValue)
	}
	if size == 64 {
		return uint64(expectedValue)
	} else {
		utils.Log("Invalid test result expected value - size")
		return nil
	}
}

func (schema *AttributeSchema) isTableAttr() bool {
	if schema.IsTable {
		return true
	}
	return false
}
func (schema *AttributeSchema) isAccess(s string) bool {
	for _, v := range schema.Access {
		if v == s {
			return true
		}
	}
	return false
}

func (schema *AttributeSchema) isSBC() bool {
	return schema.isAccess("SBC")
}

func (schema *Schema) init() *Schema {
	schema.Me.ActionMap = make(ActionMapType)
	return schema
}

func (schema *Schema) instance(inst uint16) uint16 {
	if schema.Me.Instance.Type == "Fixed" {
		return schema.Me.Instance.Value
	}
	return inst
}

func (table *AttributeStructure) sizeInBytes() (ret uint16) {
	ret = 0
	for _, field := range table.Fields {
		ret += field.sizeInBits()
	}
	ret = uint16(math.Ceil(float64(ret / 8)))
	return
}

func (field *AttributeStructureField) sizeInBits() uint16 {
	return field.Size
}

func (field *AttributeStructureField) bits(v interface{}) (ret []bool) {
	fieldBytes := uintBytes(int(field.Size), v)
	maxBits := uint16(len(fieldBytes) * 8)
	bitLoc := maxBits - field.Size
	curByteLoc := bitLoc / 8
	curByteBitLoc := bitLoc - curByteLoc*8
	for bitLoc < maxBits {
		bitValue := (fieldBytes[curByteLoc] & (1 << (7 - curByteBitLoc))) >> (7 - curByteBitLoc)
		if bitValue == 1 {
			ret = append(ret, true)
		} else {
			ret = append(ret, false)
		}
		bitLoc++
		curByteBitLoc++
		if curByteBitLoc >= 8 {
			curByteLoc++
			curByteBitLoc = 0
		}
	}
	return
}

func (schema *Schema) allAttrMask() (mask uint16) {
	mask = 0
	for count := uint8(0); count < uint8(len(schema.Me.Attributes.Attribute)); count++ {
		mask |= (1 << (15 - count))
	}
	return mask
}

func (schema *Schema) sbcAttrMask() (mask uint16) {
	mask = 0
	for count := uint8(0); count < uint8(len(schema.Me.Attributes.Attribute)); count++ {
		if schema.Me.Attributes.Attribute[count].isSBC() {
			mask |= (1 << (15 - count))
		}
	}
	return mask
}
