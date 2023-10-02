package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("mock.jpg")
	if err != nil {
		log.Fatal(err)
	}
	data, size, err := readImg(fd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data: ")
	fmt.Println(data)
	fmt.Println("size: ")
	fmt.Println(size)

	fmt.Println("writing file ...")
	err = os.WriteFile("result.jpg", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("process finish ...")

}

func readImg(fd *os.File) ([]byte, int, error) {

	buf := make([]byte, 2048)
	bufSize := 0
	for {
		n, err := fd.Read(buf)
		bufSize += n
		if err != nil {
			if err == io.EOF {
				fmt.Println("檔案結束")
				return buf[:bufSize], bufSize, nil
			}
			return nil, 0, err
		}
	}
}
