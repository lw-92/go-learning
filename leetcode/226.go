package main

/**
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1
备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
递归交换子树
这里需要判断空指针
*/
func invertTree(root *TreeNode) *TreeNode {
	// 要注意这里的npe
	if root == nil {
		return root
	}
	if root.Left != nil {
		invertTree(root.Left)
	}
	if root.Right != nil {
		invertTree(root.Right)
	}
	root.Left, root.Right = root.Right, root.Left
	return root
}
