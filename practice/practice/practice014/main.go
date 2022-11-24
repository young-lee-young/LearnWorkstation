package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
电话号码的字母组合 LeetCode No017

标签：回溯
*/
func main() {
	numStr := "23"
	numSlice := strings.Split(numStr, "")
	tempRet := make([]string, 0)
	result := solution(numSlice, tempRet)
	fmt.Println(result)
}

func solution(numSlice []string, tempRet []string) []string {
	numCharMap := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	if len(numSlice) == 0 {
		return tempRet
	}

	num, _ := strconv.Atoi(numSlice[0])
	strSlice := numCharMap[num]
	newRet := make([]string, 0)
	if len(tempRet) == 0 {
		for _, str := range strSlice {
			newRet = append(newRet, str)
		}
	} else {
		for _, str := range strSlice {
			for _, ret := range tempRet {
				ret += str
				newRet = append(newRet, ret)
			}
		}
	}

	newNumSlice := numSlice[1:]
	return solution(newNumSlice, newRet)
}
