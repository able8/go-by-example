package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
