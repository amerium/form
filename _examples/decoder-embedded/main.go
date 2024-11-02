package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/amerium/form/v6"
)

// A ...
type A struct {
	Field string
}

// B ...
type B struct {
	A
	Field string
}

// use a single instance of Decoder, it caches struct info
var decoder *form.Decoder[any]

func main() {
	decoder = form.NewDecoder[any]()

	// this simulates the results of http.Request's ParseForm() function
	values := parseFormB()

	var b B

	// must pass a pointer
	err := decoder.Decode(&b, values, nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", b)

	values = parseFormAB()

	// must pass a pointer
	err = decoder.Decode(&b, values, nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v\n", b)
}

// this simulates the results of http.Request's ParseForm() function
func parseFormB() url.Values {
	return url.Values{
		"Field": []string{"B FieldVal"},
	}
}

// this simulates the results of http.Request's ParseForm() function
func parseFormAB() url.Values {
	return url.Values{
		"Field":   []string{"B FieldVal"},
		"A.Field": []string{"A FieldVal"},
	}
}
