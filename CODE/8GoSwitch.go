package main
import "fmt"
   func main() {

       var WeekDay = 3

        switch WeekDay {
        case 1, 2, 3, 4, 5:
          fmt.Println("It's a Weekday.")
        case 6, 7:
          fmt.Println("It's Weekend.")
        default:
          fmt.Println("Invalid Day")
        }
}
