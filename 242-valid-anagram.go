package main

import "fmt"

func main() {
	s,t := "anagram", "nagaram"
	fmt.Println(isAnagram(s,t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := [26]int{0};
	
	for i := range s {
		chars[s[i]-97]++
		chars[t[i]-97]--
	}

	for _, a := range chars {
        if a != 0 {
            return false
        }
    }
    return true
}