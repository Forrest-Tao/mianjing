package main

import (
	"fmt"
	"time"
)

func main() {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		ch <- true
	}()
	<-ch
	fmt.Printf("cost %.f s\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 3)
}
