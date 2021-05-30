package main

/**
根据一棵树的中序遍历与后序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

中序遍历 inorder = [9,3,15,20,7]
后序遍历 postorder = [9,15,7,20,3]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
  中序遍历： 左子树的中序    root     右子树的中序
  后续遍历： 左子树的后序，右子树的后序，root
*/
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if inorder == nil || len(inorder) == 0 {
		return nil
	}
	//构造一个映射关系
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}
	node := new(TreeNode)
	node.Val = postorder[len(postorder)-1]
	node.Left = buildTree1(inorder[:i], postorder[:i])
	node.Right = buildTree1(inorder[i+1:], postorder[i:len(postorder)-1])
	return node
}
