package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			address := "127.0.0.1:8080"
			udpAddr, err := net.ResolveUDPAddr("udp", address)
			if err != nil {
				fmt.Println("Error resolving address:", err)
				return
			}

			conn, err := net.DialUDP("udp", nil, udpAddr)
			if err != nil {
				fmt.Println("Error connecting:", err)
				return
			}
			defer conn.Close()

			message := "Hello, server!"
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Println("Error sending:", err)
				return
			}
			//fmt.Println("Sent:", message)
		}()

	}
	wg.Wait()
}
