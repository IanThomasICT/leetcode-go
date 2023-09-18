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
	type Node struct {
		num int;
		freq int;
		next *Node;
	}

	// Get all occurrences
	allNums := make(map[int]int);
	for i := range nums {
		allNums[nums[i]]++
	}

	
	// Max size of k
	head := &Node{num: 0, freq: 0}
	for num, freq := range allNums {
		if (freq > head.freq) {
			tmp := &Node{num: num, freq: freq, next: head}
			head = tmp
			continue;
		} 			
			
		tmp, i := head, 1
		for tmp.next != nil && tmp.next.freq > freq {
			if (i >= k) {
				break;
			}
			tmp = tmp.next
			i++;
		}
		if (i >= k) {
			continue;
		}

		tmp.next = &Node{num: num, freq: freq, next: tmp.next}
	}

	res := make([]int, k);
	for i := 0; i < k; i++ {
		res[i] = head.num;
		head = head.next;
	}

	return res;
}