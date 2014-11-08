package main

import (
	"fmt"

	"github.com/takama/tags"
)

// Product is struct with tags
type Product struct {
	Name        string
	Description string
	Tags        tags.Tags
}

func main() {
	product := Product{
		Name:        "the tee",
		Description: "the boutle of ice black tee with sugar",
		Tags:        tags.Tags{"ice", "black", "sugar"},
	}

	fmt.Println("Product:", product.Description)

	fmt.Println("Is this tee black or green?")
	query := tags.Tags{"black", "green"}

	if product.Tags.IsTagged(query) {
		fmt.Println("Yes, the tee is black.")
	} else {
		fmt.Println("No, the tee has not black or green options.")
	}

	fmt.Println("Is this tee green with sugar?")
	query = tags.Tags{"+green", "+sugar"}

	if product.Tags.IsTagged(query) {
		fmt.Println("Yes, the tee is green with sugar.")
	} else {
		fmt.Println("No, the tee with sugar, but is not green.")
	}

	fmt.Println("Is this tee hot?")
	query = tags.Tags{"-ice"}

	if product.Tags.IsTagged(query) {
		fmt.Println("Yes, the tee is hot.")
	} else {
		fmt.Println("No, it is ice tee.")
	}
}
