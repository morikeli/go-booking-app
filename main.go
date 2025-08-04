package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings []string

func main() {
	// bookings := []string{} also works
	// bookings[0] = "Amanda" 	// add item in array at the 0-th index (first position)

	greetUsers()

	// user input
	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isRequiredTickets := validateUserInput(
			firstName, lastName, email, userTickets, remainingTickets,
		)
		// isValidCity := city == "Nairobi" || "New York"

		// verify if user tickets is more than what's available
		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, you can't book %v tickets!\n", remainingTickets, userTickets)
			continue

		} else if isValidName && isValidEmail && isRequiredTickets {
			bookTickets(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("Names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// terminate program if all tickets are booked
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			fmt.Printf("Invalid input! \n")

			if !isValidName {
				fmt.Println("First name or last name is too short!")
			}
			if !isValidEmail {
				fmt.Println("Invalid email address provided!")
			}
			if !isRequiredTickets {
				fmt.Println("Invalid number of tickets provided!")
			}
		}

	}

	city := "London"

	switch city {
	case "New York":
	// execute code for booking New York conference tickets

	case "Nairobi":
	case "Tokyo", "Hong Kong":
	case "Canberra":
	case "Brasilia":
	case "London":
	case "Mumbai":
	case "South Africa":

	}

	// fmt.Printf("Bookings array: %v \n", bookings)
	// fmt.Printf("First item in bookings array: %v\n", bookings[0])
	// fmt.Printf("Array type: %T\n", bookings)
	// fmt.Printf("Length of the bookings array: %v\n", len(bookings))

	// print out the pointer of `remainingTickets`
	// fmt.Printf("Remaining tickets: %v \n", remainingTickets)
	// fmt.Printf("Pointer of remaining tickets variable: %v \n\n", &remainingTickets)

	// fmt.Printf("User %v %v booked %v tickets. \n", firstName, lastName, userTickets)

	// fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	// fmt.Println("We have a total of ", conferenceTickets, "tickets. Remaining tickets: ", remainingTickets)
	// fmt.Println("Get your tickets here to attend.")
	//
	// types of variables
	// fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T", conferenceName, conferenceTickets, remainingTickets)

}

// functions
func greetUsers() {
	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Println("We have a total of ", conferenceTickets, "tickets. Remaining tickets: ", remainingTickets)
	fmt.Println("Get your tickets here to attend.")

}

func getUserInput() (string, string, string, uint) {

	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	// fmt.Println("What city will you attend %v?\n", conferenceName)
	// fmt.Scan(&city)

	return firstName, lastName, email, userTickets

}

// `[]string {}` is the output param - the type of data the function returns
func getFirstNames() []string {
	firstNames := []string{}

	// iterate through the firstNames array to get first names of each user
	for _, name := range bookings {
		var names = strings.Fields(name) // split the string with whitespace as separator & return a slice with the split elements
		firstNames = append(firstNames, names[0])
	}

	return firstNames

}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets: %v for %v \n", remainingTickets, conferenceName)

}
