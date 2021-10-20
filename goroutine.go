package main

import (
	"fmt"
	"sync"
	"time"
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

	s1 := []int{1, 2, 3, 4, 5}
	c1 := make(chan int, len(s1))
	go goroutine3(s1, c1)
	for i := range c1 {
		fmt.Println(i)
	}

	var wg1 sync.WaitGroup
	ch1 := make(chan int)

	// Producer
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go producer(ch1, i)
	}

	go consumer(ch1, &wg1)
	wg1.Wait()
	close(ch1)

	chan1 := make(chan string)
	chan2 := make(chan string)
	go goroutine4(chan1)
	go goroutine5(chan2)

	// for {
	// 	select {
	// 	case msg1 := <-chan1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-chan2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OuterLoop:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println("######################")

	cw := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			cw.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			cw.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(cw, cw.Value("Key"))

	words := []string{"test1!", "test2!", "test3!", "test4!"}
	c := make(chan string)
	go goroutine6(words, c)
	for w := range c {
		fmt.Println(w)
	}

}

func goroutine6(s []string, c chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
}

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}
func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func goroutine4(ch chan string) {
	for {
		ch <- "packet from 4"
		time.Sleep(1 * time.Second)
	}
}
func goroutine5(ch chan string) {
	for {
		ch <- "packet from 5"
		time.Sleep(1 * time.Second)
	}
}

func producer(ch chan int, i int) {
	ch <- i * 2
}
func consumer(ch chan int, wg1 *sync.WaitGroup) {
	for i := range ch {
		func() {
			defer wg1.Done()
			fmt.Println("process", i*1000)
		}()
	}
}

func goroutine3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
		c <- sum
	}
	close(c)
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
