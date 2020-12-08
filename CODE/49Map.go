

package main

import "fmt"

func main() {
	grades := make(map[string]float32)

	grades["Jhonny"] = 42
	grades["Jessy"] = 38
	grades["Jimmy"] = 18
	fmt.Println(grades)

	JhonnyGrades := grades["Jhonny"]
	fmt.Println(JhonnyGrades)

	delete(grades, "Jhonny")
	fmt.Println(grades)

	for k, v := range grades{
		fmt.Println(k,":",v)
	}

}
