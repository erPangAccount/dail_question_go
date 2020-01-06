package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	//test1()
	test2()
}

func test1() {
	true := false
	log.Println(true)
}

func test2() {
	runtime.GOMAXPROCS(1)
	ch := make(chan int)
	go func() {
		fmt.Println("start")
		ch <- 0
		fmt.Println("end")
	}()
	fmt.Println("wait")
	<-ch
	fmt.Println("done")
}
