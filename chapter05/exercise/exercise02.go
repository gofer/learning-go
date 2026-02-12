package main

import (
	"fmt"
	"io"
	"os"
)

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	totalLen := 0
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
		totalLen += count
	}
	return totalLen, nil
}

func exercise02() {
	size, err := fileLen("exercise02.go")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("size:", size)
}
