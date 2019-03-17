package main

import "fmt"

func main() {
	last := 0
	first, second := 1, 1
	for i := 0; last < 10; i++ {
		if first + second == 2 {
			fmt.Println(first)
			fmt.Println(second)
		}
		last = first + second
		fmt.Println(last)
		first, second = second, last
	}
}