Go Tags
=======

A tags package for use with Go (golang) services

[![Build Status](https://travis-ci.org/takama/tags.png?branch=master)](https://travis-ci.org/takama/tags)
[![GoDoc](https://godoc.org/github.com/takama/tags?status.svg)](https://godoc.org/github.com/takama/tags)

### Definitions

- Tags - a strings which used for tag any object
- Non-strict Tags - a strings which match to strings in data ("new" -> "new")
- Strict Tags - all strict tags have prefix "+" for strict match ("+new")
    and "-" for strict mismatch ("-old")

### Rules

- All strict Tags applied with logical operator "AND" between each other
- All non-strict Tags applied with logical operator "OR" between all tags

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
