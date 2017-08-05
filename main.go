package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./sort"
)

func main() {
	sort.InsertionSort()
	fmt.Println(len(readfile()))
}

func readfile() []int {
	var list []int
	file, _ := os.Open("./data/input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		list = append(list, i)
	}
	return list
}
