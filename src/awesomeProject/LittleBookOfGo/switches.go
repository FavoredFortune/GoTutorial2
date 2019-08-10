package LittleBookOfGo

import (
	"fmt"
	"runtime"
	"time"
)

//for loops
func Sqrt(x float64) float64 {

	z := 1.00

	for z <= x/2 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}


func main() {
	fmt.Println(Sqrt(4))

	//switch cases
	fmt.Print("My Go runs on ")
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s.", os)
	}

	//evaluation order of switch cases
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today!")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("Two days from now, sadly.")
	default:
		fmt.Println("Too far away.")
	}

	//switch cases without conditions work as if 'switch when true'
	t := time.Now()
	switch {
	case t.Hour() <12:
		fmt.Println("Morning.")
	case t.Hour() <17:
		fmt.Println("Afternoon, peoples!")
	default:
		fmt.Println("Evenin' y'all.")
	}


}

