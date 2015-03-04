package main

import(
	"net/http"
	"fmt"
	"os"
	"time"
)

func fetchUrl(url string, c chan int) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("could not fetch", url, ":", err)
		c <- 509
	}
	c <- res.StatusCode
}

func main() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c3 := make(chan int, 1)

	acc := make(map[string]int)

	go fetchUrl("https://blog.afoolishmanifesto.com", c1)
	go fetchUrl("https://blog.afoolishmanifesto.com", c2)
	go fetchUrl("https://blog.afoolishmanifesto.com", c3)

	to := make(chan bool, 1)
	go func() {
		 time.Sleep(15 * time.Second)
		 to <- true
	}()

	i := 0
	Loop:
	for {
		select {
		case a := <- c1:
			acc["a"] = a
			fmt.Println("got a", i)
			i++
			if i == 3 {
				break Loop
			}
		case b := <- c2:
			acc["b"] = b
			fmt.Println("got b", i)
			i++
			if i == 3 {
				break Loop
			}
		case c := <- c3:
			acc["c"] = c
			fmt.Println("got c", i)
			i++
			if i == 3 {
				break Loop
			}
		case <- to:
			fmt.Println("took too long!")
			os.Exit(1)
		}
	}

	for k, v := range acc {
		fmt.Println(k, ":", v)
	}
}
