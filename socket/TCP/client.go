package main

import (
	"log"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			conn, err := net.Dial("tcp", "127.0.0.1:23345")
			if err != nil {
				log.Printf("Error connecting,err:%v", err)
				return
			}
			defer conn.Close()

			msg := "hello ,world!"
			_, err = conn.Write([]byte(msg))
			if err != nil {
				log.Printf("conn.Write fialed,err:%v", err)
				return
			}
			log.Println("sent:", msg)
		}()

	}
	wg.Wait()
}
