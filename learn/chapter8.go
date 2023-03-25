package main

import "fmt"

type Person struct {
	name    string
	age     int16
	address string
}

var person Person = Person{"ZhangSan", 21, "上海"}

func main() {
	//person := Person{"ZhangSan", 21, "上海"}
	fmt.Println(person.name)
	renamePerson(person, "King")
	fmt.Println(person.name)
	renamePerson2(&person, "King")
	fmt.Println(person.name)
}

func renamePerson(person Person, s string) {
	person.name = s
}
func renamePerson2(person *Person, s string) {
	person.name = s
}
