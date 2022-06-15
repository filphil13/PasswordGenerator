package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"unicode"
)

func main() {

	//assigning variables
	//default attributes of passwordGenrator Request
	CHARMAX := 64

	lettersActive := true
	numbersActive := true
	symbolsActive := true
	length := 16
	amount := 1

	//List of Letters,Numbers, and Symbols
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters = []rune(letters)
	numbers := "1234567890"
	numbers = []rune(numbers)
	symbols := "~`! @#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
	symbols = []rune(symbols)

	//Handling Arguments
	//setting what contents is unneeded in password and storing it in a string
	//setting how many chars are needed for password
	if len(os.Args) > 1 {
		args := os.Args[1:]

		//if first arg is number this indicates size of password
		for i, arg := range args {
			if i == 0 && unicode.IsDigit(arg) {
				length, err := strconv.Atoi(arg)
			}

			//Deactivate Letters
			if arg == "-l" {
				lettersActive = false
			}

			//Deactivate Numbers
			if arg == "-n" {
				lettersActive = false
			}

			//Deactivate Symbols
			if arg == "-s" {
				lettersActive = false
			}
		}
	}

	//Password creation
	var activeArgs []string 

	
	if lettersActive{	
		append(activeArgs, "letter")
	}
	if numbersActive{	
		append(activeArgs, "number")
	}
	if symbolsActive{	
		append(activeArgs, "symbol")
	}
	
	for i := 0; i < amount; i++{
		password := ""
		for i := 0; i < length; i++{
			temp := rand.Intn(len(activeArgs)) - 1

			if temp == "letter"{
				password := password + letters[rand.Intn(len(letters)]
			}else if temp == "number"{
				password := password + number[rand.Intn(len(numbers)]
			}else if temp == "symbol"{
				password := password + symbols[rand.Intn(len(symbols)]
			}
		}
		rand.Shuffle()
		//output
		fmt.Println(password)
	}
}
