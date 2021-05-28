package ch7

import (
	"testing"
)

/**
map的相关的测试
*/
func TestInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	t.Log(m, len(m))
	m2 := make(map[string]int, 10) // 这里len 还是0，因为没有放值
	t.Log(m2, len(m2))
	m3 := map[string]int{}
	m3["four"] = 4
	t.Log(m3, len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m := map[int]int{}
	t.Log(m[1])
	m[2] = 0
	t.Log(m[2]) //都是0
	//那么怎么判断值不存在了

	if v, ok := m[3]; ok {
		t.Logf("exiting 3 %d", v)
	} else {
		t.Log("not existing 3")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m {
		t.Log(k, v)
	}

}
