package knapsack

import (
	"bufio"
	"fmt"
	"os"
	"path"
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

func (i *Item) String() string {
	return fmt.Sprintf("%-10s $%10,2f $%10,2f g", i.Name, i.Worth, i.Weight)
}

func readFile(csv string) (*[]Item, error) {
	open, err := os.Open(csv)
	if err != nil {
		return nil, fmt.Errorf("unable to open file")
	}
	defer open.Close()

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

func Knapsack() ([]Item, error) {

	if len(os.Args) != 2 {
		fmt.Printf("Provide Knapsack CSV file with file path: %s <input_file>\n",
			path.Base(os.Args[0]))
		os.Exit(0)
	}

	store, err := readFile(os.Args[1])

	if err != nil {
		return nil, err
	}

	fmt.Println("list of goods to choose from: ")
	for _, v := range *store {
		fmt.Println(" ", v)
	}

	if len(os.Args) != 1 {
		fmt.Printf("Provide empty knapsack weight(g): %s")
		os.Exit(0)
	}
	answer := bufio.NewReader(os.Stdin)

	line, err := answer.ReadString('\n')
	if err != nil {
		return nil, err
	}

	knapsackWeight, err := strconv.Atoi(line)
	if err != nil {
		return nil, err
	}

	if len(os.Args) != 1 {
		fmt.Printf("Provide max capacity knapsack can carry(g): %s")
		os.Exit(0)
	}

	response := bufio.NewReader(os.Stdin)

	cap, err := response.ReadString('\n')
	if err != nil {
		return nil, err
	}

	knapsackCapacity, err := strconv.Atoi(cap)
	if err != nil {
		return nil, err
	}

	mostItems, err := maxItemsOnly(knapsackWeight, knapsackCapacity, *store)

	if err != nil {
		return nil, err
	}

	return mostItems, nil
}

func maxItemsOnly(knapsackWeight int, knapsackCapacity int, store []Item) ([]Item, error) {
	var happySack []Item
	totalWeight := knapsackCapacity - knapsackWeight

	// sort by worth first with highest at front
	sort.Slice(store, func(i, j int) bool {
		return store[i].Worth > store[j].Worth
	})

	//sort by weight with lightest at front
	for _, i := range store {
		if totalWeight+i.Weight <= knapsackCapacity {
			happySack = append(happySack, i)
			totalWeight += i.Weight
		}

	}
	return happySack, nil
}
