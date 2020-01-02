package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
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

/**
* return byte length and character length
 */
func Rune(s string) (int, int) {
	return len(s), utf8.RuneCountInString(s)
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

	type Rope string

	var ss Rope = "sbc"
	fmt.Printf("the string is: %s\n", ss)

	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3)
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3)
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3)
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3)

	str := "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Printf("The index is: %d\n", strings.IndexRune(str, rune('o')))

	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for i, val := range sl {
		if i < len(sl)-1 {
			fmt.Printf("%s - ", val)
		} else {
			fmt.Printf("%s", val)
		}
	}
	fmt.Println()

	orig := "666"
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	ns := strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", ns)

	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	week := 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	s1 := t.Format("20160102")
	fmt.Println(t, "=>", s1)

	s2 := "good byte"
	var p *string = &s2
	*p = "hello"
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s2: %s\n", s2)

	sea := 9
	switch sea {
	case 1:
		fmt.Println("It's mon.")
	case 2:
		fmt.Println("It's sec.")
	case 3:
		fmt.Println("It's The.")
	case 9:
		fmt.Println("It's sep.")
	default:
		fmt.Println("Out of season.")
	}
}
