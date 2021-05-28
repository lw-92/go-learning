package ch14

import (
	"errors"
	"fmt"
	"testing"
)

/**
panic vs os.Exit
os.Exit 退出不会调用defer &&调用栈信息
*/
func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Print("defer")
	}()
	fmt.Println("start")
	//os.Exit(-1)// 没有堆栈信息
	panic(errors.New("error happened")) //会有堆栈，以及derfer 也执行了
}

/**
recover ,
算是一种catch 机制，让程序能继续进行下去，
*/
func TestRecover(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			//可以在这里区分具体的错误类型
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("start")
	//os.Exit(-1)// 没有堆栈信息
	panic(errors.New("error happened"))
}
