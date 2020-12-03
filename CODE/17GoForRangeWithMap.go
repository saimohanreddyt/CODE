/* In GoLang, range on map iterates over key/value pairs.
it can also iterate over just the key of a map.  */

package main
import "fmt"

func main() {

	fmt.Println("Go For Each with Map")
	fruits := map[string]string{"A": "Apple", "B": "banana", "C": "Cherry"}
    for key, value := range fruits {
        fmt.Printf("%s -> %s\n", key, value)
    }

    for key := range fruits {
        fmt.Println("key:", key)
    }
}
