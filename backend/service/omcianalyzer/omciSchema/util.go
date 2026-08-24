/*
# ------------------------------------------------------------
# -- meschema.go
# --
# -- Cao Minghui
# -- 2022-7-11
# ------------------------------------------------------------
*/

package omciSchema

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"strconv"
)

func bytesUint16(bytes []byte) uint16 {
	//return uint16(bytes[0])<<8 | uint16(bytes[1])
	return binary.BigEndian.Uint16(bytes)
}

func bytesUint64(bytes []byte) uint64 {
	//return uint16(bytes[0])<<8 | uint16(bytes[1])
	return binary.BigEndian.Uint64(bytes)
}

func uint16Bytes(v uint16) (bytes []byte) {
	//bytes = (*[2]byte)(unsafe.Pointer(&v))[:]
	bytes = make([]byte, 2)
	//binary.LittleEndian.PutUint16(bytes, v)
	binary.BigEndian.PutUint16(bytes, v)
	return
}

func pad(bytesIn []byte, width uint16) (bytesOut []byte) {
	bytesOut = bytesIn
	l := uint16(len(bytesOut))
	if l < width {
		bytesOut = append(bytesOut, make([]byte, width-l)...)
	}
	return
}

func uint32Bytes(v uint32) (bytes []byte) {
	bytes = make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, v)
	return
}

func bytesUInteger(bytes []byte, size uint16) interface{} {
	if size == 1 {
		return uint8(bytes[0])
	}
	if size == 2 {
		return binary.BigEndian.Uint16(bytes)
	}
	if size == 4 {
		return binary.BigEndian.Uint32(bytes)
	}
	if size == 5 {
		return Uint40(bytes)
	}
	return binary.BigEndian.Uint64(bytes)
}
func Uint40(b []byte) uint64 {
	_ = b[4]
	paddBytes := []byte{0, 0, 0}
	return uint64(b[4]) | uint64(b[3])<<8 | uint64(b[2])<<16 | uint64(b[1])<<24 | uint64(b[0])<<32 | uint64(paddBytes[0])<<40 | uint64(paddBytes[1])<<48 | uint64(paddBytes[2])<<56
}
func interfaceArray(v interface{}) (ret []interface{}) {
	va := reflect.ValueOf(v)
	if va.Kind() == reflect.Array || va.Kind() == reflect.Slice {
		ret = make([]interface{}, va.Len())
		for count := 0; count < va.Len(); count++ {
			ret[count] = va.Index(count).Interface()
		}
	} else {
		ret = make([]interface{}, 1)
		ret[0] = v
	}
	return
}
func uint8Bytes(v uint8) (bytes []byte) {
	bytes = make([]byte, 1)
	bytes[0] = v
	return
}

func uint64Bytes(v uint64) (bytes []byte) {
	bytes = make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, v)
	return
}
func convertToFloat64(input interface{}) (ret float64, ok bool) {
	ret = 0
	ok = true
	switch input.(type) {
	case float64:
		ret, ok = input.(float64)
	case uint:
		var tmp uint
		tmp, ok = input.(uint)
		ret = float64(tmp)
	case uint64:
		var tmp uint64
		tmp, ok = input.(uint64)
		ret = float64(tmp)
	case uint32:
		var tmp uint32
		tmp, ok = input.(uint32)
		ret = float64(tmp)
	case uint16:
		var tmp uint16
		tmp, ok = input.(uint16)
		ret = float64(tmp)
	case uint8:
		var tmp uint8
		tmp, ok = input.(uint8)
		ret = float64(tmp)
	case int:
		var tmp int
		tmp, ok = input.(int)
		ret = float64(tmp)
	case int64:
		var tmp int64
		tmp, ok = input.(int64)
		ret = float64(tmp)
	case int32:
		var tmp int32
		tmp, ok = input.(int32)
		ret = float64(tmp)
	case int16:
		var tmp int16
		tmp, ok = input.(int16)
		ret = float64(tmp)
	case int8:
		var tmp int8
		tmp, ok = input.(int8)
		ret = float64(tmp)
	// case string:
	// 	if tmp, err := strconv.ParseFloat(input.(string), 64); err == nil {
	// 		ret = tmp
	// 	}
	default:
		return ret, false
	}
	return ret, ok
}

func uintsBytes(size int, v interface{}) (bytes []byte) {
	va := reflect.ValueOf(v)
	if va.Kind() == reflect.Slice {
		for count := 0; count < va.Len(); count++ {
			bytes = append(bytes, uintBytes(size, va.Index(count).Interface())...)
		}
	} else if va.Kind() == reflect.String {
		bytes = []byte(v.(string))
	} else {
		bytes = append(bytes, uintBytes(size, v)...)
	}
	return
}

func uintBytes(size int, v interface{}) (bytes []byte) {
	//Use explicit size iso interface to solve json number type
	vf := float64(0)
	switch v.(type) {
	case float64:
		vf = v.(float64)
	case uint32:
		vf = float64(v.(uint32))
	case uint16:
		vf = float64(v.(uint16))
	}
	if size <= 8 {
		bytes = uint8Bytes(uint8(vf))
	} else if size <= 16 {
		bytes = uint16Bytes(uint16(vf))
	} else if size <= 32 {
		bytes = uint32Bytes(uint32(vf))
	} else if size <= 64 {
		bytes = uint64Bytes(uint64(vf))
	}
	return
}

func bytesInteger(bytes []byte, size uint16) interface{} {
	if size == 1 {
		return int8(bytes[0])
	}
	if size == 2 {
		return int16(binary.BigEndian.Uint16(bytes))
	}
	if size == 4 {
		return int32(binary.BigEndian.Uint32(bytes))
	}
	return int64(binary.BigEndian.Uint64(bytes))
}

func stringToInt64(input string) (output int64, e error) {
	output = 0
	v, e := strconv.Atoi(input)
	if e == nil {
		output = int64(v)
	}
	return
}

const (
	DevIdBaseline byte = 0x0A
	DevIdExtended byte = 0x0B
)

// FormatDevId returns a human-readable OMCI message format label.
func FormatDevId(devId byte) string {
	switch devId {
	case DevIdBaseline:
		return "Baseline OMCI"
	case DevIdExtended:
		return "Extended OMCI"
	default:
		return fmt.Sprintf("Unknown (0x%02X)", devId)
	}
}
