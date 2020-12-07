package main

import "fmt"

func main() {

      var i float32 = 7.7
      fmt.Printf("%v, %T\n", i, i)
      var j int
      j = int(i)
      fmt.Printf("%v, %T\n", j, j)
}
