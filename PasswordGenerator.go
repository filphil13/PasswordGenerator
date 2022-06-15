package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//assigning variables
//default attributes of passwordGenrator Request

var CHARMAX int = 64
var lettersActive bool = true
var numbersActive bool = true
var symbolsActive bool = true
var length int = 16
var amount int = 1

//List of Letters,Numbers, and Symbols
//letters = [:26]
//nums = [26:36]
//symbols = [36:]
var rootSymbols = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~`! @#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
var letters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numbers string = "1234567890"
var symbols string = "~`! @#$%^&*()_-+={[}]|\\:;\"'<,>.?/"

//Takes Args and adjusts correct rootsymbol, and length

func main() {

	//Handling Arguments
	//setting what contents is unneeded in password and storing it in a string
	//setting how many chars are needed for password

	//Password creation
	HandlingArgs()
	fmt.Println("Hello")
	for i := 0; i < amount; i++ {
		password := ""
		rootSymbolsRune := []rune(rootSymbols)

		for i := 0; i < length; i++ {

			randomIndex := rand.Intn(len(rootSymbolsRune))
			password = password + string(rootSymbolsRune[randomIndex])
		}
		//output
		fmt.Println(password)
	}
}

func HandlingArgs() {

	if len(os.Args) > 1 {
		args := os.Args[1:]

		//if first arg is number this indicates size of password
		for _, arg := range args {

			if newLength, err := strconv.Atoi(arg); err == nil {
				length = newLength

			} else if arg == "-l" {
				lettersActive = false
				rootSymbols = strings.ReplaceAll(rootSymbols, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

			} else if arg == "-n" {
				numbersActive = false
				rootSymbols = strings.ReplaceAll(rootSymbols, "1234567890", "")

			} else if arg == "-s" {
				symbolsActive = false
				rootSymbols = strings.ReplaceAll(rootSymbols, "~`! @#$%^&*()_-+={[}]|\\:;\"'<,>.?/", "")
			}

			//Deactivate Numbers

			//Deactivate Symbols

			//Deactivate Letters
			//remove letters from

		}
	}
}
