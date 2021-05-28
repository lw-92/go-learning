package ch5

import (
	"fmt"
	"runtime"
	"testing"
)

/**
case 不需要break,
 和java 一样，不限制类型
*/
func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	default:
		fmt.Print("default")
	}
}

/**
可以支持多个值
*/
func TestSwithMultiCase(t *testing.T) {

	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("even")
		case 1, 3:
			t.Log("odd")

		}
	}
}

/**
支持条件 做为switch 条件，这样和if就没有什么区别
*/
func TestSwithMultiCaseCondition(t *testing.T) {

	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("even")
		case i%2 == 1:
			t.Log("odd")

		}
	}
}
