 /* In Golang, if we remove the "initialization" and "increment" statement 
then the Go for loop behaves like a while loop as in other programming languages.*/


package main
import "fmt"

func main() {
	power := 1
	fmt.Println("Go For loop as while loop.")
	for power < 5 {
		power *= 2
	}
	fmt.Println(power)
}
