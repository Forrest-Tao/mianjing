package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:23345")

	if err != nil {
		log.Printf("net.Listen failed,err:%v", err)
		return
	}

	defer listener.Close()

	log.Println("server is listening on 127.0.0.1")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("err:%v", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()

			buffer := make([]byte, 2<<10)
			_, err := conn.Read(buffer)
			if err != nil {
				log.Printf("conn.Read failed,err:%v", err)
				return
			}
			log.Printf("Received :%s\n", string(buffer))
		}(conn)
	}

}
