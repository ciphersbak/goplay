package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// https://medium.com/technofunnel/understanding-golang-and-goroutines-72ac3c9a014d
	// runtime.GOMAXPROCS(4)
	fmt.Println("Starting concurrent calls...")
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	start := time.Now()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
		waitGroup.Done()
	}()
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
		waitGroup.Done()
	}()
	waitGroup.Wait()
	elapsedTime := time.Since(start)
	fmt.Println("Total Time for Execution: " + elapsedTime.String())
	time.Sleep(time.Second)
}
