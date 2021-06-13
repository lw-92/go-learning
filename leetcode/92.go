package main

/**
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
用一个计数器，遍历加1，   head next
这道题的递归写法可以用上面的模板思路
1.首先做一个临界条件判断，如果left>1表示没有到达翻转的起点
2.当到达起点后，就通过判断right的值来确定翻转部分的右边界
3.写法与上面的代码部分差不多，主要区别在于用temp记录了子问题部分的next（也就是翻转部分的边界的指向）

作者：505693903
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/bi-jiao-jian-dan-de-di-gui-jie-fa-by-505-p34y/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
这是反转链表的递归写法
var reverseFunc func(head *ListNode) *ListNode
	reverseFunc = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		next := reverseFunc(head.Next)
		head.Next.Next = head
		head.Next = nil
		return next
	}

解题思路为给出m和n，我们分别找到第m-1和n+1个节点，然后让第m-1个节点连在n的前面，第m个节点连在第n+1个节点的前面，就完成了反转，
此时要注意m如果等于1的情况，即m为头借点，这时直接让第m个节点连在第n+1个节点前面就完成了任务。
https://labuladong.gitee.io/algo/2/17/16/

*/
var successor *ListNode = nil // 后驱节点

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	/**
	反转以 head 为起点的 n 个节点，返回新的头结点
	*/
	var reverseN func(head *ListNode, n int) *ListNode

	reverseN = func(head *ListNode, n int) *ListNode {

		if n == 1 {
			// 记录第 n + 1 个节点
			successor = head.Next
			return head
		}
		last := reverseN(head.Next, n-1)
		head.Next.Next = head
		head.Next = successor
		return last

	}
	if left == 1 {
		// 相当于反转前 n 个元素
		return reverseN(head, right)
	}
	/**
	如果 m != 1 怎么办？如果我们把 head 的索引视为 1，那么我们是想从第 m 个元素开始反转对吧；
	如果把 head.next 的索引视为 1 呢？那么相对于 head.next，反转的区间应该是从第 m - 1 个元素开始的；那么对于 head.next.next 呢……
	*/
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head

}
