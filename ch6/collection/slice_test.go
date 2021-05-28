package collection

import "testing"

/**
切片
内部是一个结构体
 ptr 一个指针，指向内部的一个数组
len 元素的个数
cap 内部数组的容量
切片 不能做比较，只能和空做比较
*/
func TestSlice(t *testing.T) {
	var s0 []int            //不需要说明长度
	t.Log(len(s0), cap(s0)) //容量和长度都是0
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) //容量和长度都是1
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2)) //长度和容量
	//t.Log(s2[0],s2[1],s2[2],s2[3])// 只初始化了3个，不能访问下标为3的元素
	t.Log(s2[0], s2[1], s2[2]) //
	s2 = append(s2, 2)         //append 会直接加入到后面 0002
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])

}

/**
2倍扩展
*/
func TestGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i) // 存储空间扩容时,存储空间的地址有变化，所以需要重新赋值
		t.Log(len(s), cap(s), s)
	}

}

//func TestCompareSlice(t *testing.T) {
//	a := []int{1, 2, 3, 4}
//	b := []int{1, 2, 3, 4}
//	t.Log(a == b) // 是不能比较的
//}

/**
共享存储空间的测试
*/
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) //为什么cap 是9 ，因为 截取的原来的空间，后续的空间是不变的
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	// 这里修改summer的数据,Q2的数据也会改变，因为是共享空间
	summer[0] = "Unknown"
	t.Log(Q2)

}
