package main

import "fmt"

/**
给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。

两棵树重复是指它们具有相同的结构以及相同的结点值。

示例 1：

        1
       / \
      2   3
     /   / \
    4   2   4
       /
      4
下面是两个重复的子树：

      2
     /
    4
和

    4
因此，你需要以列表的形式返回上述重复子树的根结点。

*/
/**
暴力破解法
   1
  / \
 2   3
    / \
   4   5
例如上面这棵树序列化结果为 1,2,#,#,3,4,#,#,5,#,#。每棵不同子树的序列化结果都是唯一的。

算法

使用深度优先搜索，其中递归函数返回当前子树的序列化结果。把每个节点开始的子树序列化结果保存在 map 中，然后判断是否存在重复的子树。

*/

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	allSubtreeMapping := make(map[string]int)

}

func buildSubTree(node *TreeNode, allSubtreeMapping *map[string]int) string {
	if node == nil {
		return "#"
	}
	sprintf := fmt.Sprintf("%d,%s,%s", node.Val, buildSubTree(node.Left, allSubtreeMapping), buildSubTree(node.Right, allSubtreeMapping))
	if v, ok := allSubtreeMapping[sprintf]; ok {
		allSubtreeMapping[sprintf] = v + 1
	}
}
