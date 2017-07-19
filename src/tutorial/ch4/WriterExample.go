package main

import (
	"io"
	"fmt"
	"time"
	"os"
	"net/http"
)

func handle(w io.Writer, msg string) {
	fmt.Fprintln(w, msg)
}

func ioWriter() {
	msg := []string{"This", "is", "an", "example", "of", "io.Writer"}

	for _, s := range msg {
		time.Sleep(100 * time.Millisecond)
		handle(os.Stdout, s)
	}
}

func httpResponseWriter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		handle(w, r.URL.Path[1:])
	})

	fmt.Println("start listening on port 4000")
	http.ListenAndServe(":4000", nil)
}

func main() {
	//ioWriter()
	httpResponseWriter()
}