//go:build ignore
// +build ignore

package main

import "fmt"

//接口
type say interface {
	say()
}

//自定义数据类型结构体
type Chinese struct {
}

//自定义数据类型
type integer int

func (i integer) say() {
	fmt.Println("say hi = ", i)
}

//实现接口的方法
func (person Chinese) say() {
	fmt.Println("你好")
}

type American struct {
}

func (person American) say() {
	fmt.Println("hello")
}

func greet(s say) {
	s.say()
}

func main() {

	c := Chinese{}
	a := American{}

	var i integer = 20

	greet(i)
	greet(a)
	greet(c)

	// var s say 接口本身不能实例化，但是可以指向实现该接口的自定义类型变量
	// s.say()
	var s say
	s = c
	s.say()
	greet(s)

	//多态是使用不同接口实现的 多态数组
	var arr [3]say
	arr[0] = American{}
	arr[1] = Chinese{}
	arr[2] = American{}
	fmt.Println(arr)

}

/*
（1）接口中可以定义一组方法，但不需要实现，不需要方法体。并且接口中不能包含任何变量。到某个自定义类型要使用的时候（实现接口的时候）,再根据具体情况把这些方法具体实现出来。
（2）实现接口要实现所有的方法才是实现。
（3）Golang中的接口不需要显式的实现接口。Golang中没有implement关键字。
（Golang中实现接口是基于方法的，不是基于接口的）
例如：
A接口 a,b方法
B接口 a,b方法
C结构体 实现了  a,b方法 ，那么C实现了A接口，也可以说实现了B接口   （只要实现全部方法即可，和实际接口耦合性很低，比Java松散得多）
（4）接口目的是为了定义规范，具体由别人来实现即可。
*/
