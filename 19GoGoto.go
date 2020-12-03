/* In GoLang, goto statement is udes to alter the normal exection of a program and transfer control
to a labeled statement in the same program. The label is an identifier, it can be any plain text 
and can be any plain text and can be set anywhere in the Go program above or below statement is encountered,
compiler transfers the control to a label: and begin execution from there. */

/*
Syntax:-
--------
goto label
..
.
label:
*/


package main
import "fmt"
func main() {
   var age int
election:
   fmt.Print("Enter your age ")
   fmt.Scanln(&age)
   if (age <= 17) {
      fmt.Print("You are not eligible to vote.\n")
      goto election
   } else {
      fmt.Print("You are eligible to vote.\n")
   }
}
