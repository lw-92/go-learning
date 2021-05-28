package ch5

import "testing"

/**
go 中的循环

*/
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
