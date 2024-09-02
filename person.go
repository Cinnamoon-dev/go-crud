package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name  string
	Email string
	Cpf   string
	Age   uint8
}

func getPerson(people []Person) Person {
	var newPerson Person

	newPerson.Name = input("Name:")
	newPerson.Email = input("Email:")
	newPerson.Cpf = inputCpf(people)
	newPerson.Age = uint8(inputInt("Age:"))

	return newPerson
}

func createPerson(people *[]Person) {
	newPerson := getPerson(*people)
	*people = append(*people, newPerson)
}

func viewAllPerson(people []Person) {
	for i := 0; i < len(people); i++ {
		fmt.Println(people[i])
	}
}

func viewPerson(people []Person, cpf string) {
	for _, v := range people {
		if v.Cpf == cpf {
			fmt.Println(v)
			return
		}
	}

	fmt.Println("Person not found")
}

func findPerson(people []Person, cpf string) (int, error) {
	for i := 0; i < len(people); i++ {
		if (people)[i].Cpf == cpf {
			return i, nil
		}
	}

	return -1, errors.New("Person not found")
}

func editPerson(people *[]Person, cpf string) {
	personIndex, err := findPerson(*people, cpf)
	if err != nil {
		fmt.Println(err)
		return
	}

	choice := input("What field do you want do change?\n1 - Name\n2 - Email\n3 - Age\n4 - Cpf")
	switch choice {
	case "1":
		(*people)[personIndex].Name = input("Name:")
	case "2":
		(*people)[personIndex].Email = input("Email:")
	case "3":
		(*people)[personIndex].Age = uint8(inputInt("Age:"))
	case "4":
		(*people)[personIndex].Cpf = inputCpf(*people)
	}
}

func removePerson(slice []Person, index int) []Person {
	return append(slice[:index], slice[index+1:]...)
}

func deletePerson(people []Person, cpf string) []Person {
	personIndex, err := findPerson(people, cpf)
	if err != nil {
		fmt.Println(err)
		return people
	}

	return removePerson(people, personIndex)
}
