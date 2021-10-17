package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "string"
	t, f bool    = true, false
)

const pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)

// 最初に呼ばれる init
func init() {
	fmt.Println("init!")
}

// 実行される関数 main
func main() {
	fmt.Printf("Hello world\n")
	// buzz()
	variable()
	constant()
	str()
	type_conversion()
	array()
	slices()
	map_func()
	r, h := add(5, 8)
	fmt.Println(r, h)
	t := cal(10, 30)
	fmt.Println(t)

	// 右辺が無名関数
	f := func() {
		fmt.Println("inner func")
	}
	f()

	// 無名関数（即実行）
	func(x int) {
		fmt.Println("inner func", x)
	}(8)

	counter := incremantGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := circleArea(3.14)
	c2 := circleArea(3)

	fmt.Println(c1(2))
	fmt.Println(c2(2))

	foo()
	foo(10, 20)
	foo(10, 20, 30)

	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)

	f1 := 1.11
	fmt.Println(int(f1))

	s1 := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s1[2:4])

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v\n", m, m)

	num := 4
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}
	if result := 2; result%2 == 0 {
		fmt.Println("ok")
	}

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)

	l := []string{"python", "go", "java"}
	for i, v := range l {
		fmt.Println(i, v)
	}
	for _, v := range l {
		fmt.Println(v)
	}

	m2 := map[string]int{"apple": 100, "banana": 500}
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	for k := range m2 {
		fmt.Println(k)
	}
	for _, v := range m2 {
		fmt.Println(v)
	}

	os := "mac"
	switch os {
	case "mac":
		fmt.Println("mac!")
	case "windows":
		fmt.Println("windows!")
	default:
		fmt.Println("default!")
	}

	switch os1 := "windows"; os1 {
	case "mac":
		fmt.Println("mac!")
	case "windows":
		fmt.Println("windows!")
	default:
		fmt.Println("default!")
	}

	t1 := time.Now()
	fmt.Println(t1.Hour())

	foo2()
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	l1 := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	x1 := l1[0]

	for i = 0; i < len(l1); i++ {
		if x1 > l1[i] {
			x1 = l1[i]
		}
	}
	fmt.Printf("最小値は%vです\n", x1)

	m1 := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	sum1 := 0
	for _, v := range m1 {
		sum1 += v
	}
	fmt.Printf("合計は%vです\n", sum1)

	n := 100
	fmt.Println(n)

	fmt.Printf("%T %v\n", &n, &n)
	fmt.Printf("%T %v\n", n, n)
	// fmt.Printf("%T %v\n", *&n, *&n)

	// var p *int = &n
	// fmt.Println(p)
	// fmt.Println(*p)

	// one(&n)
	// fmt.Println(n)

	var p1 *int = new(int)
	fmt.Println(p1)
	// => メモリを確保しているのでポインタが出力される

	var p2 *int
	fmt.Println(p2)
	// => メモリを確保していないのでnil

}

// func one(x *int) {
// 	*x = 1
// }

func foo2() {
	defer fmt.Println("world foo")
	fmt.Println("hello foo")
}

// func buzz() {
// 	fmt.Println("hello buzz!!", time.Now())
// 	fmt.Println(user.Current())
// }

func variable() {
	// var i int = 1
	// var f64 float64 = 1.2
	// var str string = "string"
	// var t, f bool = true, false

	// var (
	// 	i    int
	// 	f64  float64
	// 	str  string
	// 	t, f bool
	// )

	fmt.Println(i, f64, s, t, f)

	// var宣言は関数外に定義できるが、「short declaration」の場合はできない（関数ないのみ）
	xi := 1 + 1
	xf64 := 1.2
	xstr := "short"
	xt, xf := true, false
	fmt.Println(xi, xf64, xstr, xt, xf)
	fmt.Printf("%T\n", xf64)
	fmt.Printf("%T\n", xi)
}

func constant() {
	fmt.Println(pi, Username, Password)
	fmt.Printf("%T %v", pi, pi)
}

func str() {
	s := "Hello World"
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Println(string(s[0]))

	fmt.Println(strings.Replace(s, "H", "X", 1))
	fmt.Println(s)

	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	fmt.Println(strings.Contains(s, "World"))
	fmt.Println(`Test
	Test`)
}

func type_conversion() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	// i := int(s) => 変換できない

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Printf("%T %v %d\n", i, i, i)

	// Atoiメソッドはintとerr（エラー）を返す。そのため、代入するときは「i, err」のようにする
	// ただし、errが返ってきても処理をしない場合は「i, _」とするとエラーにならない

	// 下記の場合、「err」を処理していないのでエラーになる
	// i, err := strconv.Atoi(s)
	// fmt.Printf("%T %v %d\n", i, i, i)

	// 下記のようにすると、エラーにならない
	xi, _ := strconv.Atoi(s)
	fmt.Printf("%T %v %d\n", xi, xi, xi)

}

func array() {
	var a [2]int
	a[0] = 1
	a[1] = 2
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)
	// b = append(b, 300) => 配列はappendできない
	fmt.Printf("%T %v\n", b, b)
}

func slices() {
	n := []int{1, 2, 3, 4, 5, 100}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:6])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}
	fmt.Println(board)

	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)

	k := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(k), cap(k), k)
	k = append(k, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(k), cap(k), k)
	k = append(k, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(k), cap(k), k)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}

func map_func() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	fmt.Println(m["nothing"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)
}

func add(x, y int) (int, int) {
	fmt.Println("add function")
	fmt.Println(x + y)
	return x + y, x - y
}

func cal(price, item int) (result int) {
	// result := price * item => エラー
	result = price * item
	return //return resultじゃなくてもOK
}

func incremantGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

// 可変調引数
func foo(params ...int) {
	fmt.Println(len(params), params)
}
