package main

import (
	"fmt"
)

func main() {
	var people []Person
	var choice string

	fmt.Println("Person registry")

	for {
		choice = input("1 - create\n2 - read all\n3 - read one\n4 - edit\n5 - delete\n6 - exit")
		switch choice {
		case "1":
			createPerson(&people)
		case "2":
			viewAllPerson(people)
		case "3":
			viewPerson(people, input("Cpf of the person"))
		case "4":
			personCpf := input("Cpf of the person")
			editPerson(&people, personCpf)
		case "5":
			personCpf := input("Cpf of the person")
			people = deletePerson(people, personCpf)
		case "6":
			return
		}
	}
}
