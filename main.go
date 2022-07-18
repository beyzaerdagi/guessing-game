package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	secret := getRandomNumber()

	for isFound := false; !isFound; {
		guessNum := guessNumber()
		isFound = isMatching(secret, guessNum)
	}

}

func getRandomNumber() int32 {
	rand.Seed(time.Now().Unix())
	return int32(rand.Int() % 1000)
}

func guessNumber() int32 {

	fmt.Println("Please enter your guess: ")
	var guessNum int32
	fmt.Scanln(&guessNum)
	fmt.Println("Your guess: ", guessNum)
	return guessNum
}

func isMatching(secret, guess int32) bool {
	if guess > secret {
		fmt.Println("Your guess too big!")
		return false
	} else if guess < secret {
		fmt.Println("Your guess too small!")
		return false
	} else {
		fmt.Println("You got it!")
		return true
	}
}
