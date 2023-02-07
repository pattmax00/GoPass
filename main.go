// GoPass
// Author: Maximilian Patterson
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var allowedCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`~!@#$%^&*()_+[]\\{}|;':,./<>?")

const (
	Version = "1.3.0"
	symbols = "`~!@#$%^&*()_+[]\\{}|;':,./<>?"
)

func matchArguments(args []string) string {
	// If there are no arguments
	if len(args) == 0 {
		return "No password length specified! (ex: gopass 16)"
	}

	// First argument is special, must be an integer, -v, or -h
	var size = 0 // Password length
	err := error(nil)
	if size, err = strconv.Atoi(args[0]); err == nil { // If first argument is an integer
	} else if args[0] == "-v" {
		return "GoPass version " + Version
	} else if args[0] == "-h" {
		return "GoPass - A simple password generator written in Go\n" +
			"Usage: gopass [length] [disallowed characters] [optional remove symbols -s]\n" +
			"     Example: gopass 16\n" +
			"     Example: gopass 16 -r=abc123!@#\n" +
			"     Example: gopass 16 -s\n" +
			"\nFor help (this output): gopass -h\n" +
			"For version: gopass -v\n"
	}

	for i := 1; i < len(args); i++ {
		v := args[i]
		if v == "-s" {
			removeDisallowed([]rune(symbols))
		} else if strings.HasPrefix(v, "-r=") { // If argument starts with -r=
			// Remove all characters after the = until next whitespace
			removeDisallowed([]rune(v[2:]))
		} else {
			return "Invalid argument (\"" + v + "\") supplied! (Type gopass -h for help)"
		}
	}

	if size <= 0 {
		return "No/invalid password length specified! (ex: gopass 16)"
	} else {
		return generatePassword(size)
	}
}

// Remove all disallowed characters from the allowedCharacters slice
func removeDisallowed(disallowed []rune) {
	for _, r := range disallowed {
		for i, v := range allowedCharacters {
			if v == r {
				allowedCharacters = append(allowedCharacters[:i], allowedCharacters[i+1:]...)
			}
		}
	}
}

func generatePassword(size int) string {
	// Make empty array of runes with size of size
	pass := make([]rune, size)

	// Assign every slot of pass to a random allowedCharacter
	for i := range pass {
		// Generate a random int greater than 0 and not to exceed the length of allowedCharacters
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(allowedCharacters))))
		if err != nil {
			println("Error securely generating random character!")
			return ""
		}
		pass[i] = allowedCharacters[index.Int64()]
	}

	return string(pass)
}

func main() {
	// Process arguments
	fmt.Println(matchArguments(os.Args[1:]))
}
