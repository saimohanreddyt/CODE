/*In GoLang, for loop with a "range" clause iterates through
  every element in an array, slice, string, map, or values received on a channel.
  A range keyword is used to iterate over slices, arrays, strings, maps or channela.*/

/*

  for key, value := range collection {
    //block of statements
}

*/

package main
import "fmt"

func main() {

	langs := []string{"Go", "C", "C++", "Java"}
	fmt.Println("Go For Each Loop.")
	for i,s := range langs{
		fmt.Println(i, s)
	}
}
/*Here, individual keys and values are held in the corresponding variable
and key or value can beomitted if one or the other is not needed using  ' ' */
