package main

import "fmt"

// Embedded
type Vertex1 struct {
	x, y int
}

// func Area(v Vertex1) int {
// 	return v.x * v.y
// }

// 値レシーバー（値渡し：値は書き変わらない）
func (v Vertex1) Area() int {
	return v.x * v.y
}

// ポインタレシーバー（参照渡し：値が書き変わる）
func (v *Vertex1) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

type Vertex2 struct {
	Vertex1
	z int
}

// func Area(v Vertex1) int {
// 	return v.x * v.y
// }

// 値レシーバー（値渡し：値は書き変わらない）
func (v Vertex2) Area2() int {
	return v.x * v.y * v.z
}

// ポインタレシーバー（参照渡し：値が書き変わる）
func (v *Vertex2) Scale2(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func New(x, y, z int) *Vertex2 {
	return &Vertex2{Vertex1{x, y}, z}
}

// non-struct
type MyInt int

// non-structメソッド
func (i MyInt) Double() int {
	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", 1, 1)
	return int(i * 2)
}

func (i MyInt) Double2() MyInt {
	return i * 2
}

// インターフェース
type Human interface {
	Say()
}

type Person struct {
	Name string
}

func (p Person) Say() {
	fmt.Println(p.Name)
}

func learning() {
	defer fmt.Println("-----------------")
	fmt.Println("-----------------")

	v := Vertex1{3, 4}
	// fmt.Println(Area(v))
	fmt.Println(v.Area())
	fmt.Println(v)

	v.Scale(10)
	fmt.Println(v.Area())
	fmt.Println(v)

	v1 := New(3, 4, 5)
	v1.Scale(10)
	v1.Scale2(10)
	fmt.Println(v1.Area())
	fmt.Println(v1.Area2())

	myInt := MyInt(10)
	fmt.Println(myInt.Double())
	fmt.Println(myInt.Double2())

	var mike Human = Person{"Mike"}
	mike.Say()

	do(10)
	do("test")
	do(true)

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

}

type UserNotFound struct {
	Username string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User not found: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{Username: "mike"}
}

func do(i interface{}) {
	// タイプアサーション
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	// ss := i.(string)
	// fmt.Println(ss + "!!")

	// switch type文
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!!")
	default:
		fmt.Printf("I don't know %T\n", v)
	}

	// int をfloat64に変換することをコンバージョンという
	// 上記はタイプアサーション
}
