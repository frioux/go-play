package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 0

	for {
		n := <-c
		fmt.Println(n)
		c <- n + <-c
		c <- n
	}
}
