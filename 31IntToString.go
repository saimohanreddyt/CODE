package main

import  (
        "fmt"
       "strconv"

)  //We need to use String Conversion Package 
func main() {

      var i int = 7
      fmt.Printf("%v, %T\n", i, i)
      var j string
      j = strconv.Itoa(i)
      fmt.Printf("%v, %T\n", j, j)
}
