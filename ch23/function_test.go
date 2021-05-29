package ch23

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 go test -v -cover 可以看覆盖率
Fail Error 该测试失败，该测试继续，其他测试继续执行
FailNow，Fatal 该测试失败，该测试终止，其他测试继续执行
*/
func TestSquare(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := sqare(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected is %d,the actual is %d", inputs[i], expected[i], ret)
		}
	}
}

func sqare(i int) int {
	return i * i
}

/**
Fatal 不会继续后面代码执行
*/
func TestFatalInCode(t *testing.T) {
	fmt.Print("start")
	t.Fatal("error")
	fmt.Println("end")
}

/**
https://github.com/stretchr/testify
*/
func TestWithTestify(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
}
