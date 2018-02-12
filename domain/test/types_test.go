package test

import (
	"testing"
	"springo/domain"
	"fmt"
	"reflect"
	"github.com/stretchr/testify/assert"
)

func TestTypesDomainGeneric(t *testing.T) {
	d := domain.GenericDomain{Extra:"nenhum"}
	fmt.Println(d.ID)
	d.ChangeId()
	assert.Equal(t, "nenhum", d.Extra)
	assert.NotNil(t, d.ID)
	assert.IsType(t, domain.GenericDomain{}, d)

	e := domain.User{}
	e.Extra = "nenhum2"
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(e.ID)
	e.ChangeId()
	assert.Equal(t, "nenhum2", e.Extra)
	assert.NotNil(t, e.ID)
	assert.IsType(t, domain.User{}, e)


	var f interface{} = e
	assert.IsType(t, domain.User{}, f)
}