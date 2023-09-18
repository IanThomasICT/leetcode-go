package main

import "fmt"

func main() {
	input, k := []int{1,1,1,2,2,3}, 2
	res := topKFrequent(input, k);
	fmt.Println(res)
}


// max
// ...
// ...
// ...
// min

func topKFrequent(nums []int, k int) []int {
	// Get all occurrences
	allNums := make(map[int]int);
	for i := range nums {
		allNums[nums[i]]++
	}

	buckets := make([][]int, len(nums)+1);
	for num, freq := range allNums {
		buckets[freq] = append(buckets[freq], num);
	}

	res := make([]int, 0);
	for i := len(buckets)-1; i > 0; i-- {
		if (len(res) >= k) {
			break;
		}

		res = append(res, buckets[i]...);
	}

	return res;
}