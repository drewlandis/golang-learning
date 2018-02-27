package main

import (
	"fmt"
	"strconv"
)

type book struct {
	title  string
	author string
}

func (b book) printInfo() {
	fmt.Println(b.getInfo())
}

func (b book) getInfo() string {
	return b.title + " written by " + b.author
}

type shelf struct {
	length int
	slot   int
	book   book
}

func (s *shelf) create(length int) {
	s.length = length
}

func (s *shelf) shelveBook(b book, slot int) {
	s.slot = slot
	s.book = b
}

func (s *shelf) getInfo() string {
	return "This shelf has " + s.book.getInfo() + " in slot " + strconv.Itoa(s.slot)
}

func (s *shelf) printInfo() {
	fmt.Println(s.getInfo())
}

type library struct {
	height int
	row    int
	shelf  shelf
}

func (l *library) create(height int) {
	l.height = height
}

func main() {
	bk1 := book{title: "Don Quixote", author: "Miguel de Cervantes"}
	bk2 := book{title: "Three Musketeers", author: "Alexandre Dumas"}
	bk3 := book{title: "Atlas Shrugged", author: "Ayn Rand"}
	bk4 := book{title: "Grapes of Wrath", author: "John Steinbeck"}
	bk5 := book{title: "Of Mice and Men", author: "John Steinbeck"}

	sh1 := &shelf{}
	sh1.create(10)
	sh1.shelveBook(bk1, 4)
	sh1.printInfo()

	bk1.printInfo()
	bk2.printInfo()
	bk3.printInfo()
	bk4.printInfo()
	bk5.printInfo()
}
