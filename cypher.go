package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
			transformed = append(transformed, 'a'+(r-'a'+13)%26)
		case r >= 'A' && r <= 'Z':
			transformed = append(transformed, 'A'+(r-'A'+13)%26)
		default:
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
			transformed = append(transformed, 'z'-(r-'a'))
		} else if r >= 'A' && r <= 'Z' {
			transformed = append(transformed, 'Z'-(r-'A'))
		} else {
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
			if r == 'z' {
				transformed = append(transformed, 'a')
			} else {
				transformed = append(transformed, r+1)
			}
		case r >= 'A' && r <= 'Z':
			if r == 'Z' {
				transformed = append(transformed, 'A')
			} else {
				transformed = append(transformed, r+1)
			}
		default:
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
			if r == 'z' {
				transformed = append(transformed, 'a')
			} else {
				transformed = append(transformed, r-1)
			}
		case r >= 'A' && r <= 'Z':
			if r == 'Z' {
				transformed = append(transformed, 'A')
			} else {
				transformed = append(transformed, r-1)
			}
		default:
			transformed = append(transformed, r)
		}
	}
	result := string(transformed)
	fmt.Println("Transformed message using simple substitution:", result)
}
