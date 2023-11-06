package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	ch := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second)
		<-ch
	}()
	ch <- true
	fmt.Printf("cost %.f s\n", time.Now().Sub(st).Seconds())
	ch <- true
	fmt.Printf("cost %.f s\n", time.Now().Sub(st).Seconds())
	ch <- true
	time.Sleep(time.Second * 3)
}
