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
 reflect.ValueOf(*e).FiledByName(name)// 获取某个成员的值
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

type Employee struct {
	Id   string
	Name string `for:"normal""`
	Age  int
}

func (e *Employee) UpdateAge(age int) {
	e.Age = age
}

/**
&符号的意思是对变量取地址
*符号的意思是对指针取值
*/
func TestInvokeByName(t *testing.T) {
	employee := &Employee{"1", "mike", 30}
	// 为什么这里要用*
	t.Log("Name=", reflect.ValueOf(*employee).FieldByName("Name"))
	reflect.ValueOf(employee).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(100)})
	t.Log(employee)
	if name, b := reflect.TypeOf(*employee).FieldByName("Name"); b {
		t.Log("tag:for=", name.Tag.Get("for"))
	}
}

/**
获取struct tah
*/
func TestTag(t *testing.T) {

}
