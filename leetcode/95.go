package main

/**
给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。



示例 1：


输入：n = 3
输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
示例 2：

输入：n = 1
输出：[[1]]


提示：

1 <= n <= 8

*/

func generateTrees(n int) []*TreeNode {

	if n == 0 {
		return nil
	}
	var helper func(start int, end int) []*TreeNode
	helper = func(start int, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		var alltrees []*TreeNode
		//枚举所有的跟节点
		for i := start; i <= end; i++ {
			//左边的树
			leftTree := helper(start, i-1)
			////右边的树
			rightTree := helper(i+1, end)
			//// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
			for _, left := range leftTree {
				for _, right := range rightTree {
					currentTree := new(TreeNode)
					currentTree.Left = left
					currentTree.Right = right
					currentTree.Val = i
					alltrees = append(alltrees, currentTree)
				}
			}
		}
		return alltrees

	}
	return helper(1, n)
}
