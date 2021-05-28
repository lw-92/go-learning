package ch15

import (
	"learning/common" // 为什么是learning 下的common
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
