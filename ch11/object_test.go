package ch11

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
面向对象
 不支持继承
new 返回的是指向实例的指针
*/
func TestCreateEmployee(t *testing.T) {
	e := new(Employee)
	e.Age = 10
	e.Id = "1"
	e.Name = "liwei"
	employee := Employee{"0", "liwei1", 10}

	e2 := Employee{Name: "liwei3", Age: 30}
	t.Log(e, e2, employee) // 指针类型

}

/**
对象的定义
 添加行为、方法
*/
type Employee struct {
	Id   string
	Name string
	Age  int
}

/**
这两种的区别
这种会复制一个,会有跟多的内存开销
*/
func (e Employee) String() string {
	fmt.Printf("adress is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("id:%s-name :%s-age:%d", e.Id, e.Name, e.Age)
}

/**
不会复制对象
*/
func (e *Employee) String1() string {
	fmt.Printf("adress is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("id:%s-name :%s-age:%d", e.Id, e.Name, e.Age)
}

func (e *Employee) change() string {
	e.Id = "new id"
	return fmt.Sprintf("id:%s-name :%s-age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	employee := Employee{"0", "liwei", 20}
	fmt.Printf("1111adress is %x", unsafe.Pointer(&employee.Name))
	t.Log(employee.String())
	fmt.Println("===============================")
	// 使用指针访问,是可以调用的
	e := &Employee{"0", "liwei", 20}
	fmt.Printf("adress is %x", unsafe.Pointer(&e.Name))
	t.Log(e.String1())
	// 指针的值会改变
	e.change()
	fmt.Println("after change", e)
	employee.change() // 对象也会改变
	fmt.Println("employedd change", employee)
}
