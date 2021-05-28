package ch9

import (
	"strconv"
	"strings"
	"testing"
)

/**
string 是数据类型，不是引用或指针
string 是只读的byte slice,len 函数可以是byte 数，不是字符数
rune(s) 可以取出string 得unicode
unicode 是规范，utf-8 是实现
https://blog.csdn.net/zhusongziye/article/details/84261211
*/
func TestInit(t *testing.T) {
	s := "中"
	t.Log(len(s))  // 这里不是字符长度，而是byte 的长度
	r := []rune(s) //rune(s) 可以取出string 得unicode  %x 代表16进制
	t.Logf("中的unicode %x", r[0])
	t.Logf("中的UTF-8 %x", s)

}

/**
1、%d表示按整型数据的实际长度输出数据。

2、%c用来输出一个字符。
*/
func TestStringToRune(t *testing.T) {

	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c) //%[1]c %[1]d 都取用第一个参数，用c 和d 格式化
	}
}

func TestStrFunc(t *testing.T) {
	s := "A,B,C"
	split := strings.Split(s, ",")
	// 第一参数位error
	for _, part := range split {
		t.Log(part)
	}
	join := strings.Join(split, "=")
	t.Log(join)

}
func TestStrToInt(t *testing.T) {
	i := 10
	s := strconv.Itoa(i)
	t.Log(s)
	atoi, err := strconv.Atoi("10")
	if err == nil {
		t.Log(10 + atoi)
	}

}
