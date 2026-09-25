package main

import (
	"fmt";
	"os";
	"totp/code"
	
)

func main () {
	fmt.Println("Working")
	if len(os.Args) == 1 {
		fmt.Println("Help:")
	}
	if len(os.Args) == 2{
		Key := os.Args[1]
		fmt.Print("Code:")
		code.GetCode(Key)
	}
}