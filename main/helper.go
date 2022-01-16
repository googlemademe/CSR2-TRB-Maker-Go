package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

//Function to check if the end date is before the start date
func checkDate(startDate time.Time, endDate time.Time) bool {
	return startDate.Before(endDate)
}

//Function to check the use inputted date format is valid
func validateDate(checkDate string) bool {
	//regex expression for valid date format
	re := regexp.MustCompile(`^(\d\d\d\d)\/(0?[1-9]|1[0-2])\/(0?[1-9]|[12][0-9]|3[01])$`)

	return re.MatchString(checkDate)
}

//Function the print out Today's date
func printTodayDate() {
	fmt.Printf("Today's Date: %d/%02d/%02d\n", time.Now().Year(), int(time.Now().Month()), time.Now().Day())
}

//Function for various print out calls
func printEnterDate(typeOfDate string) {

	//Switch statement for various print out calls
	switch typeOfDate {
	case "End":
		fmt.Printf("Enter %s Date (ex. 2009/01/01): ", typeOfDate)
	case "Start":
		fmt.Printf("Enter %s Date (ex. 2009/01/01): ", typeOfDate)
	default:
		return
	}
}

//Function for various print out based upon user inputted option
func printOutput(menuOption int, code string, endDate string) {

	var startEnd string

	//Switch for various print out based upon user inputted option
	switch code {
	case "endDate":
		fmt.Printf("End Date: %s\n", endDate)
	case "startDate":
		fmt.Printf("Start Date: %s\n", endDate)
	case "inputErorr":
		fmt.Printf("Invalid input - %s Try again...\n", endDate)
	case "validErorr":
		fmt.Printf("Invalid date - %s Try again...\n", endDate)
	case "lessDate":
		if menuOption == 2 {
			startEnd = "Start"
		}
		startEnd = "End"
		fmt.Printf("The %s Date: %s cannot be less the Start Date: %d/%02d/%02d - Try again...\n", startEnd, endDate, time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	default:
		return
	}
}

//Function to detects if the required directories - "Original", "Decrypted", and "Finished" folder - exists.
//If not, they are created
func checkIfDirectoryExist(path string) {
	//IF statement to check if the folder exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//If does not exist, creates the folder
		err := os.Mkdir(path, 0755)
		//Any error that is present will exit the program
		if err != nil {
			fmt.Printf("\nFatal Error while creating %s\n...", path)
			fmt.Println("Exiting...")
			fmt.Println(err)
			os.Exit(-1)
		}
	}
}

//Function to delete "TRBMaker" folder
func deleteAllFiles(folderName string) {

	//OS call to remove the folder
	err := os.RemoveAll(folderName)

	//IF statement to check if there is any exception and exit the program
	if err != nil {
		fmt.Println("\nFatal Error while deleting files...")
		fmt.Println("Exiting...")
		fmt.Println(err)
		os.Exit(3)
	} else {
		fmt.Printf("\nSucessfully delete all files in %s\n\n", folderName)
	}
}

//Function to delete "TRBMaker" folder
func deleteFiles() {

	//Get the current working directory
	currentDIR, err := os.Getwd()

	//IF statement to check if there is any exception and exit the program
	if err == nil {
		folderName := filepath.FromSlash(currentDIR + "/TRBMaker/")
		//Call to remove the files
		deleteAllFiles(folderName)
	}
}

//Function to call TRB making function
func makeTRB(startDate time.Time, endDate time.Time) {

	//Function call to start making the TRB files with user inputted start and end date
	createTRB(startDate, endDate, getSeasonRestArray(startDate, endDate))
}

//Function to get the list of Season Reset dates
func getSeasonRestArray(startDate time.Time, endDate time.Time) []time.Time {

	//Slice to hold all the Season Reset dates from the user inputted start date
	tempTimeArray := []time.Time{}

	//Time variable for the start date of Epoch time - 1970/01/01
	start := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	//FOR loop that start from start date to end date by incrementing 14 days
	for d := start; !d.After(endDate); d = d.AddDate(0, 0, 14) {

		//IF statement to check if the calculate date is before the start date
		if startDate.Before(d) {
			//If a match is found it is added to the temporary date slice
			tempTimeArray = append(tempTimeArray, d)
		}
	}

	//return the temporary date slice for all the calculated season reset dates
	return tempTimeArray
}

//Function to check if the current date for making the TRB is a season rest date
func checkSeasonRestDay(checkDay time.Time, seasonResetDayArray []time.Time) bool {

	//FOR loop to check if the date of TRB is a season rest day or not
	for _, j := range seasonResetDayArray {
		if j == checkDay {
			return true
		}
	}
	return false
}

//Function to print out various messages
func printMessages(show int) {

	//SWITCH statement based on user action
	switch show {
	case 1:
		fmt.Println()
		fmt.Println("Deleting previous TRBMaker folder.....")
	case 2:
		fmt.Println("Making the TRB files....")
		fmt.Println("Please Wait....")
	default:
		fmt.Println("Successfully Completed making the TRB files....")
		fmt.Println("Thank you and have a nice day....")
	}
}
