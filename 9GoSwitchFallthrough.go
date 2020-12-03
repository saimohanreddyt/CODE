/* In Switch sttement a fallthrough statement allows to execute
     all the following statements after a match found*/

package main
import "fmt"

   func main() {

     var x = 2
        fmt.Println("Go Switch fallthrough")

        switch x{
        case 1:
	fmt.Println("1")
	fallthrough
	case 2:
	fmt.Println("2")
	fallthrough
	case 3:
	fmt.Println("3")
	}

}
