package main

import "fmt"

// you can declare variables outside of the func main() but you cannot assign them here.
// var cardName string -> WORKS (and then you can assign the value inside of func main())
// var cardName string = "Ace of Spades" -> ERROR (Cant assign outside of func main())
// cardName := "Ace of Spades" -> ERROR (Cant assign outside of func main())

func main() {
	/*
		var -> declaring a new variable
		card -> the name of the variable will be "card"
		string -> Only a "string" type can be assigned here - similar to const in JS
			Go is a Statically typed language (so are C++ & Java);
				Statically typed are declared once and stay that way always and are expecting that data type or will throw an error
			Languages like Ruby, JS, Python are Dynamically Typed languages;
				Dynamically means that you can change what type of variable is assigned to a variable - integer can change to string
			*** NOT REQUIRED *** it'll assume that it'll always be the initial value's type
		= -> the next thing is what the variable will be
			Go basic data types - bool, string, int, float64 (there are many other types available)
	*/
	// var card string = "Ace of Spades"
	// above is long form and very explicit

	// Also note camelCase is observed

	// short form way to declare a NEW variable.  If you are reassigning, just use plain =
	card := "Ace of Spades"
	// Example of reassignment
	card = "Five of Diamonds"

	fmt.Println(card)
}

//prints out "Five of Diamonds"
