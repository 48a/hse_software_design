package main

import "fmt"

type Engine struct {
	Power int
}

type Radio struct {
	Power int
}

type Car struct {
	Engine
	Radio
}

func main() {
	c := Car{
		Engine: Engine{},
		Radio:  Radio{},
	}
	fmt.Println(c.Engine.Power, c.Radio.Power)
}
