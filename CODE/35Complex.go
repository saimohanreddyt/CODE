
package main

import (
       "fmt"
       )
func main() {

       var n complex64 = 1 + 2i
       fmt.Printf("%v, %T\n", real(n), real(n))
       fmt.Printf("%v, %T\n", imag(n), imag(n))
}



/*

package main

import (
       "fmt"
       )
func main() {

       var n complex64 = complex(5, 12)
       fmt.Printf("%v, %T\n", n, n)
}

*/
