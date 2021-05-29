package ch20_1

import (
	"fmt"
	"testing"
	"time"
)

/**
任务的取消
通过channel 的close 来实现

*/

func TestCancel(t *testing.T) {

	cancelChan := make(chan struct{}, 0)

	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			if _, ok := <-cancelChan; ok {
				fmt.Print("running")
			} else {
				fmt.Println("canceled ", i)
			}

		}(i, cancelChan)
	}
	Cancel(cancelChan)
	time.Sleep(time.Second * 3)
}

func Cancel(ch chan struct{}) {
	time.Sleep(time.Second * 1)
	close(ch)
}
