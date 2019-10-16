package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	<-done
	// 如果你把 <- done 这行代码从程序中移除，程序甚至会在 worker 还没开始运行时就结束了。
}
