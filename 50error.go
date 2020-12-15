//In other languages we have trt-catch, try-except to handle the error
//In go lang we have error . Here we return the error to function

package main
import (
	"fmt"
	"errors"
	)
func main() {
	total, err := sumOf(10, 8)
	if err != nil {
		fmt.Println("There was an error", err)
	}else {
		fmt.Println(total)
	}
}
func sumOf(start int, end int)(int, error) {
	if start > end {
		return 0, errors.New("Start is greather than end")
	}
	total := 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total, nil
}
