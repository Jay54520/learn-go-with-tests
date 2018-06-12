package main

import "fmt"

const helloPrefix  = "Hello, "
const world  = "World"
const spanish  = "Spanish"
const spanishHelloPrefix  = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = world
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	return helloPrefix + name
}

func main()  {
	fmt.Println(Hello("world", ""))
}
