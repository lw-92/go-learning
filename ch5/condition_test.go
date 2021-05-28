package ch5

import (
	"errors"
	"testing"
)

/**
和其他语言不一样的地方在与，可以在if 中赋值
*/
func TestCondition(t *testing.T) {
	if a := 1; a < 2 {
		t.Log("T%", a)
	}
}

/**
  主要作用,多返回值的情况下，可以直接判断某一个返回值
*/
func TestIfMultiSec(t *testing.T) {
	if v, erre := someFunc(); erre == nil {
		t.Logf("v is %d", v)
	} else {
		t.Logf("has error v is %d", v)
	}
}

func someFunc() (int, error) {

	return 1, errors.New("test")
}
