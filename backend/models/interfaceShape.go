package models

import (
	"encoding/json"
	"reflect"
)

func GetContentFromInterface(src interface{}, dest interface{}) (e error) {
	arr, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(arr, dest)
	if err != nil {
		return err
	}
	return nil
}

func MapFormToStruct(values map[string][]string, out interface{}) {
	v := reflect.ValueOf(out).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		key := field.Name

		if valArr, ok := values[key]; ok && len(valArr) > 0 {
			v.Field(i).SetString(valArr[0])
		}
	}
}
