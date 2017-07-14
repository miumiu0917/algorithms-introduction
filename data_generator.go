package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"fmt"

	"github.com/cheggaaa/pb"
)

func main() {
	count := 100000
	set := make(map[int]struct{})
	bar := pb.StartNew(count)
	current := 0
	for len(set) < count {
		set[randomNumber()] = struct{}{}
		if current < len(set) {
			bar.Increment()
			current = len(set)
		}
	}

	file, err := os.Create("data/input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()
	for key := range set {
		file.Write(([]byte)(fmt.Sprintf("%08s\n", strconv.Itoa(key))))
	}
}

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000000)
}
