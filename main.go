package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println("")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Select operation (1/2):\n1. Encrypt.\n2. Decrypt.")
		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("No input detected, try again.")
				continue
			}

			option, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input, please enter a number (1 or 2).")
				continue
			}

			switch option {
			case 1:
				encrypt()
				return
			case 2:
				decrypt()
				return
			default:
				fmt.Println("Invalid input", option, "please enter a number (1 or 2).")
			}
		}
	}
}

/* 	input := 0

for {
	fmt.Println("Select operation (1/2):\n 1. Encrypt.\n 2. Decrypt.")
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
}*/

func encrypt() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println("Select cypher (1/2/3):\n1. ROT13.\n2. Reverse.\n3. Simple Substitution.")
		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("No input detected, try again.")
				continue
			}

			option, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input, please enter a number (1 or 2 or 3).")
				continue
			}

			switch option {
			case 1:
				ROT13()
				return
			case 2:
				reverse()
				return
			case 3:
				simpleSubstitutionEn()
				return
			default:
				fmt.Println("Invalid input", option, "please enter a number (1 or 2 or 3).")
			}

		}
	}
	
}	

/* inputEncrypt := 0

for {
	fmt.Println("")
	fmt.Println("Select cypher (1/2/3):\n1. ROT13.\n2. Reverse.\n3. fun3.")

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
		fun3En()
		return
	default:
		fmt.Println("Invalid input", inputEncrypt, "try again")

	}
} */

func decrypt() {
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("")
		fmt.Println("Select cypher (1/2/3):\n1. ROT13.\n2. Reverse.\n3. Simple Substitution.")
		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			if input == "" {
				fmt.Println("No input detected, try again.")
				continue
			}

			option, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input, please enter a number (1 or 2 or 3).")
				continue
			}

			switch option {
			case 1:
				ROT13()
				return
			case 2:
				reverse()
				return
			case 3:
				simpleSubstitutionDe()
				return
			default:
				fmt.Println("Invalid input", option, "please enter a number (1 or 2 or 3).")
			}

		}
	}

}

