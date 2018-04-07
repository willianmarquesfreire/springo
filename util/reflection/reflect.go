package reflection

import (
	"strings"
	"reflect"
	"springo/domain"
	"errors"
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

func Validate(obj interface{}) error {
	var erros string = ""
	s := GetElem(obj)

	typeof := s.Type()

	for i := 0; i < typeof.NumField(); i++ {
		if strings.Contains(typeof.Field(i).Tag.Get("validate"), "required") && (s.Field(i).Interface() == nil || s.Field(i).Interface() == "" || s.Field(i).Interface() == 0) {
			erros = erros + typeof.Field(i).Name + " has empty, "
		}
	}

	if erros == "" {
		return nil
	}
	return errors.New(erros)
}

func GetAttrs(obj interface{}) []Attr {
	var attrs []Attr = make([]Attr, 0)
	s := GetElem(obj)

	typeof := s.Type()

	for i := 0; i < typeof.NumField(); i++ {
		if !Equal(s.Field(i).Interface(), domain.GenericDomain{}) && s.Field(i).Interface() != nil && s.Field(i).Interface() != "" {
			f := Attr{
				Field: typeof.Field(i),
			}
			f.Value = s.Field(i)
			attrs = append(attrs, f)
		} else if Equal(s.Field(i).Interface(), domain.GenericDomain{}) {
			for j := 0; j < s.Field(i).NumField(); j++ {
				field := s.Field(i).Field(j)
				if field.Interface() != nil && field.Interface() != "" {
					f := Attr{
						Field: s.Field(i).Type().Field(j),
					}
					f.Value = s.Field(i).Field(j)
					attrs = append(attrs, f)
				}
			}
		}
	}
	return attrs
}
