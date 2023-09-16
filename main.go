package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go conference" // alternate syntax for (conferenceName := "Go conference")
var remainingTickets uint = 50       //uint accpets positive integers only
var bookings = make([]UserData, 0)

// var bookings = make([]map[string]string, 0) // creating an empty list of maps
// var bookings = []string{}            //using slices
// var bookings := []string{} alternate syntax for slices
// var bookings [50]string using array

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName) logic to find out data type of different variables

	firstName, lastName, email, userTickets := getUserInput()
	// isValidCity := city == "Singapore" || city =="London"
	// isInvalidCity := city!= "Singapore" && city!="London" equals !isValidCity

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		bookTicket(userTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		//calling first name function
		printFirstNames()

		if remainingTickets == 0 {
			//end of program
			fmt.Println("Our conference is booked, come back next year")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("The first or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("The email you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}

	}
	wg.Wait()

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings { // old syntax "for index, booking := range bookings" but we replaced "index" with "_" as we were not using it which caused an error but we can't just remove it either so to ignore "index" we placed "_" there.
		firstNames = append(firstNames, booking.firstName)
	}
	fmt.Printf("The first names of the bookings are %v\n", firstNames)

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// asks user to enter name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName) //here we have used pointer to point at the memory address of firstName to get input from the user

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName) //here we have used pointer to point at the memory address of lastName to get input from the user

	fmt.Println("Enter your email:")
	fmt.Scan(&email) //here we have used pointer to point at the memory address of email to get input from the user

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets) //here we have used pointer to point at the memory address of userTickets to get input from the user

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// fmt.Printf("The whole slice: %v.\n", bookings)
	// fmt.Printf("The first value : %v.\n", bookings[0])
	// fmt.Printf("The type of slice: %T.\n", bookings)
	// fmt.Printf("The lenght of slice: %v.\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will get a confirmation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n ", userTickets, firstName, lastName)
	fmt.Println("**********************")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("**********************")
	wg.Done()

}
