package ch13

import (
	"fmt"
	"testing"
)

/**
go 中多态的支持
*/

/**
接口变量
、
*/

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (input *GoProgrammer) WriteHelloWorld() string {
	return "fmt.println(\"hello world\")"
}

type JavaProgrammer struct {
}

func (input *JavaProgrammer) WriteHelloWorld() string {
	return "sout(hello word)"
}

func writeFisrtProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

/**
多肽测试
*/
func TestPolymorphism(t *testing.T) {

	goP := new(GoProgrammer)

	javaP := new(JavaProgrammer)
	writeFisrtProgram(goP)
	writeFisrtProgram(javaP)
}

// 空接口和断言
/**
通过断言江空接口转换为制定类型
*/

func TestVoidInterface(t *testing.T) {
	doSomeThing(2)
	doSomeThing("33333")

}

func doSomeThing(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println("Integer", i)
	//	return
	//}
	//if i, ok := p.(string); ok {
	//	fmt.Println("string", i)
	//	return
	//}
	//fmt.Println("unknown")
	//另外一种写法
	switch i := p.(type) {
	case string:
		fmt.Println("string", i)
	case int:
		fmt.Println("Integer", i)
	default:
		fmt.Println("unknown")
	}
}
