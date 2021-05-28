package ch12

import (
	"fmt"
	"testing"
	"time"
)

/**\
接口变量
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

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	p = new(JavaProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 自定义类型,用来替换一些复杂的类型，比如函数作为参数

/**
计算某一个函数的执行时间
*/
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		now := time.Now()
		i := inner(n)
		fmt.Println("time spend:", time.Since(now).Seconds())
		return i
	}
}

/**
模拟慢函数
*/
func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestCustomerType(t *testing.T) {
	spent := timeSpent(slowFunc)
	spent(1)
}
