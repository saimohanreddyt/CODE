
  /* block of variable can be decleare together */

     /*
       var ( key type = "val"
             key type = "val"
             key type = "val"    )
                               */


package main
import "fmt"

 var (
      Name string = "Rajini Kanth"
      Profession string = "Actor"
      DoB      string = "12 December 1950"
      RemunerationCR int = 63
     )

  func main() {
     fmt.Printf("%v, %T\n", Name, Name)
     fmt.Printf("%v, %T\n", Profession, Profession)
     fmt.Printf("%v, %T\n", DoB, DoB)
     fmt.Printf("%v, %T\n", RemunerationCR, RemunerationCR)
}

