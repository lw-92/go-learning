package ch21

import (
	"fmt"
	"testing"
	"time"
)

/**
等待所有的任务都返回，利用channel
*/

func RunTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

/**
这里会有会又协程被堵塞再这里，是有问题的，要用带有buffer的channel
*/
func AllResponse() string {
	numOfRuner := 10
	ch := make(chan string, numOfRuner)
	// 没有缓冲的buffer , 写
	for i := 0; i < numOfRuner; i++ {
		go func(i int) {
			task := runTask(i)
			ch <- task
		}(i)
	}
	finalRes := ""
	for i := 0; i < numOfRuner; i++ {
		finalRes += <-ch + "\n"
	}
	return finalRes

}

func TestAllRes(t *testing.T) {
	response := AllResponse()
	t.Log(response)
}
