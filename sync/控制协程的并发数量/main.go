package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	c := make(chan struct{}, 3)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- struct{}{}
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			log.Println(x)
			time.Sleep(time.Second)
			<-c
		}(i)
	}
	wg.Wait()
}
