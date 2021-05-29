package ch25

import (
	"reflect"
	"testing"
)

/**
map 和slice 不能==，就需要用deepequal
*/
func TestDeepEqual(t *testing.T) {
	m1 := map[int]string{1: "one", 2: "two"}
	m2 := map[int]string{1: "one", 2: "two"}
	t.Log(reflect.DeepEqual(m1, m2))
}
