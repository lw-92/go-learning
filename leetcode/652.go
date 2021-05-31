package main

import (
	"fmt"
)

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
	var nodes []*TreeNode
	buildSubTree(root, allSubtreeMapping, nodes)
	return nodes
}

/**
意思是，一开始也是写作 *map[key]value，后来发现所有的map都是当作指针来用的，于是就省略简写了。
*/
func buildSubTree(node *TreeNode, allSubtreeMapping map[string]int, nodes []*TreeNode) string {
	if node == nil {
		return "#"
	}
	sprintf := fmt.Sprintf("%d,%s,%s", node.Val, buildSubTree(node.Left, allSubtreeMapping, nodes), buildSubTree(node.Right, allSubtreeMapping, nodes))

	if v, ok := allSubtreeMapping[sprintf]; ok {
		allSubtreeMapping[sprintf] = v + 1
		_ = append(nodes, node)
	} else {
		allSubtreeMapping[sprintf] = 1
	}
	return sprintf
}
