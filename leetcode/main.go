package main

import "fmt"

func main() {
	// 654
	//tree := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	//fmt.Print(tree)
	//105
	//tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	//fmt.Print(tree)
	//tree := buildTree1([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	//fmt.Print(tree)
	trees := numTrees(3)
	fmt.Print(trees)
}
