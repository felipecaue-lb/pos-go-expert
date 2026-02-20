package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, error := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
	if error != nil {
		panic(error)
	}

	res, error := http.DefaultClient.Do(req)
	if error != nil {
		panic(error)
	}
	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
