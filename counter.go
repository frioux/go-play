package main

import(
	"bufio"
	"fmt"
	"os"
)

func main () {
	acc := make(map[string]int)

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		acc[s.Text()]++
	}

	for k, v := range acc {
		fmt.Println(k, "count is", v)
	}
}
