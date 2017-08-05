package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"./sort"
)

func main() {
	list := readfile()
	start := time.Now()
	sort.InsertionSort(list)
	end := time.Now()
	writefile(list)
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
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

func writefile(list []int) {
	file, err := os.Create("./data/sorted.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	for _, e := range list {
		file.Write(([]byte)(fmt.Sprintf("%08s\n", strconv.Itoa(e))))
	}
}
