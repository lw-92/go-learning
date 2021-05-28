package operator_test

import "testing"

/**
数组的比较
1. 数组长度不相等的，不能==，会编译错误
*/
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4, 5}
	c := [...]int{1, 2, 3, 4}
	//t.Log(a==b) 数组长度不等，不能做比较
	t.Log(b)
	t.Log(a == c)
}
