package main

import (
	"fmt"
	"strings"
)

func main() {
	// out := removeDuplicates("abbaca")
	// fmt.Println(out)
	out := removeDuplicates("aaaaa")
	fmt.Println(out)
}

func removeDuplicates(s string) string {
	prev := 0
	i := 1
	delete := 0
	strs := strings.Split(s, "")
	for i < len(strs) {
		fmt.Println(prev, i, strs)
		if strs[prev] == strs[i] {
			// fmt.Println(strs)
			if i == len(strs) - 1 {
				fmt.Println("last")
				fmt.Println(prev, i, strs)
				strs = strs[:len(strs)-2]
				continue
			}
			strs = append(strs[:prev], strs[prev+1:]...)
			delete = 1
		} else {
			if delete == 1 {
				delete = 0
				strs = append(strs[:prev], strs[prev+1:]...)
				if prev == 0 {
					continue
				}
				i--
				prev --
				continue
			}
			i++
			prev++
			delete = 0
		}
	}
	return strings.Join(strs, "")
}

