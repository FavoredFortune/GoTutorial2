package main

func main() {
	
}

type Calculator interface {
	Add (a, b int64) int64
	Subtract (a,b int64) int64
	Divide (a,b int64) int64
	Multiply (a,b int64) int64
}
