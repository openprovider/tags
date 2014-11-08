// Copyright 2014 Igor Dolzhikov. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

/*
Package tags 0.1.3

Definition:

	- Tags - a strings which used for tag any object
	- Non-strict Tags - a strings which match to strings in data ("new" -> "new")
	- Strict Tags - all strict tags have prefix "+" for strict match ("+new")
	    and "-" for strict mismatch ("-old")

Rules:

	- All strict Tags applied with logical operator "AND" between each other
	- All non-strict Tags applied with logical operator "OR" between all tags

Queries:

	{"a","b"} this mean that we ask "a" OR "b" tag
	{"a","+b","+c"} this mean that we ask "a" OR ("b" AND "c") tag
	{"a","+b", "c", "-d"} this mean that we ask "a" OR "c" OR ("b" AND NOT "d") tag

Example:

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

		// We ask for any tee "black" or "green"
		fmt.Println("Is this tee black or green?")
		query := tags.Tags{"black", "green"}

		if product.Tags.IsTagged(query) {
			fmt.Println("Yes, the tee is black.")
		} else {
			fmt.Println("No, the tee has not black or green options.")
		}

		// We ask fot strict match "green" and "sugar"
		fmt.Println("Is this tee green with sugar?")
		query = tags.Tags{"+green", "+sugar"}

		if product.Tags.IsTagged(query) {
			fmt.Println("Yes, the tee is green with sugar.")
		} else {
			fmt.Println("No, the tee with sugar, but is not green.")
		}

		// We ask for strict mismatch, not "ice"
		fmt.Println("Is this tee hot?")
		query = tags.Tags{"-ice"}

		if product.Tags.IsTagged(query) {
			fmt.Println("Yes, the tee is hot.")
		} else {
			fmt.Println("No, it is ice tee.")
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
