package main

import (
	"fmt"
	"log"
)

func main() {
	var s Storage = NewMemoryStorage()

	r := NewMemoryRecord(10, "hi")
	if err := s.write(r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Succesfully written: %v\n", r)

	if r, err := s.read(10); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully read: %v\n", r)
	}
}
