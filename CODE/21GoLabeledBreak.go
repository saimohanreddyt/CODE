/* The standard unlabeled break statement is used to terminates the nearest enclosing statement.
In Golang, there is another form of break(labeled break) statement is used to terminate specified "for", 
"switch", or "select" statement. In GoLang, Label is an identifier which is followed by a colon (:)sign*/


package main
import "fmt"

func main() {

	var count = 0

	fmt.Println("Go labeled break statement.")

	outer:
	for count <= 10 {
		count++
		if (count == 5) {
            break outer
        }
        fmt.Printf("Inside loop %d\n", count)
	}
	fmt.Println("Out of loop")

}
