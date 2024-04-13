package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println("")

	input := 0

	for {
		fmt.Println("Select operation (1/2):\n 1. Encrypt.\n 2. Decrypt.\n")
		_, err := fmt.Scan(&input)
		if err != nil {
			continue

		}

		switch input {
		case 1:
			encrypt()
			return

		case 2:
			decrypt()
			return

		default:
			fmt.Println("Invalid input", input, "try again")

		}
	}
}

func encrypt() {
	inputEncrypt := 0

	for {
		fmt.Println("")
		fmt.Println("Select cypher (1-3):\n 1. ROT13.\n 2. Reverse.\n 3. fun3.")

		_, err := fmt.Scan(&inputEncrypt)
		if err != nil {
			continue
		}

		switch inputEncrypt {
		case 1:
			ROT13()
			return
		case 2:
			reverse()
			return
		case 3:
			fun3()
			return
		default:
			fmt.Println("Invalid input", inputEncrypt, "try again")

		}
	}
}

func decrypt() {
	inputDecrypt := 0
	for {
		fmt.Println("Select cypher (1-3):\n 1. ROT13.\n 2. Reverse.\n 3. func31")
		_, err := fmt.Scan(&inputDecrypt)
		if err != nil {
			continue
		}

		switch inputDecrypt {
		case 1:
			ROT13()
			return
		case 2:
			reverse()
			return
		case 3:
			fun3()
			return
		default:
			fmt.Println("Invalid input", inputDecrypt, "try again")

		}
	}
}

func ROT13() {

	input := bufio.NewReader(os.Stdin)
	fmt.Print("You chooce ROT13\nWhat whould You like to to encrypt: ")
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("You entered: %s \n", text)

	//answer := ""
	//for i:=0; i<len(input); i++{
	//
	//}

	//fmt.Println("Decrypted message using ROT13:\n"+input)
}

func reverse() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("You chooce Reverse!\nWhat whould You like to to encrypt: ")
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("You entered: %s \n", text)
}

func fun3() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("You chooce FUNC3\nWhat whould You like to to encrypt: ")
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("You entered: %s \n", text)
}
