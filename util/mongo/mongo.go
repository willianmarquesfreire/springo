package mongo

import (
	"gopkg.in/mgo.v2/bson"
	"springo/util/reflection"
)

func GetBsonMSet(obj interface{}) bson.M {
	var attrs []reflection.Attr = reflection.GetAttrs(obj)
	bsons := make(map[string]interface{})
	for _, e := range attrs {
		fieldName := e.Field.Name
		if e.Field.Tag.Get("json") != "" {
			fieldName = e.Field.Tag.Get("json")
		}
		bsons[fieldName] = e.Value.Interface()
	}
	return bsons
}
