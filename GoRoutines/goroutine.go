//Go Routines
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func f1(c chan string) {
// 	c <- "Hello World! Inside f1 function"
// }

// func f2(c chan string) {
// 	msg := <-c
// 	fmt.Println("Inside f2 function", msg)
// }

// func f3(c chan string) {
// 	msg := <-c
// 	fmt.Println("Inside f3 function", msg)
// }
// func main() {
// 	var c chan string = make(chan string)
// 	go f2(c)
// 	go f1(c)
// 	go f3(c)
// 	time.Sleep(100 * time.Millisecond)
// }

//Un Buffered Channel
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(done chan bool) {
// 	fmt.Println("working start")
// 	time.Sleep(time.Second)
// 	fmt.Println("working end")

// 	done <- true
// }

// func main() {

// 	done := make(chan bool, 1)
// 	go worker(done)

// 	<-done

// 	fmt.Println("Main func")
// }

// Buffered Channel
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 		time.Sleep(time.Second)
// 	}
// 	close(ch)
// }
// func main() {

// 	// creates capacity of 2
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 	}
// }

// Sum of Squares
// package main

// import (
// 	"fmt"
// )

// func square(nums []int, ch chan int) {
// 	for _, num := range nums {
// 		ch <- num * num
// 	}
// }

// func main() {
// 	nums := []int{1, 3, 5}
// 	ch := make(chan int)
// 	go square(nums, ch)
// 	sum := 0
// 	v := 0
// 	for v < len(nums) {
// 		sum += <-ch
// 		v++
// 	}
// 	fmt.Println(sum)
// }

// Find occurence of all characters in a string
package main

import "fmt"

func main() {
	str := "Hello World"
	ch := make(chan rune)
	go func() {
		for _, r := range str {
			ch <- r
		}
		close(ch)
	}()
	m := make(map[string]int)
	for r := range ch {
		m[string(r)]++
	}
	fmt.Println(m)
}
