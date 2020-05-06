package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello function to return a string to say hello
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
