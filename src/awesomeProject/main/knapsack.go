package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Setup problem:
// The set up
// references for inspiration
// https://victoria.dev/verbose/knapsack-problem-algorithms-for-my-real-life-carry-on-knapsack/
// https://rosettacode.org/wiki/Knapsack_problem/0-1
// https://github.com/ivanpesin/golang-knapsack/blob/master/knapsack.go

//INPUT:
// CSV file with three columns:
// the item’s name (a string),
// a representation of its worth (an integer),
// and its weight in grams (an integer)
// There are 40 items in total.
//
// worth is done ranking each item from 40 to 1, 40 most important
//
//  Total weight of all items and bag: 9000g
//
//  Bag weight: 1400g
//
//  Airline limit: 7000g
//
//  Maximum weight of items I can pack: 5500g
//

//The challenge: Pack as many items as the weight limit allows while maximizing the total worth.

// Brute force solution

// 1. read in file by line assigning values and failing if part of data is missing or file is empty
// 2. sort data by highest worth then by lowest weight
// 3. create a data structure and start putting in items by highest worth and and lowest weight into map
// 4. keep checking total value of data structure for weight using an iterative loop.
// when weight exceeds limit, break out with only last prior value added to data structure.
// return data structure as knapsack list.

// Read file

type Item struct {
	Name   string
	Worth  int
	Weight int
}

//func (i *Item) String() string {
//	return fmt.Sprintf("%-10s $%10,2f $%10,2f g", i.Name, i.Worth, i.Weight)
//}

func readFile(csv string) (*[]Item, error) {
	open, err := os.Open(csv)
	if err != nil {
		return nil, fmt.Errorf("unable to open file")
	}
	defer  func(){
	err = open.Close()
	if err !=nil {
		 fmt.Printf("THIS IS BROKEN %s", err.Error())
	}}()

	var knapsack []Item

	read := bufio.NewScanner(open)
	for read.Scan() {
		line := strings.TrimSpace(read.Text())
		if len(line) > 0 && (line[0] == ';' || line[0] == '#') {
			continue
		}
		
		fields := strings.Fields(line)
		if len(fields) != 3 {
			return nil, fmt.Errorf("data file not in correct format")
		}
		
		worth, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, fmt.Errorf("unable to read file data")
		}
		
		weight, err := strconv.Atoi(fields[2])
		if err != nil {
			return nil, fmt.Errorf("unable to read file data")
		}
		
		knapsack = append(knapsack, Item{Name: fields[0], Worth: worth, Weight: weight})
	}

	if len(knapsack) < 1 {
		return nil, fmt.Errorf("file data appears to be incomplete")
	}
	
	return &knapsack, nil
}

func main() {
	
	//if len(os.Args) != 2 {
	//	fmt.Printf("Provide Knapsack CSV file with file path: %s <input_file>\n",
	//		path.Base(os.Args[0]))
	//	os.Exit(0)
	//}
	//
	//store, err := readFile(os.Args[1])
	//if err != nil {
	//	fmt.Printf("BAD JOB %s", err.Error())
	//}
	//
	////fmt.Println("list of goods to choose from: ")
	////for _, v := range *store {
	////	fmt.Println(" ", v)
	////}
	//
	//_, _ = fmt.Printf("file read")
	//
	//if len(os.Args) != 1 {
	//	fmt.Printf("Provide empty knapsack weight(g): %s")
	//	os.Exit(0)
	//}
	//answer := bufio.NewReader(os.Stdin)
	//
	//line, err := answer.ReadString('\n')
	//if err != nil {
	//	fmt.Printf(" SO STRANGE  %s", err.Error())
	//}
	//
	//knapsackWeight, err := strconv.Atoi(line)
	//if err != nil {
	//	fmt.Printf("WEIRDO %s", err.Error())
	//}
	//
	//if len(os.Args) != 1 {
	//	fmt.Printf("Provide max capacity knapsack can carry(g): %s")
	//	os.Exit(0)
	//}
	//
	//response := bufio.NewReader(os.Stdin)
	//
	//capacity, err := response.ReadString('\n')
	//if err != nil {
	//	fmt.Printf("%s", err.Error())
	//}
	//
	//knapsackCapacity, err := strconv.Atoi(capacity)
	//if err != nil {
	//	fmt.Printf("%s", err.Error())
	//}
	
	store :=[]Item{
		{Name: "kitten", Worth:40, Weight:240},
		{Name: "food", Worth:38, Weight:220},
		{Name: "water", Worth:39, Weight:120},
		{Name: "blanket", Worth:33, Weight:80},
		{Name: "toys", Worth:20, Weight:100},
		{Name: "book", Worth:24, Weight:90},
		{Name: "hat", Worth:28, Weight:30},
		{Name: "camera", Worth:26, Weight:110},
		{Name: "cat-food", Worth:37, Weight:160},
	}
	
	knapsackWeight := 370
	knapsackCapacity := 2440
	
	mostItems, err := MaxItemsOnly(knapsackWeight, knapsackCapacity, store)
	
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	
	fmt.Printf("here is the most items for your bag, prioritized by worth, %v. " +
		"remember your bag can only hold %v (g) including it's own weight %v (g)",
		mostItems, knapsackCapacity, knapsackWeight)
}

func MaxItemsOnly(knapsackWeight int, knapsackCapacity int, store []Item) ([]Item, error) {
	var happySack []Item
	totalWeight := knapsackCapacity - knapsackWeight
	fmt.Println(totalWeight)
	// sort by worth first with highest at front
	sort.Slice(store, func(i, j int) bool {
		return store[i].Worth > store[j].Worth
	})
	fmt.Println(store)
	//sort by weight with lightest at front
	for _, i := range store {
		if totalWeight- i.Weight > 0 {
			happySack = append(happySack, i)
			totalWeight -= i.Weight
			fmt.Println(totalWeight)
			fmt.Println(happySack)
		}

	}
	return happySack, nil
}
