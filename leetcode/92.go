package main

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
用一个计数器，遍历加1，
*/
var count = 1

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	currentNode := head
	for {
		if currentNode.Next != nil {

		} else {

		}
	}

}
