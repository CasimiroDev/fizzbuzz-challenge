package main

import (
	"fmt"
	"strconv"
)

type combinationType struct {
	multiple int
	result    string
}

func main() {
	combinations := []combinationType{
		{multiple: 3, result: "Fizz"},
		{multiple: 5, result: "Buzz"},
	}
	for i := 1; i <= 100; i++ {
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
