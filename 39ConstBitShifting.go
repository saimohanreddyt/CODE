

package main

import (
       "fmt"
       )
const(
    _ = iota // ignore first value by assigning to blank identifier
    KB = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
func main() {

      fileSize := 4000000000.
      fmt.Printf("%.2fGB\n", fileSize/GB)
}




/*


package main

import (
       "fmt"
       )

const(
     isAdmin = 1 << iota
     isHeadquarters
     canSeeFinancials

     canSeeAfrica
     canSeeAsia
     cabSeeNorthAmerica
     canSeeSouthAmerica

)
func main() {

      var roles byte = isAdmin | canSeeFinancials | canSeeEurope
      fmt.Printf("%b\n", roles)
      fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin)
      fmt.Printf("Is HQ? %v", isHeadquaeters & roles == isHeadquarters)
}


*/
