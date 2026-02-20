package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Second}
	resp, error := c.Get("https://google.com")
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		panic(error)
	}
	println(string(body))
}
