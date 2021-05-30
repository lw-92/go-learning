package main

/**
给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：

struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。



进阶：

你只能使用常量级额外空间。
使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。


示例：



输入：root = [1,2,3,4,5,6,7]
输出：[1,#,2,3,#,4,5,6,7,#]
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
解题思路：
  拆解为两种情况, left 直接指向right  right 要指向 parent.parent.right.left
   这里要help 这个函数,作用，不要一次全部把子节点的连接做完，还是要拆分的很细
 用辅助函数,  直接给到两个节点
*/
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	helpConnect(root.Left, root.Right)
	return root
}

func helpConnect(node1 *Node, node2 *Node) {
	if node2 == nil || node1 == nil {
		return
	}
	node1.Next = node2
	//// 连接相同父节点的两个子节点
	helpConnect(node1.Left, node1.Right)
	helpConnect(node2.Left, node2.Right)
	//// 连接跨越父节点的两个子节点
	helpConnect(node1.Right, node2.Left)
}
