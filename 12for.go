package main
import "fmt"

   func main() {
   sum := 0
    fmt.Println("Go for loop: for initialization; condition; increment/decrement")

     for i:=1; i<5; i++ {
     sum += i
     }
     fmt.Println(sum)
}
