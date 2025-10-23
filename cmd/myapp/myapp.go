package main

import "github.com/lunarisnia/go-project-structure/internal/myapp/greeter"

func main() {
	greeter := greeter.CreateNewGreeter()
	greeter.SayHello()
}
