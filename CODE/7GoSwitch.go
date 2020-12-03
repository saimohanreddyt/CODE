package main
import "fmt"
  func main(){

   var WhichMonth=5;

     switch WhichMonth{

       case 1:
          fmt.Println("JAN")
       case 2:
          fmt.Println("FEB")
       case 3:
          fmt.Println("MAR")
       case 4:
          fmt.Println("APR")
       case 5:
          fmt.Println("MAY")
       case 6:
          fmt.Println("JUN")
       case 7:
          fmt.Println("JUL")
       case 8:
          fmt.Println("AUG")
       case 9:
          fmt.Println("SEP")
       case 10:
          fmt.Println("OCT")
       case 11:
          fmt.Println("NOV")
       case 12:
          fmt.Println("DEC")
       default:
          fmt.Println("Invalid Entry")
       }
}
