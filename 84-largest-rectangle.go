package main

import (
	"fmt"
)

func main() {
	input := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(input))
	input = []int{2, 1, 2}
	fmt.Println(largestRectangleArea(input))
	input = []int{1, 1}
	fmt.Println(largestRectangleArea(input))
	input = []int{0, 2, 0}
	fmt.Println(largestRectangleArea(input))
	input = []int{1, 2, 2}
	fmt.Println(largestRectangleArea(input))
	input = []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	fmt.Println(largestRectangleArea(input))
	input = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(largestRectangleArea(input))
}

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	// iterate over array
	stk := Stack{}
	prev, count, largestRect := 0, 0, 0
	for i, val := range heights[:len(heights)-1] {
		next := heights[i+1]
		if val != 0 {
			count++
		}
		if val == next {
			continue
		}

		stk.push([2]int{val, count})
		count = 0
		prev = val

		if val > next {
			for !stk.isEmpty() {
				if stk.peek()[0] < next {
					largestRect = max(largestRect, next*count)
					break
				}

				layer := stk.pop()
				prev = layer[0]
				count += layer[1]
				//fmt.Println("popped layer:", prev, ",count:", count, "->", prev*count)
				largestRect = max(largestRect, prev*count)
			}
		}
		prev = val
	}

	lastVal := heights[len(heights)-1]
	if prev == lastVal {
		count++
	}
	stk.push([2]int{lastVal, count + 1})
	count = 0
	fmt.Println(prev, count, stk.data)

	largestRect = max(largestRect, prev*count)
	for !stk.isEmpty() {
		layer := stk.pop()
		prev = layer[0]
		count += layer[1]
		//fmt.Println("popped layer:", prev, ",count:", count, "->", prev*count)
		largestRect = max(largestRect, prev*count)
	}

	return largestRect
}

type Stack struct {
	data [][2]int
}

func (s *Stack) push(i [2]int) {
	//fmt.Println("pushing", i)
	s.data = append(s.data, i)
}

func (s *Stack) pop() [2]int {
	i := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return i
}

func (s *Stack) peek() [2]int {
	return s.data[len(s.data)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
