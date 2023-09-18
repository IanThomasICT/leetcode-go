package main

import "fmt"

func main() {
	input := []int{1,2,3,4}
	res := productExceptSelf(input);
	fmt.Println(res)
}

func productExceptSelf(nums []int) []int {
    startingInts := make(map[int][]int);
    totals := make(map[int]int);

	for i := range nums {
		startingInts[nums[i]] = append(startingInts[nums[i]], i);
	}

	res := make([]int, len(nums));
	for key, _ := range startingInts {
		totals[key] = 1;
		for num, _ := range totals {
			if (num != key) {
				totals[key] *= num;
			}
		}
	}

	return res;
}