package main

import (
	"context"
	"fmt"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("goroutine out")
				return
			case dst <- n:
				n++
			}
		}
		//fmt.Println("goroutine out")
	}()
	fmt.Println("gen out  ...", n)
	return dst
}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	defer func() {
		ctx.Done()
		fmt.Println("main cancel ...")
	}()

	for n := range gen(ctx) {
		fmt.Println("n", n)
		if n == 5 {
			break
		}
	}
}
