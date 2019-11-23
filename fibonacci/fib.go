package main

import (
	"fmt"
	"time"
)

var x = map[int]int{}

func timed(f func(n int) int, n int) int {
	t := time.Now()
	q := f(n)
	since := time.Since(t)
	fmt.Println(since)
	return q
}

// fib is having time complexity of O(2^n)
// Since many calculations are repeated
// We can optimize this using dynamic programming approach
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// for smaller number difference will not be huge
// but as you proceed with large numbers
// differences can be observed easily.
// It is obvious that space complexity has been increased
func fibDP(n int) int {
	_, ok := x[n]
	if !ok {
		if n <= 1 {
			return n
		}
		x[n] = fibDP(n-1) + fibDP(n-2)
		return x[n]
	}
	return x[n]
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println("Dynamic Approach: ", timed(fibDP, n))
	fmt.Println("Naive Approach: ", timed(fib, n))
}
