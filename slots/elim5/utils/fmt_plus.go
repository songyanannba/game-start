package utils

import (
	"fmt"
	"reflect"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: StructToMap
//@description: 利用反射将结构体转化为map
//@param: obj interface{}
//@return: map[string]interface{}

func StructToMap(obj interface{}) map[string]interface{} {
	obj2 := reflect.ValueOf(obj)
	if obj2.Kind() == reflect.Ptr {
		obj2 = obj2.Elem()
	}
	obj1 := obj2.Type()

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

// StructToSpreadMap 接口体转展开的map
func StructToSpreadMap(obj interface{}) map[string]interface{} {
	objValue := reflect.ValueOf(obj)

	// Dereference if it's a pointer
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	data := make(map[string]interface{})
	for i := 0; i < objType.NumField(); i++ {
		fieldType := objType.Field(i)
		fieldValue := objValue.Field(i)

		// Check if the field is an embedded struct
		if fieldType.Anonymous && fieldValue.Kind() == reflect.Struct {
			// Recursively convert the nested struct to a map
			for k, v := range StructToSpreadMap(fieldValue.Interface()) {
				data[k] = v
			}
		} else {
			key := fieldType.Name
			if jsonTag := fieldType.Tag.Get("json"); jsonTag != "" {
				key = jsonTag
			}
			data[key] = fieldValue.Interface()
		}
	}
	return data
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ArrayToString
//@description: 将数组格式化为字符串
//@param: array []interface{}
//@return: string

func ArrayToString(array []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
