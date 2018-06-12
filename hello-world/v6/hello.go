package main

import "fmt"

const helloPrefix  = "Hello, "
const world  = "World"
const spanish  = "Spanish"
const spanishHelloPrefix  = "Hola, "
const french  = "French"
const frenchHelloPrefix  = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = world
	}

	prefix := helloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main()  {
	fmt.Println(Hello("world", ""))
}
