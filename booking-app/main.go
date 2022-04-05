package main

import (
	"fmt"
	"strings"
    "booking-app/helper"
)

/*import (
  "fmt"
  "strings")*/

func main() {
    //var conferenceName string = "Go Conference"
    conferenceName := "Go Conference"
    const conferenceTickets int = 50
    var remainingTickets uint = 50
    bookings := []string{}
    helper.greetingUser()
    greetUsers(conferenceName)
    //Print type of a variable
    fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v remaining tickets are still available.\n", conferenceTickets, remainingTickets)
    fmt.Printf("Get your tickets here to attend\n")
    
    //var bookings = [50]string{"Brian", "Nhu", "Trang", "Son"}
    //var bookings[] string
    //var bookings = []string{}

    // for remainingTickets > 0 && len(bookings) < 50{
    //     ...
    // }
    for remainingTickets > 0 && len(bookings) < 50{
        var firstName string
        var lastName string
        var email string
        var userTickets uint
        // ask user for their name
        fmt.Println("Enter your first name:")
        fmt.Scan(&firstName)

        fmt.Println("Enter your last name:")
        fmt.Scan(&lastName)

        fmt.Println("Enter your email:")
        fmt.Scan(&email)
        if !isValidInput(email) {
            fmt.Println("Your email is invalid format, it must contain @")
        }
        
        fmt.Println("How many tickets you have:")
        fmt.Scan(&userTickets)
        
        
        yourAge()
        
        isValidName := len(firstName) >= 2 && len(lastName) >= 2
        isValidEmail := strings.Contains(email, "@")
        isValidTicketNumber := userTickets > 0
        //isValidCity := city == "Singapore" || city == "London"
        if isValidName && isValidEmail && isValidTicketNumber{
            if userTickets < remainingTickets {
                remainingTickets = remainingTickets - userTickets
                bookings = append(bookings, firstName + " " + lastName)
                fmt.Printf("The whole array: %v\n", bookings)
                fmt.Printf("The first value: %v\n", bookings[0])
                //fmt.Printf("Array type: %T\n", bookings)
                fmt.Printf("Slice type: %T\n", bookings)
                //fmt.Printf("Array length: %v\n", len(bookings))
                fmt.Printf("Slice length: %v\n", len(bookings))
    
                fmt.Printf("User %v %v booked %v tickets, and you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
                fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
                
            
                firstNames := []string{}
                for _, booking := range bookings {
                    var names = strings.Fields(booking)
                    firstNames = append(firstNames, names[1])
                }
                fmt.Printf("The first name of bookings are: %v\n", firstNames)
                //noTicketsRemaining := remainingTickets == 0
                // if remainingTickets == 0{
                //     // end program
                //     fmt.Println("Our conference is booked out. Come back next year.")
                //     break
                // }
            } else if userTickets == remainingTickets{
                fmt.Println("Our conference is booked out. Come back next year.")
                    break
            } else {
                fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
            }
        } else {
            if !isValidName {
                fmt.Println("First name or last name is you entered is invalid")
            }
            if !isValidEmail{
                fmt.Println("Email address you entered doesn't contain @ sign")
            }
            fmt.Println("The input information is invalid, try again!")
        }
        
    }

    city := "London"
    switch city {
        case "New York", "California":
            // execute code for booking New York conference tickets
        case "Berlin":
        case "Hongkong":
        case "Mexico":
        case "Singapore":
        default:
            fmt.Print("No valid city")
    }
}

func greetUsers(conferenceName string) {
    fmt.Printf("Welcome to our %v \n", conferenceName)
}

func isValidInput(email string) (bool) {
    isValidEmail := strings.Contains(email, "@")
    return isValidEmail
}
func yourAge() {
    var age uint
    fmt.Println("Your age is: ")
    fmt.Scan(&age)
    fmt.Printf("Your age is %v\n", &age)
}

