package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

//Function for Main UI of the program
func mainUI(email string) {

	//variable for reading user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		//Main UI menu
		fmt.Println("************************************************************************")
		fmt.Println("               Automated CSR TRBMaker by AppleCoreOne")
		fmt.Println()
		fmt.Println("How to use:")
		fmt.Println("Choose from the given menu options")
		fmt.Println("Have fun with this")
		fmt.Println()
		fmt.Println("Licensed to: ", email)
		fmt.Print("************************************************************************")
		fmt.Println()
		fmt.Println()
		fmt.Println("1 - Automated from \"Today's date\" to \"End Date\"")
		fmt.Println("2 - Automated from \"Start date\" to \"End Date\"")
		fmt.Println("3 - Delete all files and folders")
		fmt.Println("4 - Exit")
		fmt.Println()
		fmt.Print("Type your choice: ")
		//scans in user input
		scanner.Scan()

		//converts user input from a string to an int
		usrInput, _ := strconv.Atoi(scanner.Text())

		//Switch statement determines which action the user wants to execute
		switch usrInput {
		case 1:
			//Variable to store the user input
			var endDate time.Time

			//setting the time location to "UTC"
			loc, _ := time.LoadLocation("UTC")

			//Screen UI formatting
			fmt.Println()

			//Open FOR loop to get correct user input from the user
			for {
				//Prints out the current date to screen for user information
				printTodayDate()

				//UI prompt for user input
				printEnterDate("End")

				//Std.IO buffer to scan in user input
				scanner.Scan()

				//Set user input to a variable
				scanDate := scanner.Text()

				//Prints out user input to the screen for visual verification
				printOutput(usrInput, "endDate", scanDate)

				//IF statement validate user input format
				if !validateDate(scanDate) {
					//Prints out user input error to the screen
					printOutput(usrInput, "inputErorr", scanDate)
				} else {
					//Create a timeDate object from the user input
					endDateTemp, err := time.Parse("2006/01/02", scanDate)

					//IF statment check if the user inputted date is in the past
					//If not, it set the user input to end date of the TRB maker
					if err == nil {
						if !checkDate(time.Now(), endDateTemp) {
							printOutput(usrInput, "lessDate", scanDate)
						} else {
							endDate = endDateTemp
							break
						}
					} else {
						printOutput(usrInput, "validErorr", scanDate)
					}
				}
			}

			//Function call for displaying a message for deleting previous TRB folder
			//to make way for the new folder
			printMessages(1)

			//Function call to delete the previous TRB folder
			deleteFiles()

			//Function call for displaying a message the program will start making
			//the desired TRB from user inputs
			printMessages(2)

			//Function call to make the TRB files
			makeTRB(time.Now().In(loc), endDate)

			//Function call for displaying a message the program has finished making
			//the desired TRB from user inputs and the profrom will now close
			printMessages(3)

			//OS call to close the program successfully
			os.Exit(0)
		case 2:

			//Variables to store the user inputs
			var startDate, endDate time.Time

			fmt.Println()

			//Open FOR loop to get correct user input from the user
			for {
				//UI prompt for user input for the start date of the TRN
				printEnterDate("Start")

				//Std.IO buffer to scan in user input
				scanner.Scan()

				//Set user input to a variable
				scanDate := scanner.Text()

				//Prints out user input to the screen for visual verification
				printOutput(usrInput, "startdate", scanDate)

				//IF statement validate user input format
				if !validateDate(scanDate) {
					//Prints out user input error to the screen and continues to asked user
					//for correct input
					printOutput(usrInput, "inputErorr", scanDate)
					continue
				} else {
					//Create a timeDate object from the user input
					dateTemp, err := time.Parse("2006/01/02", scanDate)

					//IF statment check if the user inputted date is in the past
					//If not, it set the user input to start date of the TRB maker
					if err == nil {
						if !checkDate(time.Now(), dateTemp) {
							printOutput(usrInput, "lessDate", scanDate)
						} else {
							startDate = dateTemp
							break
						}
					} else {
						printOutput(usrInput, "validErorr", scanDate)
						continue
					}
				}
			}

			for {
				//UI prompt for user input for the end date of the TRN
				printEnterDate("End")

				//Std.IO buffer to scan in user input
				scanner.Scan()

				//Set user input to a variable
				scanDate := scanner.Text()

				//Prints out user input to the screen for visual verification
				printOutput(usrInput, "enddate", scanDate)

				//IF statement validate user input format
				if !validateDate(scanDate) {
					//Prints out user input error to the screen and continues to asked user
					//for correct input
					printOutput(usrInput, "inputErorr", scanDate)
					continue
				} else {
					//Create a timeDate object from the user input
					dateTemp, err := time.Parse("2006/01/02", scanDate)

					//IF statment check if the user inputted date is in the past
					//If not, it set the user input to start date of the TRB maker
					if err == nil {
						if !checkDate(startDate, dateTemp) {
							printOutput(usrInput, "lessDate", scanDate)
						} else {
							endDate = dateTemp
							break
						}
					} else {
						printOutput(usrInput, "validErorr", scanDate)
					}
				}
			}

			//Function call for displaying a message for deleting previous TRB folder
			//to make way for the new folder
			printMessages(1)

			//Function call to delete the previous TRB folder
			deleteFiles()

			//Function call for displaying a message the program will start making
			//the desired TRB from user inputs
			printMessages(2)

			//Function call to make the TRB files
			makeTRB(startDate, endDate)

			//Function call for displaying a message the program has finished making
			//the desired TRB from user inputs and the profrom will now close
			printMessages(3)

			//OS call to close the program successfully
			os.Exit(0)
		case 3:
			//Functional call to delete the previous TRB folder
			deleteFiles()
			continue
		case 4:
			//Screen messages that the user wants to exit the program
			fmt.Println("Exiting...... Program......")
			fmt.Println("Thank you......")

			//OS call to close the program successfully
			os.Exit(0)
		default:
			//Wrong user input - the program will exit with an error
			fmt.Println("Invalid input...... Exiting......")
			os.Exit(3)

		}
	}
}
