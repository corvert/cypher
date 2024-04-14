package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ROT13() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Print("Enter text for ROT13 transformation:\n")
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("No input detected, try again.")
			fmt.Print("Enter text for ROT13 transformation:\n")
			continue
		}

		var transformed []rune
		for _, r := range input {
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
		fmt.Println("")
		fmt.Print("Transformed message using ROT13 cipher:\n")
		fmt.Println(result)
		return
	}
}

func reverse() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Print("Enter text for reverse transformation:\n")
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("No input detected, try again.")
			fmt.Print("Enter text for reverse transformation:\n")
			continue
		}

		var transformed []rune
		for _, r := range input {
			if r >= 'a' && r <= 'z' {
				transformed = append(transformed, 'z'-(r-'a'))
			} else if r >= 'A' && r <= 'Z' {
				transformed = append(transformed, 'Z'-(r-'A'))
			} else {
				transformed = append(transformed, r)
			}
		}
		result := string(transformed)
		fmt.Println("")
		fmt.Print("Transformed message using reverse cipher:\n")
		fmt.Println(result)
		return
	}
}

func simpleSubstitutionEn() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Print("Enter text for simple substitution transformation:\n")
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("No input detected, try again.")
			fmt.Print("Enter text for simple substitution transformation:\n")
			continue
		}

		var transformed []rune
		for _, r := range input {
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
		fmt.Println("")
		fmt.Print("Transformed message using simple substitution:\n")
		fmt.Println(result)
		return
	}
}

func simpleSubstitutionDe() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("")
	fmt.Print("Enter text for simple substitution transformation:\n")
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("No input detected, try again.")
			fmt.Print("Enter text for simple substitution transformation:\n")
			continue
		}

		var transformed []rune
		for _, r := range input {
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
		fmt.Println("")
		fmt.Print("Transformed message using simple substitution:\n")
		fmt.Println(result)
		return
	}
}
