package main
import "fmt"
func main() {

	fmt.Println("Go switch with short statement.")
	switch dayOfWeek := 5; dayOfWeek {
		case 1, 2, 3, 4, 5:
			fmt.Println("It's a Weekday.")
		case 6, 7:
			fmt.Println("Its Weekend.")
		default:
			fmt.Println("Invalid Day")		
	}

}
