package main

/**
斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1
给你 n ，请计算 F(n)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fibonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
var fibs = make(map[int]int)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	if val, ok := fibs[n]; ok {
		return val
	} else {
		i := fib(n-1) + fib(n-2)
		fibs[n] = i
		return i
	}

}
