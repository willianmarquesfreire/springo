package reflect

import (
	"strings"
	"reflect"
)

func IsPointer(obj interface{}) bool {
	return strings.Contains("*", reflect.TypeOf(obj).String())
}
