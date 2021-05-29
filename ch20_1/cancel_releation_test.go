package ch20_1

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/**
取消关联任务
根Context 通过context.Background()创建
子Context  :context.WithCancel(parentContext)
  cxt,cancel:=context.withCancel(context.Background())
当前的context 被取消时，基于他的子context 都会被取消
接收取消通知 <-cxt.Done()
*/
func TestCancelRelation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			// while 一直循环
			for {
				if IsCancelled(ctx) {
					break
				}
				fmt.Println("running")
				time.Sleep(time.Millisecond * 5)

			}
			fmt.Println(i, " Cancelled")
		}(i, ctx)
	}
	cancel()
}

func IsCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}

}
