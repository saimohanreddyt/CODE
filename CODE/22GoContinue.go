/* The continue statement gives you way to skip over the current iteration of any loop.
When a continue statement is encountered in the loop, the Go interpreter ignores rest of
statements in the loop body for current iteration and returns the program execution to the very
first statement in the loop body. It does not terminates the loop rather continues with the next iteration */

/*

Syntax:-
--------
for <condition1> {
    // loop body
    if (condition2) {
        continue
    }
    // loop body
}

*/

package main
import "fmt"

func main() {

	var ctr = 0

	fmt.Println("Go continue statement.")

	for ctr < 10 {
		ctr++
		if (ctr == 5) {
            println("5 is skipped")
            continue
            println("This won't be printed too.")
        }
        fmt.Printf("Number is %d\n", ctr)

	}
	fmt.Println("Out of loop")

}
