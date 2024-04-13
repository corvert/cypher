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
	}
}

func encrypt() {
	inputEncrypt := 0

	for {
		fmt.Println("")
		fmt.Println("Select cypher (1/2/3):\n 1. ROT13.\n 2. Reverse.\n 3. fun3.")

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
	}
}

func decrypt() {
	inputDecrypt := 0
	for {
		fmt.Println("Select cypher (1/2/3):\n 1. ROT13.\n 2. Reverse.\n 3. func31")
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
			fun3De()
			return
		default:
			fmt.Println("Invalid input", inputDecrypt, "try again")

		}
	}
}

func ROT13() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text for ROT13 transformation: ")
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	// fmt.Printf("You entered: %s \n", text)

	var transformed []rune
	for _, r := range text {
		switch {
		case r >= 'a' && r <= 'z':
			// Shift character within the lowercase letters
			transformed = append(transformed, 'a'+(r-'a'+13)%26)
		case r >= 'A' && r <= 'Z':
			// Shift character within the uppercase letters
			transformed = append(transformed, 'A'+(r-'A'+13)%26)
		default:
			// Non-alphabetical characters are unchanged
			transformed = append(transformed, r)
		}
	}
	result := string(transformed)
	fmt.Println("Transformed message using ROT13 cipher:", result)
}

func reverse() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text for reverse transformation: ")
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	//	fmt.Printf("You entered: %s \n", text)

	var transformed []rune
	for _, r := range text {
		if r >= 'a' && r <= 'z' {
			// Calculate the mirrored lowercase letter
			transformed = append(transformed, 'z'-(r-'a'))
		} else if r >= 'A' && r <= 'Z' {
			// Calculate the mirrored uppercase letter
			transformed = append(transformed, 'Z'-(r-'A'))
		} else {
			// Non-alphabetical characters are unchanged
			transformed = append(transformed, r)
		}
	}
	result := string(transformed)
	fmt.Println("Transformed message using reverse cipher:", result)
}

func fun3En() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text for simple substitution transformation: ")
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	// fmt.Printf("You entered: %s \n", text)

	var transformed []rune
	for _, r := range text {
		switch {
		case r >= 'a' && r <= 'z':
			// Shift character to the next within the lowercase letters
			// Wrap around from 'z' to 'a'
			if r == 'z' {
				transformed = append(transformed, 'a')
			} else {
				transformed = append(transformed, r+1)
			}
		case r >= 'A' && r <= 'Z':
			// Shift character to the next within the uppercase letters
			// Wrap around from 'Z' to 'A'
			if r == 'Z' {
				transformed = append(transformed, 'A')
			} else {
				transformed = append(transformed, r+1)
			}
		default:
			// Non-alphabetical characters are unchanged
			transformed = append(transformed, r)
		}
	}
	result := string(transformed)
	fmt.Println("Transformed message using simple substitution:", result)
}

func fun3De() {
	input := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text for simple substitution transformation: ")
	text, _ := input.ReadString('\n')
	text = strings.TrimSpace(text)
	// fmt.Printf("You entered: %s \n", text)

	var transformed []rune
	for _, r := range text {
		switch {
		case r >= 'a' && r <= 'z':
			// Shift character to the next within the lowercase letters
			// Wrap around from 'z' to 'a'
			if r == 'z' {
				transformed = append(transformed, 'a')
			} else {
				transformed = append(transformed, r-1)
			}
		case r >= 'A' && r <= 'Z':
			// Shift character to the next within the uppercase letters
			// Wrap around from 'Z' to 'A'
			if r == 'Z' {
				transformed = append(transformed, 'A')
			} else {
				transformed = append(transformed, r-1)
			}
		default:
			// Non-alphabetical characters are unchanged
			transformed = append(transformed, r)
		}
	}
	result := string(transformed)
	fmt.Println("Transformed message using simple substitution:", result)
}
