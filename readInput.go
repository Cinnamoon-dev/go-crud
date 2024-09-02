package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inputCpf(people []Person) string {
	cpf := input("Cpf:")

	for {
		_, err := findPerson(people, cpf)
		if err == nil {
			cpf = input("This cpf already exists. Choose another:")
			continue
		}

		break
	}

	return cpf
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
