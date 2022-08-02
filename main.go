package GoPass

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
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
	rand.Seed(time.Now().UTC().UnixNano())

	// Assign every slot of pass to a random rune (generate rand int of length runes to select)
	for i := range pass {
		pass[i] = runes[rand.Intn(len(runes))]
	}

	// Print the pass :D
	fmt.Println(string(pass))
}
