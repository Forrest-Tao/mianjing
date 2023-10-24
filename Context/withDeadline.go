package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer func() {
		cancel()
		fmt.Println("main cancel ...")
	}()
	//cancel()//context canceled

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overtime")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) //context deadline exceeded
	}

}
