package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const arabic = "Arabic"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const arabicHelloPrefix = "مرحباً, "

// Hello function to return a string to say hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language)
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case arabic:
		prefix = arabicHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", "spanish"))
}
