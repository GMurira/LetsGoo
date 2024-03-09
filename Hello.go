package main

import "fmt"

func main() {

	fmt.Println("Welcome to my Quiz Game.")

	var name string
	fmt.Printf("Enter your name: \n")
	fmt.Scan(&name)
	fmt.Printf("Welcome to the game %v\n", name)
	var age uint
	fmt.Printf("Enter your age %v:", name)
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Printf("Yeeeah!, You can Play")
	} else {
		fmt.Printf("Sorry your too young %v", name)
		return
	}
	//Variable to keep track of the game score
	Score := 0
	num_questions := 2
	fmt.Printf("What is better, the RTX  3080 or RTX 3090")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer+" "+answer2 == "RTX 3080" || answer+" "+answer2 == "rtx 3080" {
		fmt.Printf("Correct")
		Score++
	} else {
		fmt.Printf("Icorrect")
	}

	fmt.Printf("How many cores does the Ryzen 9 have?")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Printf("Correct")
		Score++
	} else {
		fmt.Printf("Correct")
	}

	fmt.Printf("You scored %v out of %v.\n", Score, num_questions)

	//Percentage Score
	percent := (float64(Score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%", percent)
}
