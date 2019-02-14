package main

import (
	"fmt"
	"runtime"
)

func Sqrt(x float64) float64 {

	z := 1.00

	for z <= x/2 {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Print("My Go runs on ")
	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s.", os)
	}
}

