package main

import "learngo/test/interface/dog"

type Dog interface {
	GetName() string
	GetAge() float32
	GetDad() string
	eat()
	shout()
}

func main() {
	r := dog.Dog{"麦兜",1.5,"ydf"}
	r.GetName()
	r.GetAge()
	r.GetDad()
	r.Eat()
	r.Shout()
}
