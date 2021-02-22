package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}
func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	//fmt.Println(resp)
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))

	//fmt.Println("DOING THIS A DIFF WAY USING IO.COPY ====================")
	////a you can also do this to print out html to screen. copy implements the writer interface which can do the above code.
	//io.Copy(os.Stdout, resp.Body)

	fmt.Println("WRITING NOW USING custom writer!!!")

	io.Copy(logWriter{}, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}