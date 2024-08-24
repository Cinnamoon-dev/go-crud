package main

import (
	"bufio"
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
	newPerson.Name = input("Name:")
	newPerson.Age = uint8(inputInt("Age:"))

	return newPerson
}

func createPerson(people *[]Person) {
	newPerson := getPerson()
	*people = append(*people, newPerson)
}

func editPerson(people *[]Person) {
	// find person by cpf
	// choose what field change
	// insert new value
}

func input(message string) string {
	// Use case
	// response := input("say your name")
	// fmt.Println("Hi, " + response + "!")

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
	list := []int{1, 2, 3, 4, 5}
	list = remove(list, 1)
	showMainOptions()
}

func showMainOptions() {
	fmt.Print(`
1 - employee
2 - car
	`)
}

func mainOptions() {
	var people []Person
	var cars []Car

	_, _ = people, cars

	items :=
		`1 - employee
2 - car`

	var choice string

	for {
		fmt.Println(items)
		choice = input("choose a item")

		switch choice {
		case "1":
			fmt.Println("you chose employee!")
		case "2":
			fmt.Println("you chose car!")
		default:
			fmt.Println("none")
		}
	}
}
