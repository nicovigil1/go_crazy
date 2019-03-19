package main

import "testing"

func TestFibonacci(t *testing.T) {
	fib := fibonacci()
	results := []int{1, 1, 2, 3, 5}
	for _, i := range results {
		if fib() != i {
			t.Errorf("expected the returned value %d, to equal %d", fib()-1, i)
		}
	}
}