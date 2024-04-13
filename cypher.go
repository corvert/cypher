package main

import (
	
	"fmt"

)

func main(){
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println("")
	fmt.Println("Select operation (1/2):\n 1. Encrypt.\n 2. Decrypt.")
	input := 0
     fmt.Scan(&input)

	switch input{
	case 1:
		encrypt()
	case 2:
		decrypt()
	default:
		fmt.Println("Invalid input", input,  "try again")
		input = 0
		return

		}
	}
		
func encrypt(){
	fmt.Println("Select cypher (1/2):\n 1. ROT13.\n 2. Reverse.")
	inputEncrypt := 0
     fmt.Scan(&inputEncrypt)

	switch inputEncrypt{
	case 1:
		ROT13()
	case 2:
		reverse()
	case 3:
		fun3()
	default:
		fmt.Println("Invalid input", inputEncrypt,  "try again")

		}
	}


func decrypt(){
	fmt.Println("Select cypher (1/2):\n 1. ROT13.\n 2. Reverse.")
	inputDecrypt := 0
	fmt.Scan(&inputDecrypt)

	switch inputDecrypt{
	case 1:
		ROT13()
	case 2:
		reverse()
	case 3:
		fun3()
	default:
		fmt.Println("Invalid input", inputDecrypt,  "try again")

		}
}

func ROT13(){
	fmt.Println("ROT13")
}

func reverse(){
	fmt.Println("Reverse")
}

func fun3(){
	fmt.Println("fun3")
}