package main

import (
	"fmt";
	"os";
	
)

func main () {
	fmt.Println("Working")
	if len(os.Args) == 1 {
		fmt.Println("Help:")
	}
}