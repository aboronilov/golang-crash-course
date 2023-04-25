package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	booking := []string{}

	fmt.Printf("conferenceTickets type is %T, remainingTickets type is %T, conferenceName type is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 && len(booking) < 50 {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("How many tickets would you like to buy?")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			booking = append(booking, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets left for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, value := range booking {
				var names = strings.Fields(value)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The client names: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("All the tickets are sold")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your name is invalid")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid")
			}
		}
	}
}
