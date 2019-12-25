package main

import (
	"fmt"
	"math"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
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

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
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

	var n int16 = 34
	var m int32
	m = int32(n)

	fmt.Printf("32 bit int is: %d\n", m)
	fmt.Printf("16 bit int is: %d\n", n)

	x, error := Uint8FromInt(100000)
	fmt.Printf("the result is: %d, info: %s\n", x, error)
	flag := 1 | 2
	fmt.Printf("the flag is: %d\n", flag)
}
