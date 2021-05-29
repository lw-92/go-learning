package ch25

import (
	"fmt"
	"reflect"
	"testing"
)

/**
go 中反射
 reflect.TypeOf 返回类型(reflect.Type)
 reflect.ValueOf 返回值（reflect.Value)
 可以从reflect.Value 获得类型
 通过Kind来判断类型 ，kind 是枚举类型
*/

func CheckType(v interface{}) {
	of := reflect.TypeOf(v)
	switch of.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Print("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	default:
		fmt.Println("unknown", of)

	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)

	t.Log("done")
}

func TestTyoeAndValue(t *testing.T) {
	var i int64 = 12
	t.Log(reflect.TypeOf(i), reflect.ValueOf(i))
	t.Log(reflect.ValueOf(i).Type())
}
