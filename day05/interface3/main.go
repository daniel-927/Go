package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步....")
}
func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...!🐟\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动!")
}

func (c chicken) eat(kfc string) {
	fmt.Printf("吃鸡%s!🐓\n", kfc)
}

func main() {

	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	a1.eat("小黄鱼")
	fmt.Printf("%T\n", a1)
	fmt.Println(a1)

	kfc := chicken{
		feet: 2,
	}
	a1 = kfc
	a1.eat("KFC")
	fmt.Printf("%T\n", a1)
}
