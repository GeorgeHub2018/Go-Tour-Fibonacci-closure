package main

import "fmt"

func fib(n int) int64 {
	if n < 2 {
		return int64(n)
	}
	return fib(n-1) + fib(n-2)
}

func fibonacci() func() int64 {
	i := -1
	return func() int64 {
		i++
		return fib(i)
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
