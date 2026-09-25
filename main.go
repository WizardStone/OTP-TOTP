package main

import (
	"fmt";
	"os";
	"totp/totp"
	
)

func main () {
	fmt.Println("Working")
	if len(os.Args) == 1 {
		fmt.Println("Help:")
	}
	if len(os.Args) == 2{
		Key := os.Args[1]
		fmt.Print("Code:")
		totp.GetCode(Key)
	}
}