package main 

import "fmt"

func main() {
	fib := fibonacci()
	for i:=0; i<10; i++ {
		fmt.Println(fib())
	}
}

func fibonacci() func() int {
	int1 := 0
	int2 := 1
	return func() int {
		next := int1 + int2
		int1, int2 = int2, next
		return int1  
	}
}