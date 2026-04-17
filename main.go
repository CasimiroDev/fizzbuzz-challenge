package main

import (
	"fmt"
	"strconv"
)

type combinationType struct {
	multiple int
	result    string
}

var combinations = []combinationType{
	{multiple: 3, result: "Fizz"},
	{multiple: 5, result: "Buzz"},
	{multiple: 7, result: "Bazz"},
	{multiple: 11, result: "Jazz"},
}

func main() {
	for i := 1; i <= 300; i++ {
		result := ""
		for _, combo := range combinations {
			if i%combo.multiple == 0 {
				result += combo.result
			}
		}
		if result == "" {
			result = strconv.Itoa(i)
		}
		fmt.Println(result)
	}
}
