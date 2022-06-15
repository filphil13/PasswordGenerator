package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//default attributes of passwordGenrator Request
var CHARMAX int = 64
var length int = 16
var amount int = 1

func main() {

	//List of Letters,Numbers, and Symbols
	var rootSymbols string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~`! @#$%^&*()_-+={[}]|\\:;\"'<,>.?/"

	//handles arguments
	rootSymbols = HandlingArgs(rootSymbols)

	//turns rootSymbols into rune to be able to read all symbols
	rootSymbolsRune := []rune(rootSymbols)

	for i := 0; i < amount; i++ {

		//Password creation
		password := createPassword(rootSymbolsRune)

		//output
		fmt.Println(password)
	}
}

//Handling Arguments
//Takes Args and adjusts correct rootsymbol, and length
func HandlingArgs(rootSymbols string) string {
	newRootSymbols := rootSymbols

	if len(os.Args) > 1 {
		args := os.Args[1:]

		//if first arg is number this indicates size of password
		for _, arg := range args {

			if newLength, err := strconv.Atoi(arg); err == nil {
				length = newLength

				//Disables all letters
			} else if arg == "-al" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

				//disables all lower letters
			} else if arg == "-ll" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "abcdefghijklmnopqrstuvwxyz", "")

				//disables all capital letters
			} else if arg == "-cl" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

				//disables all numbers
			} else if arg == "-n" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "1234567890", "")

				//disbales all symbols
			} else if arg == "-s" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/", "")
			}
		}
	}
	//returns new string with argument modifications
	return newRootSymbols

}

func createPassword(rootSymbolsRune []rune) string {
	password := ""

	for i := 0; i < length; i++ {

		//chooses random index of rune array and and adds it to a variable
		//keeps repeating this until password length is satisfied
		randomIndex := rand.Intn(len(rootSymbolsRune))
		password = password + string(rootSymbolsRune[randomIndex])
	}

	return password
}
