package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	const goRuntimes = 7
	num := 0
	var waitGroup sync.WaitGroup
	waitGroup.Add(goRuntimes)
	mutex.Lock()
	for i := 0; i < goRuntimes; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			num++
			fmt.Println("insied func 1", runtime.NumGoroutine())
			waitGroup.Done()
		}()
	}
	mutex.Unlock()

	go func() {
		time.Sleep(time.Second / 10)
		mutex.Lock()
		num--
		mutex.Unlock()
		fmt.Println("insied func 2", runtime.NumGoroutine())
		waitGroup.Done()
	}()

	waitGroup.Wait()
	fmt.Println("num", num)
}
