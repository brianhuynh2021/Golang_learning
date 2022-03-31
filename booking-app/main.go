package main

import "fmt"
import "strings"

func main() {
    //var conferenceName string = "Go Conference"
    conferenceName := "Go Conference"
    const conferenceTickets int = 50
    var remainingTickets uint = 50
    bookings := []string{}
    //Print type of a variable
    fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
    fmt.Printf("Welcome to %v booking application\n", conferenceName)
    fmt.Printf("We have total of %v tickets and %v remaining tickets are still available.\n", conferenceTickets, remainingTickets)
    fmt.Printf("Get your tickets here to attend\n")
    
    //var bookings = [50]string{"Brian", "Nhu", "Trang", "Son"}
    //var bookings[] string
    //var bookings = []string{}
    
    for {
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
        fmt.Println("How many tickets you have:")
        fmt.Scan(&userTickets)
        remainingTickets = remainingTickets - userTickets
        bookings = append(bookings, firstName + " " + lastName)
        fmt.Printf("The whole array: %v\n", bookings)
        fmt.Printf("The first value: %v\n", bookings[0])
        //fmt.Printf("Array type: %T\n", bookings)
        fmt.Printf("Slice type: %T\n", bookings)
        //fmt.Printf("Array length: %v\n", len(bookings))
        fmt.Printf("Slice length: %v\n", len(bookings))

        fmt.Print()
        fmt.Printf("User %v %v booked %v tickets, and you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
        fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
        
        firstNames := []string{}
        for _, booking := range bookings {
            var names = strings.Fields(booking)
            firstNames = append(firstNames, names[1])
        }
        fmt.Printf("The first name of bookings are: %v\n", firstNames)
    }
}
