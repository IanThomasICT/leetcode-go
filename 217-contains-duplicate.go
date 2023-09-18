package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

var numOfInputs = 10

func main() {
	rand.NewSource(time.Now().UnixNano());
	
	iter := 2
	var elapsedTime time.Duration
	for i := 0; i < iter; i++ {
		fmt.Printf("Round %d\n", i + 1)
		input := rand.Perm(rand.Intn(10000));
		fmt.Println(input[0: min(25, len(input))])
		start := time.Now()
		containsDuplicate(input)
		elapsedTime += time.Since(start);
	}


	fmt.Printf("Overall Average: %f", float64(elapsedTime.Milliseconds()) / float64(iter))
}


func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool);
	return slices.ContainsFunc(nums, func(n int) bool {
		if _, ok := numMap[n]; ok {
			return true
		}
		numMap[n] = true
		return false;
	});
}

func min(a, b int) int {
    if a <= b {
        return a
    }
    return b
}