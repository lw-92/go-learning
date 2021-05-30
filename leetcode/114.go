package main

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。


示例 1：


输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [0]
输出：[0]


提示：

树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100


进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
给 flatten 函数输入一个节点 root，那么以 root 为根的二叉树就会被拉平为一条链表。

我们再梳理一下，如何按题目要求把一棵树拉平成一条链表？很简单，以下流程：

1、将 root 的左子树和右子树拉平。

2、将 root 的右子树接到左子树下方，然后将整个左子树作为右子树。

上面三步看起来最难的应该是第一步对吧，如何把 root 的左右子树拉平？其实很简单，按照 flatten 函数的定义，对 root 的左右子树递归调用 flatten 函数即可
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	/**** 后序遍历位置 ****/
	// 1、左右子树已经被拉平成一条链表
	left := root.Left
	right := root.Right
	//// 2、将左子树作为右子树
	root.Left = nil
	root.Right = left

	// // 3、将原先的右子树接到当前右子树的末端,先找到最右的树
	node := root
	for {

		if node.Right == nil {
			break
		}
		node = node.Right
	}
	node.Right = right
}
