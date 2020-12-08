

package main

import (
	"fmt"
)

func main()	{		//calling function
	var res int
	res = add(10, 30)
	fmt.Printf("Sum value:%d\n",res)
}

// called function
func add(a int, b int)int {	// The function return type is 'int' so give the 'int' followed by func
	var c int
	c = a+b
	return c
}
