
package main

import (
       "fmt"
       )
func main() {

      var students [3] string
       fmt.Printf("Students:  %v\n", students)
     students[0] = "Lisa"
     students[2] = "john"
     students[1] = "karter"
       fmt.Printf("Students #1: %v\n", students[1])
       fmt.Printf("Students: %v\n", students)
       fmt.Printf("Number of Students: %v\n",len(students))
}






/*

package main

import (
       "fmt"
       )
func main() {

      grades := [3]int{97, 98, 99}
      fmt.Printf("Grades: %v",grades)

}
*/
