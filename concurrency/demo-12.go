package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(os.Stdout, res.Body)
}
