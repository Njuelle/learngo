package main

import "fmt"

const defaultName = "World"
const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello "
const spanishPrefix = "Hola "
const frenchPrefix = "Bonjour "

func main() {
	fmt.Println(Hello("world", "English"))
}

// Hello someone in many languages
func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}
