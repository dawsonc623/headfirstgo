package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100) + 1

	fmt.Println("I have chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	var i int

	for i = 10; i > 0; i -= 1 {
		fmt.Print(i, " guesses left > ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("\nOops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("\nOops. Your guess was HIGH.")
		} else {
			break
		}
	}

	if i == 0 {
		fmt.Println("\nSorry, you did not guess my number. It was", target, "\b.")
	} else {
		fmt.Println("\nGood job! You guessed it!")
	}
}
