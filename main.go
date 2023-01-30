// GoPass
// Author: Maximilian Patterson
// Version: 1.2.1
package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"os"
	"strconv"
)

var allowedCharacters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`~!@#$%^&*()_+[]\\{}|;':,./<>?")

func main() {
	// Take in all OS arg
	args := os.Args[1:]
	if len(args) < 1 {
		println("No password length specified! (ex: ./gopass 16)")
		return
	}

	// Convert String arg to int
	size, err := strconv.Atoi(args[0])
	if err != nil {
		println("First argument supplied must be an integer! (ex: 16)")
		return
	}

	// Grab second argument (if it exists) and use it as a disallowed character(s)
	var disallowed []rune
	if len(args) == 2 {
		// Break apart the string into a slice of runes
		disallowed = []rune(args[1])

		// Remove all disallowed characters from the allowedCharacters slice
		for _, r := range disallowed {
			for i, v := range allowedCharacters {
				if v == r {
					allowedCharacters = append(allowedCharacters[:i], allowedCharacters[i+1:]...)
				}
			}
		}
	}

	// Make empty array of runes with size of size
	pass := make([]rune, size)

	// Seed rand with time
	var b [8]byte
	_, err = cryptorand.Read(b[:])
	if err != nil {
		println("Error securely seeding crypto/rand!")
		return
	}
	mathrand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	// Assign every slot of pass to a random allowedCharacter

	for i := range pass {
		// Generate a random int greater than 0 and not to exceed the length of allowedCharacters
		index, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(len(allowedCharacters))))
		if err != nil {
			println("Error securely generating random character!")
			return
		}
		pass[i] = allowedCharacters[index.Int64()]
	}

	// Print the password
	fmt.Println(string(pass))
}
