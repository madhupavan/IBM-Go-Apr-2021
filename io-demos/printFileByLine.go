package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	inputHandle, inputError := os.Open("sample.txt")
	defer inputHandle.Close()
	if inputError != nil {
		log.Fatalf("Error opening the file: %v", inputError)
	}
	inputReader := bufio.NewReader(inputHandle)
	for {
		line, err := inputReader.ReadString('\n')
		fmt.Printf("Line : %s\n", line)
		if err == io.EOF {
			break
		}

	}
}
