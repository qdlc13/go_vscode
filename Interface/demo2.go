//go:build ignore
// +build ignore

package main

import "fmt"

type AInterface interface {
	a()
}
type BInterface interface {
	b()
}

//接口组合 要实现c接口必须把AB接口的方法全部实现
type CInterface interface {
	AInterface
	BInterface
	c()
}

type EInterface interface {
	e()
}
type Stu struct {
}

func (s Stu) a() {
	fmt.Println("aaa")
}
func (s Stu) b() {
	fmt.Println("bbb")
}
func (s Stu) c() {
	fmt.Println("ccc")
}

func main() {
	//一个自定义类型可以实现多个接口
	var s Stu
	var a AInterface = s
	var b BInterface = s
	var c CInterface = s
	a.a()
	b.b()
	c.a()
	c.b()
	c.c()
	//interface类型默认是一个指针(引用类型)，如果没有对interface初始化就使用,那么会输出nil
	fmt.Println(c) //{}
	fmt.Println(s) //{}

	//interface类型默认是一个指针(引用类型)，如果没有对interface初始化就使用,那么会输出nil

	//空接口没有任何方法,所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口。
	var e EInterface
	fmt.Println(e) //<nil>

	var e2 interface{} = s
	fmt.Println(e2) //{}
	var num float64 = 3.2
	var e3 interface{} = num
	fmt.Println(e3) //3.2

}
