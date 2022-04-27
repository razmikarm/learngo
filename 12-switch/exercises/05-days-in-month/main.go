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
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	month := os.Args[1]
	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	switch m := strings.ToLower(month); m {
	case "january", "march", "may", "july", "august", "october", "december":
		fmt.Printf("%q has 31 days.\n", m)
	case "april", "june", "september", "november":
		fmt.Printf("%q has 30 days.\n", m)
	case "february":
		if leap {
			fmt.Println("february has 29 days.")
		} else {
			fmt.Println("february has 28 days.")
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
	}

}
