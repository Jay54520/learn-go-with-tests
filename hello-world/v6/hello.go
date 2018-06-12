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

	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	}

	return helloPrefix + name
}

func main()  {
	fmt.Println(Hello("world", ""))
}
