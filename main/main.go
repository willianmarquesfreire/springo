package main

import (
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func main()  {
	fmt.Println(bson.NewObjectId())
}