

// In Go language, defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns.




package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
        defer fmt.Println("hii")
        fmt.Println("Last")
}

