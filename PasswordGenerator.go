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
var length int = 16
var amount int = 1

//List of Letters,Numbers, and Symbols
var rootSymbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~`! @#$%^&*()_-+={[}]|\\:;\"'<,>.?/"

//Takes Args and adjusts correct rootsymbol, and length

func main() {

	var rootSymbols string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890~`! @#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
	//Handling Arguments
	//setting what contents is unneeded in password and storing it in a string
	//setting how many chars are needed for password

	//Password creation
	rootSymbols = HandlingArgs(rootSymbols)
	rootSymbolsRune := []rune(rootSymbols)

	fmt.Println("Hello")
	for i := 0; i < amount; i++ {
		password := createPassword(rootSymbolsRune)
		//output
		fmt.Println(password)
	}
}

func HandlingArgs(rootSymbols string) string {
	newRootSymbols := rootSymbols

	if len(os.Args) > 1 {
		args := os.Args[1:]

		//if first arg is number this indicates size of password
		for _, arg := range args {

			if newLength, err := strconv.Atoi(arg); err == nil {
				length = newLength

			} else if arg == "-al" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

			} else if arg == "-ll" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "abcdefghijklmnopqrstuvwxyz", "")

			} else if arg == "-cl" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

			} else if arg == "-n" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "1234567890", "")

			} else if arg == "-s" {
				newRootSymbols = strings.ReplaceAll(newRootSymbols, "~`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/", "")
			}
		}
	}
	return newRootSymbols

}

func createPassword(rootSymbolsRune []rune) string {
	password := ""

	for i := 0; i < length; i++ {

		randomIndex := rand.Intn(len(rootSymbolsRune))
		password = password + string(rootSymbolsRune[randomIndex])
	}

	return password
}
