package main

import (
	"fmt"
	"os"
)

func main() {

	//Funciton call to read the license file
	email, _, err := readLicense()

	//IF statement if there any error reading the license file
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	//Call to main UI to start the application
	mainUI(email)
}
