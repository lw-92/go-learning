package ch20

import (
	"fmt"
	"sync"
	"testing"
)

/**
多个消费者的情况下，消费者不知道什么情况下会结束，需要close channel 来通知所有的
消费者
向关闭的channel 发送数据时，会导致panic
*/
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok {
				fmt.Print(data)
			} else {
				fmt.Println("channel close")
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg) //这里为什么要传指针
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
	t.Log("main end")
}
