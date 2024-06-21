package main

import "fmt"

type AnimalI interface {
	MakeSound()
	GetName() string
}

type Dog struct {
	name string
}

type Cat struct {
	name string
}

func (d Dog) MakeSound() {
	fmt.Println("Dog make sound: Gav")
}

func (d Dog) GetName() string {
	return d.name // Возвращаем имя собаки
}

func (c Cat) MakeSound() {
	fmt.Println("Cat make sound: Meow")
}

func (c Cat) GetName() string {
	return c.name
}

func AnimalAction(animal AnimalI) {
	fmt.Printf("Имя животного: %s\n", animal.GetName())
	animal.MakeSound()
}

func main() {
	dog := Dog{name: "Мухтар"}
	cat := Cat{name: "Мошон"}
	AnimalAction(dog)
	AnimalAction(cat)
}
