package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const port = ":8787"

func main() {
	fmt.Println("listening ", port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		err = handleUpload(conn)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func handleUpload(conn net.Conn) error {

	data := make([]byte, 2048)
	dataSize := 0

	for {

		n, err := conn.Read(data)
		dataSize += n
		if err != nil {
			if err == io.EOF {
				fmt.Println("資料讀取結束...")
				break
			}
			return err
		}

	}

	err := os.WriteFile("result.jpg", data[:dataSize], 0644)
	if err != nil {
		return err
	}
	return nil
}
