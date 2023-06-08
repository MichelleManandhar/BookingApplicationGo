package main

import "fmt"

// import (
// 	"booking-app/helper"
// 	"fmt"
// 	"time"
// )

// //inside function is local, package outside the functions and global is first letter capital and declared outside all the the finctions

// // Package variable but cannot be denoted by :=
// // conferenceName := "Go Conference" // cant be done for the constant value or where u delare the variable type like uint
// const conferenceTickets = 50

// var conferenceName = "Go Conference"
// var remainingTickets uint = 50

// // var bookings []string //using the slice or use bookings := []string{}
// // var bookings = make([]map[string]string, 0) // for the map 0 or 1 the inital size we can edit and add more to it
// var bookings = make([]UserData, 0)

// type UserData struct {
// 	firstName       string
// 	lastName        string
// 	email           string
// 	numberOfTickets uint
// }

// func main() {
// 	//greetUsers(conferenceName, conferenceTickets, remainingTickets)
// 	greetUsers() //removed parameters as we have package variable now

// 	//for remainingTickets > 0 && len(bookings) < 50 { //len counts from zero to len()-1
// 	for {
// 		firstName, lastName, email, userTickets := getUserInput()

// 		//Validation of user inputs
// 		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

// 		//isValidCity := city == "Kathmandu" || city == "London"
// 		//isInvalidCity := city != "Kathmandu" && city != "London"

// 		if isValidName && isValidEmail && isValidTicketNumber {
// 			//userName = "Michelle"
// 			//userTickets = 2
// 			//fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)
// 			bookTicket(userTickets, firstName, lastName, email)
// 			go sendTicket(userTickets, firstName, lastName, email) //multithreading and concurrency

// 			//userDataDetail := bookTicket(userTickets, firstName, lastName, email)
// 			//bookings = append(bookings, userDataDetail)

// 			//fmt.Println("bookkkkkkings ===\n", bookings)

// 			//call funcion for the name
// 			firstNames := getFirstNames()
// 			fmt.Printf("The first name of the people who booked the tickets are: %v\n", firstNames)

// 			if remainingTickets == 0 {
// 				fmt.Println("Our conference ticket are sold out. Come back next year.")
// 				break //end program
// 			}
// 		} else {
// 			//fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n", remainingTickets, userTickets)
// 			//continue
// 			if !isValidName {
// 				fmt.Println("First name or last name you entered is too short.")
// 			}
// 			if !isValidEmail {
// 				fmt.Println("Email address you entered doesn't have @ sign.")
// 			}
// 			if !isValidTicketNumber {
// 				fmt.Println("Number of tickets you entered is invalid.")
// 			}
// 		}

// 	}
// 	/* switch statements example:

// 	city := "London"

// 	switch city{
// 	case "New York":
// 		//execute code for booking New York conference tickets
// 	case "Singapore":
// 		//execute code for booking Singapore conference tickets
// 	case "Berlin", "Delhi":
// 		//execute code for booking Berlin and Delhi conference tickets
// 	default:
// 		fmt.Println("No valid city selected")
// 	}*/

// }

// func greetUsers() {
// 	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
// 	//fmt.Printf("conferenceTicket is %T, remainingTickets is %T, and conferenceName is %T data type.\n", conferenceTickets, remainingTickets, conferenceName)
// 	//fmt.Println("Welcome To", conferenceName, "Booking Application")
// 	//fmt.Printf("Welcome To %v Booking Application\n", conferenceName)
// 	//fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are sill available")
// 	fmt.Printf("We have total of %v tickets and %v are sill available.\n", conferenceTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

// func getFirstNames() []string {
// 	firstNames := []string{}
// 	for _, booking := range bookings { // Blank Identifer which variable you don't want to use
// 		//firstNames = append(firstNames, booking["firstName"]) // for the map
// 		firstNames = append(firstNames, booking.firstName) //for struct
// 		//var names = strings.Fields(booking) //Fields ->  Splits the string with white space as separator and returns a slice with the split elements
// 		//firstNames = append(firstNames, names[0])
// 	} //for arrays and slices, range provides the index and value for each element
// 	//fmt.Printf("The first name of the people who booked the tickets are: %v\n", firstNames)
// 	return firstNames
// }

// func getUserInput() (string, string, string, uint) {
// 	var firstName string
// 	var lastName string
// 	var email string
// 	var userTickets uint
// 	//ask user for their name
// 	fmt.Println("Enter you first name:")
// 	fmt.Scan(&firstName) //Passing memory address

// 	fmt.Println("Enter you last name:")
// 	fmt.Scan(&lastName)

// 	fmt.Println("Enter you email address:")
// 	fmt.Scan(&email)

// 	fmt.Println("Enter number of tickets:")
// 	fmt.Scan(&userTickets)

// 	return firstName, lastName, email, userTickets
// }

// func bookTicket(userTickets uint, firstName string, lastName string, email string) {
// 	remainingTickets = remainingTickets - userTickets

// 	//Create a map for user
// 	/*userData := make(map[string]string)
// 	userData["firstName"] = firstName
// 	userData["lastName"] = lastName
// 	userData["email"] = email
// 	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)*/

// 	//For struct
// 	var userData = UserData{
// 		firstName:       firstName,
// 		lastName:        lastName,
// 		email:           email,
// 		numberOfTickets: userTickets,
// 	}

// 	//bookings[0] = firstName + " " + lastName // for the array
// 	//bookings = append(bookings, firstName+" "+lastName) for the slices
// 	bookings = append(bookings, userData)

// 	//fmt.Printf("The whole array: %v\n", bookings)
// 	//fmt.Printf("The first value: %v\n", bookings[0])
// 	//fmt.Printf("Array type %T\n", bookings)
// 	//fmt.Printf("Array length : %v\n", len(bookings))

// 	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
// 	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)
// 	//return bookings["numberOfTickets","firstName","lastName", "email"]
// 	//return userData
// }

// func sendTicket(userTickets uint, firstName string, lastName string, email string) {
// 	time.Sleep(10 * time.Second)                                                       //Sleep stops the thread for 10 sec
// 	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName) //Sprintf to store the value in the variable ticket
// 	fmt.Println("#####################")
// 	fmt.Printf("Sending tickets:\n%v\nto email address %v\n", ticket, email)
// 	fmt.Println("#####################")
// }

// //wg add, wait and done

func main() {
	var x int
	fmt.Println(x)
}
