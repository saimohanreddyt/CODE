/* In Golang, if we further remove the "condition", 
then the Go for loop turns into an infinite loop */

package main
import "fmt"

func main() {
	/*power := 1*/
	fmt.Println("Go For loop as infinite loop.")
	sum := 0
	for {
		sum++ // repeated forever
	}
	fmt.Println(sum) // unreachable
}
