package main

import (
	"fmt"
	"github.com/LoliE1ON/srakes-bot/files"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// Create files
	err := files.Create()
	if err != nil {
		fmt.Println("main", err)
		return
	}

	// Get files
	f, err := files.Get()
	if err != nil {
		return
	}
	fmt.Println("Hello", f)

}
