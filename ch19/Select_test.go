package ch19

import (
	"fmt"
	"testing"
	"time"
)

/**
多路选择器

多个case 多个buffer 如何处理
*/

func TestService(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 130): // 超时控制
		t.Error("time out")
	}
}

func Service() string {
	time.Sleep(50 * time.Millisecond)
	return "Done"
}
func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		service := Service()
		fmt.Println("returned result")
		retCh <- service //没有buffer 时就阻塞在这里
		fmt.Println("service exited")

	}()
	return retCh
}
