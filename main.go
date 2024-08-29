package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	// "strconv"
)

type Car struct {
	Plate     string
	Color     string
	Model     string
	Brand     string
	PersonCpf string
}

type Person struct {
	Name  string
	Email string
	Cpf   string
	Age   uint8
}

func getPerson() Person {
	var newPerson Person

	newPerson.Name = input("Name:")
	newPerson.Email = input("Email:")
	newPerson.Cpf = input("Cpf:")
	newPerson.Age = uint8(inputInt("Age:"))

	return newPerson
}

func createPerson(people *[]Person) {
	newPerson := getPerson()
	*people = append(*people, newPerson)
}

func viewAllPerson(people []Person) {
	for i := 0; i < len(people); i++ {
		fmt.Println(people[i])
	}
}

func viewPerson(people *[]Person, cpf string) {
	for _, v := range *people {
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

	// TODO
	// Check if cpf already exists when changing
	choice := input("What field do you want do change?\n1 - Name\n2 - Email\n3 - Age\n4 - Cpf")

	switch choice {
	case "1":
		(*people)[personIndex].Name = input("Name:")
	case "2":
		(*people)[personIndex].Email = input("Email:")
	case "3":
		(*people)[personIndex].Age = uint8(inputInt("Age:"))
	case "4":
		(*people)[personIndex].Cpf = input("Cpf:")
	}
}

func deletePerson(people []Person, cpf string) []Person {
	personIndex, err := findPerson(people, cpf)
	if err != nil {
		fmt.Println(err)
		return people
	}

	return removePerson(people, personIndex)
}

func input(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message, "\n>> ")

	text, _ := reader.ReadString('\n')

	return strings.Replace(text, "\n", "", -1)
}

func inputInt(message string) int {
	var newInt int
	fmt.Print(message, "\n>> ")

	fmt.Scanf("%d\n", &newInt)
	return newInt
}

func remove(slice []int, index int) []int {
	// Use case
	// list := []int{1, 2, 3, 4, 5}
	// list = remove(list, 1)
	return append(slice[:index], slice[index+1:]...)
}

func removePerson(slice []Person, index int) []Person {
	// Use case
	// list := []int{1, 2, 3, 4, 5}
	// list = remove(list, 1)
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	var people []Person
	var choice string

	for {
		choice = input("1 - create\n2 - read\n3 - edit\n4 - delete\n5 - exit")
		switch choice {
		case "1":
			createPerson(&people)
		case "2":
			viewAllPerson(people)
		case "3":
			personCpf := input("Cpf of the person")
			editPerson(&people, personCpf)
		case "4":
			personCpf := input("Cpf of the person")
			people = deletePerson(people, personCpf)
		case "5":
			return
		}
	}
}
