package main

import "fmt"

const englishHelloPrefix = "Hello, "
const world  = "World"
const spanish  = "Spanish"
const spanishHelloPrefix  = "Hola, "
const french  = "French"
const frenchHelloPrefix  = "Bonjour, "

func greetingPrefix(language string) (prefix string)  {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = world
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func main()  {
	fmt.Println(Hello("world", ""))
}
