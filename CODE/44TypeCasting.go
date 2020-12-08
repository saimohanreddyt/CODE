


//      Type Casting
//	String & Float

package main

import (
	"fmt"
        "strconv"
)

func main() {
	price := "Price of Iphone is" + strconv.FormatFloat( 948.04,'f',2,64) + "dollar"
	fmt.Println(price)
}


//	Int and Float

package main

import (
        "fmt"
)

func main() {
	numberOFMobiles := 10
        price := 9480.02
	priceOFOneMobile := price/float64(numberOFMobiles)
        	fmt.Println(price)
		fmt.Printf("%T\n",price)
		fmt.Println(priceOFOneMobile)

}


