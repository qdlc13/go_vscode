//go:build ignore
// +build ignore

package main

import "fmt"

//Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok := element.(T)，
//这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
/*
value, ok := x.(T)
x表示要断言的接口变量；
T表示要断言的目标类型；
value表示断言成功之后目标类型变量；
ok表示断言的结果，是一个bool型变量，true表示断言成功，false表示失败，如果失败value的值为nil。
*/
type Say interface {
	saySomething()
}

type Chinese struct {
	name string
}
type American struct {
	name string
}

func (c Chinese) saySomething() {
	fmt.Println("你好")
}
func (a American) saySomething() {
	fmt.Println("hi")
}

//特有方法
func (c Chinese) sayEating() {
	fmt.Println("吃了吗")
}
func (a American) disco() {
	fmt.Println("disco")
}
func greet(s Say) {
	s.saySomething()
	//断言
	ch, flag := s.(Chinese)

	if flag {
		ch.sayEating()
	} else {
		fmt.Println("美国朋友")
	}

	/*简洁写法
	  if ch, flag := s.(Chinese); flag {
	  		ch.sayEating()
	  	} else {
	  		fmt.Println("美国朋友")
	  	}
	*/
	//type属于go中的一个关键字，固定写法
	switch s.(type) {
	case Chinese:
		ch := s.(Chinese)
		ch.sayEating()
	case American:
		am := s.(American)
		am.disco()
	}

}
func main() {
	c := Chinese{"小明"}
	a := American{"撒旦撒旦"}

	greet(c)
	greet(a)
}
