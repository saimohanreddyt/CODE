/* For a String, the for loop iterates over each character's
Unicode code points. The first value is the starting byte 
index of the rune and the second the rune itself  */

package main
import "fmt"

func main() {

	fmt.Println("Go For Each with Strings")
	for i,s := range "GO LANG"{
		fmt.Printf("%U represents %c and it is at position %d\n", s, s, i)
	}
}
