package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

type traceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := traceCode("Trace_Code")
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
	log.Printf("worker func trace %s ...\n", traceCode)

LOOP:
	for {
		fmt.Printf("woker,trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:

		}
	}
	wg.Done()
	fmt.Println("done")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	key := traceCode("Trace_Code")
	value := "123456"
	ctx = context.WithValue(ctx, key, value)
	wg.Add(1)
	log.Printf("main func traceCode:%s", ctx.Value(key))
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}

/*

 */
