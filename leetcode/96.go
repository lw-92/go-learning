package main

import (
	"fmt"
)

/**
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。



示例 1：


输入：n = 3
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/**
穷举：
  依次用每个节点做为跟节点， 左右的积就是根节点的树的总数，超时了
  需要用备忘录记住,因为超过循环过数字的一半之后，会重复计算之前的值(左右调换了)
*/

var cache = make(map[string]int) //全局变量要用car

func numTrees(n int) int {
	return countFc(1, n)
}

/**
但是超出时间限
*/
func countFc(lo int, hi int) int {
	if lo > hi {
		return 1 // null 也算是一种情况
	}

	if val, ok := cache[fmt.Sprintf("%d-%d", lo, hi)]; ok {
		return val
	} else {
		result := 0 // 这里需要限制是lo
		for i := lo; i <= hi; i++ {
			countLeft := countFc(lo, i-1)
			countRight := countFc(i+1, hi)
			result += countLeft * countRight
		}
		cache[fmt.Sprintf("%d-%d", lo, hi)] = result
		return result
	}

}
