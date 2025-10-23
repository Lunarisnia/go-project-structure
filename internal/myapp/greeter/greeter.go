package greeter

import "fmt"

type Greeter interface {
	SayHello()
}

type greeterImpl struct {
}

func CreateNewGreeter() Greeter {
	return &greeterImpl{}
}

func (greeterImpl *greeterImpl) SayHello() {
	fmt.Println("Hello, World")
}
