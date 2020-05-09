package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet just greets
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler Handle greeting
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
func main() {
	// Greet(os.Stdout, "Ismail")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
