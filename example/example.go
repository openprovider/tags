package main

import (
	"fmt"

	"github.com/openprovider/tags"
)

// Product is struct with tags
type Product struct {
	Name        string
	Description string
	Tags        tags.Tags
}

func main() {
	product := Product{
		Name:        "the tea",
		Description: "the boutle of ice black tea with sugar",
		Tags:        tags.Tags{"ice", "black", "sugar"},
	}

	fmt.Println("Product:", product.Description)

	// We ask for any tea "black" or "green"
	fmt.Println("Is this tea black or green?")
	query := tags.Tags{"black", "green"}

	if product.Tags.IsTagged(query) {
		fmt.Println("Yes, the tea is black.")
	} else {
		fmt.Println("No, the tea has not black or green options.")
	}

	// We ask fot strict match "green" and "sugar"
	fmt.Println("Is this tea green with sugar?")
	query = tags.Tags{"+green", "+sugar"}

	if product.Tags.IsTagged(query) {
		fmt.Println("Yes, the tea is green with sugar.")
	} else {
		fmt.Println("No, the tea with sugar, but is not green.")
	}

	// We ask for strict mismatch, not "ice"
	fmt.Println("Is this tea hot?")
	query = tags.Tags{"-ice"}

	if product.Tags.IsTagged(query) {
		fmt.Println("Yes, the tea is hot.")
	} else {
		fmt.Println("No, it is ice tea.")
	}
}
