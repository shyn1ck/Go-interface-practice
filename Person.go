package main

import "fmt"

type PersonI interface {
	GetName() string
	GetAge() int
}

type Student struct {
	name string
	age  int
}

func (s *Student) GetName() string {
	return s.name
}

func (s *Student) GetAge() int {
	return s.age
}

func printPersonInfo(person PersonI) {
	fmt.Printf("Имя: %s, Возраст: %d\n", person.GetName(), person.GetAge())
}

func main() {
	student := Student{name: "Vasya", age: 20}
	printPersonInfo(&student)
}
