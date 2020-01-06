package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("第一题：")
	one()

	fmt.Println("第二题：")
	two()
}

func one() {
	fmt.Println(one_doubleScore(0))
	fmt.Println(one_doubleScore(20))
	fmt.Println(one_doubleScore(50))
}

func one_doubleScore(source float32) (score float32) {
	defer func() {
		if source < 1 || score > 100 {
			score = source
		}
	}()
	return source * 2
}

var mu sync.RWMutex
var count int

func two() {
	go tow_A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func tow_A() {
	mu.RLock()
	defer mu.RUnlock()
	tow_B()
}

func tow_B() {
	time.Sleep(5 * time.Second)
	tow_C()
}

func tow_C() {
	mu.RLock()
	defer mu.RUnlock()
}
