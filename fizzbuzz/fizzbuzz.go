package main

import ("os"
				"strconv"
        "fmt")
// test := []int

func main() {
	for _, j := range buildArray() { 
		switch {
		case j % 15 == 0:
			fmt.Println("fizz buzz")
		case j % 3 == 0:
			fmt.Println("fizz")
		case j % 5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(j) 
		}
	}
}

func buildArray() []int {
	length, _ := strconv.Atoi(os.Args[1])
	test_slice := make([]int, length)
	for i := 0; i <= length-1; i++ {
		test_slice[i] += i + 1
	}
	return test_slice
}