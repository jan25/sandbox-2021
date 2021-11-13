package main

import (
	"fmt"

	// Errors package providing primitives is not part of stable Go (v1) version yet.
	// That is why we use absolute import here.
	errors "github.com/pkg/errors"
)

func main() {
	err := errors.New("root cause")
	err1 := errors.Wrap(err, "first wrap")
	err2 := errors.Wrapf(err1, "wrap num %d", 2)

	// recursively prints UnWrap(ed) version of errors.
	// custom error types can implement UnWrap() to show here.
	fmt.Printf("STACKED ERROR - %v\n", err2)

	// more info about error messages with stackframes.
	fmt.Printf("ERROR STACK TRACE - %+v\n", err2)

	// recurse wrapped errors until no more Cause() is implemented.
	fmt.Printf("FIND MAIN ERROR - %s\n", errors.Cause(err2))
}
