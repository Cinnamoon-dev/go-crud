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

func findPerson(people *[]Person, cpf string) (*Person, error) {
	// Use case
	// person1, err := findPerson(&people, "1234")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	for i := 0; i < len(*people); i++ {
		if (*people)[i].Cpf == cpf {
			var person *Person = &(*people)[i]
			return person, nil
		}
	}

	return nil, errors.New("Person not found")
}

func editPerson(people *[]Person, cpf string) {
	person, err := findPerson(people, cpf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// TODO
	// Check if cpf already exists when changing
	choice := input("What field do you want do change?\n1 - Name\n2 - Email\n3 - Age\n4 - Cpf")

	switch choice {
	case "1":
		person.Name = input("Name:")
	case "2":
		person.Email = input("Email:")
	case "3":
		person.Age = uint8(inputInt("Age:"))
	case "4":
		person.Cpf = input("Cpf:")
	}
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

func main() {
	var people []Person
	var cars []Car
	_ = cars

	fmt.Println(people)

	createPerson(&people)
	createPerson(&people)

	fmt.Println(people)

	editPerson(&people, "1234")

	fmt.Println(people)

}
