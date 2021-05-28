package ch15

import (
	cm "github.com/easierway/concurrent_map" // 从git hub 上导包
	"learning/common"
	"testing"
)

/**
基本的复用的模块单元
  以首字母大写表明可被包外代码访问
 代码的package 可以和所在的目录不一致
 同一目录里的go 代码的package 要保持一致
gopath:
 moudle:
*/
func TestP(t *testing.T) {
	common.DoCommonThings()
	t.Log("complete")
}

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
