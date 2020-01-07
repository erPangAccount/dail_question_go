package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	//test1()

	test2()
}

func test1() {
	var ch chan int // nil channel  如果用make 则返回初始化(分配内存)好的channel

	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch)
	}()
	<-ch
	log.Println(count)
}

func test2() {
	var ch chan int
	go func() {
		ch = make(chan int)
		ch <- 1
	}()

	go func(ch chan int) {
		time.Sleep(time.Second)
		<-ch
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Println(runtime.NumGoroutine())
	}
}
