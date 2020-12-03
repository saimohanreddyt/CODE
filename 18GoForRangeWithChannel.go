/* For channel, it iterates over values sent on the channel until closed. */

package main
import "fmt"

func main() {

	fmt.Println("Go For Each with Channel")
	ch := make(chan string)
	go func() {
		ch <- "G"
		ch <- "O"
		ch <- "L"
		ch <- "A"
		ch <- "N"
		ch <- "G"
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}
}
