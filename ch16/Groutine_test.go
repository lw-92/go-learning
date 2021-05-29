package ch16

import (
	"testing"
	"time"
)

/**
groutine 的堆栈更小，默认只有2k, thread 是 1M
*/
func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			t.Log(i)
		}(i)
	}
	t.Log("main completed")
	time.Sleep(time.Millisecond * 10)
}
