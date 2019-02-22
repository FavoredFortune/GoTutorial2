package main

import "fmt"

//
//func c() (i int){
//	//defers an anonymous function and then calls it immediately
//	defer func(){ i++}()
//	return 1
//}

func main() {

	////defers an anonymous function and then calls it immediately
	//defer func(s string) {fmt.Println(s)}("Today is ...")
	//
	//fmt.Println(c())
	First()
	fmt.Println("Returned normally from First function.")
}

func First(){
	defer func(){
		if r := recover(); r!=nil{
			fmt.Println("recovered in First function", r)
		}
	}()
	fmt.Println("Calling Second function.")
	Second(0)
	fmt.Println("Returned normally from Second function.")
}

func Second(i int){
	if i > 3 {
		fmt.Println("PANICKING!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer happening in Second function", i)
	fmt.Println("Printing in Second function.", i)
	Second(i+1)
}