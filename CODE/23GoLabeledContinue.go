/* The standard unlabeled continue statement is used to skip the current iteration of nearest
enclosing loop. In GoLang, there is another form of continue (labeled continue) statement is used to skip iteration of specified loop (can be outer loop).
 In Golang, Label is an identifier which is an identifier which is followed by colon(:) sign*/

package main
import "fmt"

func main() {

	var i = 0
	var j = 0

	fmt.Println("Go continue statement.")

	    outerfor:
	    for i < 3 {
	    	j = 0
	    	i++
	        for j < 3 {
	        	j++
	            fmt.Printf("i is %d and j is %d\n",i,j)
	            if(i == 2){
	                fmt.Println("I am Skipped")
	                continue outerfor
	            }
	            fmt.Println("I am Printed")
	        }
	    }


}
