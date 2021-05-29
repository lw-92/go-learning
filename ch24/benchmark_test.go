package ch24

import (
	"bytes"
	"testing"
)

/**
性能测试
*/
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2"}
	b.ResetTimer()
	var buf bytes.Buffer
	for _, ele := range elems {
		buf.WriteString(ele)
	}
	b.StopTimer()
}

func BenchmarkConcatStringByAdd2(b *testing.B) {
	elems := []string{"1", "2"}
	b.ResetTimer()
	var buf string
	for _, ele := range elems {
		buf += ele
	}
	b.StopTimer()
}
