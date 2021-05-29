package ch21

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
只运行一次， 再java 中的单例模式
*/

type Singleton struct {
}

func TestOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			ton := getSingleTon()
			fmt.Printf("addr is %p", ton)
		}()
	}
	time.Sleep(time.Second * 1)
	t.Log("done")
}

var singletonInstance *Singleton
var once sync.Once

func getSingleTon() *Singleton {
	once.Do(func() {
		fmt.Print("create instance")
		singletonInstance = new(Singleton)
	})
	return singletonInstance
}
