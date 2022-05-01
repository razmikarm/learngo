// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func main() {
	const (
		attempt = 5
		descr   = `Guess a number, enter it and win!
Computer will keep %d random numbers,
If you'll guess one of them, you'll win!
`
		usage = "Usage: [positive number]"
	)
	rand.Seed(time.Now().UnixMicro())
	if len(os.Args) != 2 {
		fmt.Printf(descr, attempt)
		fmt.Println(usage)
		return
	}
	guess, err := strconv.Atoi(os.Args[1])
	if err != nil || guess <= 0 {
		fmt.Println(usage)
		return
	}
	msgIdx := rand.Intn(10) + 1
	for i := 0; i < attempt; i++ {
		if r := rand.Intn(guess) + 2; r == guess {
			msgIdx *= -1
		}
	}
	switch {
	case msgIdx < -5:
		fmt.Println("  YOU WON")
	case msgIdx < 0:
		fmt.Println("  YOU'RE AWESOME")
	case msgIdx < 6:
		fmt.Println("  LOSER!")
	default:
		fmt.Println("  YOU LOST. TRY AGAIN")
	}
}
