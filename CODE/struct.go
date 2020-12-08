

package main

import (
	"fmt"
)

type car struct{
	gas_pedal uint16 //min 0 max 65535
	break_pedal uint16
	steering_wheel int16 // -32k - +32k
	top_speed_kmh float64
}

func main() {
	a_car := car{gas_pedal: 2321,
	break_pedal:0, 
	steering_wheel:1234, 
	top_speed_kmh:230}

	fmt.Println(a_car.gas_pedal)
}

