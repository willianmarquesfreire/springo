package reflection

import (
	"strings"
	"reflect"
	"springo/domain"
)

type Attr struct {
	Field reflect.StructField
	Value reflect.Value
}

func IsError(obj interface{}) bool {
	_, b := reflect.TypeOf(obj).MethodByName("Error")
	return b
}

func IsPointer(obj interface{}) bool {
	return strings.Contains(reflect.TypeOf(obj).String(), "*")
}

func GetElem(obj interface{}) reflect.Value {
	if IsPointer(obj) {
		return reflect.ValueOf(obj).Elem()
	}
	return reflect.ValueOf(obj)
}

func Equal(obj interface{}, obj2 interface{}) bool {
	return reflect.TypeOf(obj) == reflect.TypeOf(obj2)
}

func GetAttrs(obj interface{}) []Attr {
	var attrs []Attr = make([]Attr, 0)
	s := GetElem(obj)

	typeof := s.Type()

	for i := 0; i < s.NumField(); i++ {
		if !Equal(s.Field(i).Interface(), domain.GenericDomain{}) && s.Field(i).Interface() != nil && s.Field(i).Interface() != "" {
			f := Attr{
				Field: typeof.Field(i),
			}
			f.Value = s.Field(i)
			attrs = append(attrs, f)
		}
	}
	return attrs
}
