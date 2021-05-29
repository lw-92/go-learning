package ch21

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
  任意一个任务完成即可
*/

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

/**
这里会有会又协程被堵塞再这里，是有问题的，要用带有buffer的channel
*/
func FirstResponse() string {
	numOfRuner := 10
	ch := make(chan string, numOfRuner)
	// 没有缓冲的buffer , 写
	for i := 0; i < numOfRuner; i++ {
		go func(i int) {
			task := runTask(i)
			ch <- task
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("End:", runtime.NumGoroutine())
}
