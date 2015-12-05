// Copyright 2015 Openprovider Authors. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

/*
Package tags 0.2.0

Definitions:

	- Tags - a strings which used for tag any object
	- Non-strict query tag - a string which match to string in tag ("new" -> "new")
	- Strict query tag - all strict query tags have prefix "+" for strict match ("+new")
    	and "-" for strict mismatch ("-old")

Rules:

	- All strict query tags applied with logical operator "AND" between each other
	- All non-strict query tags applied with logical operator "OR" between all tags in query

Queries:

	{"a","b"} this mean that we ask "a" OR "b" tag
	{"a","+b","+c"} this mean that we ask "a" OR ("b" AND "c") tag
	{"a","+b", "c", "-d"} this mean that we ask "a" OR "c" OR ("b" AND NOT "d") tag

Example:

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


Go Tags
*/
package tags

import (
	"strings"
)

// Tags is a slice of strings which used for tag any object
type Tags []string

// IsTagged - this method checks if Tags are matched with tags in query
func (t Tags) IsTagged(query Tags) bool {
	if len(query) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	for _, qTag := range query {
		for _, tag := range t {
			if tag == qTag {
				return true
			}
		}
	}

	var strict uint8
	for _, qTag := range query {
		if strings.HasPrefix(qTag, "+") || strings.HasPrefix(qTag, "-") {
			strict++
		}
	}
	if strict == 0 {
		return false
	}
	for _, qTag := range query {
		if strings.HasPrefix(qTag, "+") {
			for _, tag := range t {
				if tag == strings.TrimPrefix(qTag, "+") {
					strict--
					break
				}
			}
		}
		if strings.HasPrefix(qTag, "-") {
			fit := true
			for _, tag := range t {
				if tag == strings.TrimPrefix(qTag, "-") {
					fit = false
					break
				}
			}
			if fit {
				strict--
			}
		}
	}

	return strict == 0
}
