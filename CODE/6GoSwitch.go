package main
import "fmt"
func main() {

	var dayOfWeek = 5
	fmt.Println("Go switch statement.")
	switch dayOfWeek {
	case 1:
		fmt.Println("Today is Monday.")
	case 2:
		fmt.Println("Today is Tuesday.")
	case 3:
		fmt.Println("Today is Wednesday.")
	case 4:
		fmt.Println("Today is Thursday.")
	case 5:
		fmt.Println("Today is Friday.")
	case 6:
		fmt.Println("Today is Saturday.")
	case 7:
		fmt.Println("Today is Sunday.")
	default:
		fmt.Println("Invalid Weekday.")
	}

}
