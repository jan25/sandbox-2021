package main

import (
	"fmt"
	"log"

	types "github.com/jan25/sandbox-2021/storage/types"
	memory "github.com/jan25/sandbox-2021/storage/types/memory"
)

func main() {
	var s types.Storage = memory.NewStorage()

	r := memory.NewRecord(10, "hi")
	if err := s.Write(r); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Succesfully written: %v\n", r)

	if r, err := s.Read(10); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully Read: %v\n", r)
	}
}
