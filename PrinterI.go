package main

import "fmt"

type PrinterI interface {
	Print()
}

type Document struct {
	content string
}

func (d Document) Print() {
	fmt.Println(d.content)
}

func printDocument(printer PrinterI) {
	printer.Print()
}

func main() {
	doc := Document{content: "THIS SHOULD BE YOUR TEXT"}
	printDocument(doc)
}
