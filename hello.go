package main

import (
	"fmt"
)

type People struct {
	Name string
	Age  int
}

func (self *People) sayHello() {
	fmt.Printf("I'm %s, %d years old, i'm glad to meet you!\n", self.Name, self.Age)
}

func muliReturn(a, b string) (x, y string) {
	x = a
	y = b
	return x, y
}

func main() {
	xiaoming := People{
		Name: "xiaoming",
	}

	var array [2]string
	array[0] = "pitone"
	array[1] = "allday"

	fmt.Println(array)

	i := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(i)

	s := i[1:5]

	fmt.Println(s)

	fmt.Printf("s, leng: %d, cap: %d\n", len(s), cap(s))

	st := []int{6, 7, 8, 9, 10, 11}

	fmt.Println(st)
	fmt.Printf("st, length: %d, cap: %d\n", len(st), cap(st))

	xiaoming.Age = 10
	xiaoming.sayHello()
	fmt.Println(muliReturn("hello", "world"))
}
