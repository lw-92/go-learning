package main

/**
给定一个不含重复元素的整数数组 nums 。一个以此数组直接递归构建的 最大二叉树 定义如下：

二叉树的根是数组 nums 中的最大元素。
左子树是通过数组中 最大值左边部分 递归构造出的最大二叉树。
右子树是通过数组中 最大值右边部分 递归构造出的最大二叉树。
返回有给定数组 nums 构建的 最大二叉树 。

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if nums == nil {
		return nil
	}
	// find max and index
	idex, maxVal := 0, 0
	for idx, val := range nums {
		if val >= maxVal {
			maxVal = val
			idex = idx
		}
	}
	root := new(TreeNode)
	root.Val = maxVal
	root.Left = constructMaximumBinaryTree(nums[0, idex])
root.Right = constructMaximumBinaryTree(nums[idex, len(nums)])

return root

}
