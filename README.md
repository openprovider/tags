Go Tags
=======

A tags package for use with Go (golang) services

[![Build Status](https://travis-ci.org/takama/tags.png?branch=master)](https://travis-ci.org/takama/tags)
[![GoDoc](https://godoc.org/github.com/takama/tags?status.svg)](https://godoc.org/github.com/takama/tags)

### Definitions

- Tags - a strings which used for tag any object
- Non-strict query tag - a string which match to string in tag ("new" -> "new")
- Strict query tag - all strict query tags have prefix "+" for strict match ("+new")
    and "-" for strict mismatch ("-old")

### Rules

- All strict query tags applied with logical operator "AND" between each other
- All non-strict query tags applied with logical operator "OR" between all tags in query

### Queries

{"a","b"} this mean that we ask "a" OR "b" tag

{"a","+b","+c"} this mean that we ask "a" OR ("b" AND "c") tag

{"a","+b", "c", "-d"} this mean that we ask "a" OR "c" OR ("b" AND NOT "d") tag

### Example

```go
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
```

## Author

[Igor Dolzhikov](https://github.com/takama)

## Contributors

All the contributors are welcome. If you would like to be the contributor please accept some rules.
- The pull requests will be accepted only in "develop" branch
- All modifications or additions should be tested
- Sorry, I'll not accept code with any dependency, only standard library

Thank you for your understanding!

## License

[MIT Public License](https://github.com/takama/tags/blob/master/LICENSE)
