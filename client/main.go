package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const url = "localhost:8787"

func main() {

	data, err := readFile("mock.jpg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	if err := sendFile(data, url); err != nil {
		log.Fatal(err)
	}

}

func sendFile(data []byte, url string) error {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("write done")
	return nil
}

func readFile(path string) ([]byte, error) {

	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	data := make([]byte, 2048)

	n, err := fd.Read(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data[:n], nil
}
