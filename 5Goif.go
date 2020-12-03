package main
import "fmt"
func main() {

fmt.Println("Go If Statement with short statement.")

if x := 10; x%2 == 0 {
    fmt.Printf("%d is even\n", x)
} else {
    fmt.Printf("%d is odd\n", x)
}


}
