package ch10

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
可以有多个返回值
 所有参数都是值传递，slice,map,channel 会有引用传递的错觉，因为操作的都是结构体中的指针
函数可以作为变量的值
函数可以作为参数和返回值
*/
func TestMultiValues(t *testing.T) {
	i, i2 := multivalues()
	t.Log(i, i2)
}

/**
多返回值
*/
func multivalues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

/**
计算某一个函数的执行时间
*/
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		now := time.Now()
		i := inner(n)
		fmt.Println("time spend:", time.Since(now).Seconds())
		return i
	}
}

/**
模拟慢函数
*/
func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	spent := timeSpent(slowFunc)
	i := spent(10)
	t.Log(i)
	t.Log(sum(1, 2, 3))
	t.Log(sum(3, 4, 5, 7))
}

/**
可变长参数
*/
func sum(ops ...int) int {
	s := 0
	for _, op := range ops {
		s += op
	}
	return s
}

/**
函数 延迟执行
*/
func TestDeferFunc(t *testing.T) {
	// 相当于finally
	defer func() {
		t.Log("clear resource")
	}()

	t.Log(sum(1, 2, 3))
	panic("error") // 抛出错误
}
