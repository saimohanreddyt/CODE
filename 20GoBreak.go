/* In Go, break statement gives you way to break or terminate the execution of innermost "for",
"switch", or "select" statement that containing ir, 
and transfers the execution to the next statement following it */


package main
import "fmt"

func main() {

	var count = 0
	fmt.Println("Go break statement.")
	for count <= 10 {
		count++
		if (count == 5) {
            break
        }
        fmt.Printf("Inside loop %d\n", count)
	}
	fmt.Println("Out of loop")

}
