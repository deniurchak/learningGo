package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {

}
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Printf("Just wrote %d bytes", len(bs))
	return len(bs), nil
}

func main() {
	url := "http://google.com"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, res.Body)
}