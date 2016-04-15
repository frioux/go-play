package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func fetch(url string, c chan int) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching", url, ":", err)
		c <- 509
	}
	c <- res.StatusCode
}

func main() {
	g := make(chan int, 1)
	b := make(chan int, 1)
	y := make(chan int, 1)

	go fetch("http://google.com", g)
	go fetch("http://bing.com", b)
	go fetch("http://yahoo.com", y)
	to := time.After(15 * time.Second)

	select {
	// case i := <-b:
	// 	fmt.Println("Bing's status was:", i)
	// case i := <-g:
	// 	fmt.Println("Google's status was:", i)
	case i := <-y:
		fmt.Println("Yahoo's status was:", i)
	case <-to:
		fmt.Println("TOO SLOW")
		os.Exit(1)
	}
}
