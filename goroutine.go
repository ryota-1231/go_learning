package main

import (
	"fmt"
)

func goroutine() {
	defer fmt.Println("-------go----------")
	fmt.Println("--------go---------")

	// s := []int{1, 2, 3, 4, 5}
	// c := make(chan int)

	// go goroutine1(s, c)
	// go goroutine2(s, c)
	// x := <-c
	// y := <-c
	// fmt.Println(x)
	// fmt.Println(y)

	// Goroutine:並列処理
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go goroutineFunc("world", &wg)
	// normal("hello")
	// wg.Wait()

	ch := make(chan int, 2)
	ch <- 100
	fmt.Println(len(ch))
	ch <- 200
	fmt.Println(len(ch))

	x1 := <-ch
	x2 := <-ch
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(len(ch))

	ch <- 100
	ch <- 200
	close(ch)

	for c := range ch {
		fmt.Println(c)
	}

}
func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// func normal(s string) {
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(1000 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
// func goroutineFunc(s string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 5; i++ {
// 		// time.Sleep(1000 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
