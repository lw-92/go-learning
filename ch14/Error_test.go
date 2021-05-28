package ch14

import (
	"errors"
	"testing"
)

/**
go 中的异常处理机制
在嵌套处理的时候，一般判断error 是null 的时候做后续的处理
*/
func TestError(t *testing.T) {
	if thing, err := dosomeThing(-1); err != nil {
		if err == LessThanZeroError {
			t.Log("less 0")
		}
		t.Log(thing)
	} else {
		t.Log(err)
	}
}

// 预先定义好错误的类型
var LessThanZeroError = errors.New("not be less than 0")
var MoreThanHundrundError = errors.New("more than 100")

/**
如何区分错误类型
*/
func dosomeThing(input int) (int, error) {
	if input < 0 {
		return 0, LessThanZeroError
	}
	if input > 100 {
		return 0, MoreThanHundrundError
	}
	return input * 10, nil
}
