package ch18

import (
	"fmt"
	"testing"
	"time"
)

/**
csp 机制,通过共享通道解决共享数据的问题
*/

func Service() string {
	time.Sleep(10 * time.Millisecond)
	return "Done"
}

func otherTask() {
	fmt.Print("Working on something else")
	time.Sleep(time.Millisecond * 100)
}

/**
并行的执行
*/
func TestService(t *testing.T) {
	fmt.Println(Service())
	otherTask()
	t.Log("all task done")
}

func TestAsyncService(t *testing.T) {
	service := AsyncService()
	otherTask()
	fmt.Println(<-service) //这里会阻塞等待
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

func AsyncServiceWithBuffer() chan string {
	retCh := make(chan string, 1) //带有一个buffer,
	go func() {
		service := Service()
		fmt.Println("returned result")
		retCh <- service //有buffer 不会阻塞在这里
		fmt.Println("service exited")

	}()
	return retCh
}

func TestAsyncWithBuffer(t *testing.T) {
	service := AsyncServiceWithBuffer()
	otherTask()
	fmt.Println(<-service) //这里会阻塞等待
}
