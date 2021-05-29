package ch17

import (
	"sync"
	"testing"
	"time"
)

/**
共享内存.
得不到1000
*/
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Log("counter is", counter)
}

/**
安全 使用锁
*/
func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Log("counter is", counter)
}

/**
类似java 中的join

*/
func TestWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("counter is", counter)

}
