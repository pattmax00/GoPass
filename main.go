// GoPass
// Author: Maximilian Patterson
// Version: 1.1
package main

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"fmt"
	mathrand "math/rand"
	"os"
	"strconv"
)

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`~!@#$%^&*()_+[]\\{}|;':,./<>?")

func main() {
	// Take OS arg (only accepts one arg)
	arg := os.Args[1]

	// Convert String arg to int
	size, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}

	// Make empty array of runes with size of size
	pass := make([]rune, size)

	// Seed rand with time
	var b [8]byte
	_, err = cryptorand.Read(b[:])
	if err != nil {
		panic("Error securely seeding rand!")
	}
	mathrand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	// Assign every slot of pass to a random rune (generate rand int of length runes to select)
	for i := range pass {
		pass[i] = runes[mathrand.Intn(len(runes))]
	}

	// Print the pass :D
	fmt.Println(string(pass))
}
