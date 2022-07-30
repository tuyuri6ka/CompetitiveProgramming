package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 標準入力を取得する
	var N int
	fmt.Scan(&N)
	input := inputSource(N)

	// 問題に適した形に整形する
	result := inputProcessing(input)

	fmt.Println(result)
}

func inputSource(N int) []string {
	input := make([]string, N, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&input[i])
	}
	return input
}

func inputProcessing(input []string) []map[string]int {
	result := make([]map[string]int, len(input), cap(input))
	for i, _ := range input {
		hash := make(map[string]int)
		hash[input[i][:1]], _ = strconv.Atoi(input[i][1:])

		result[i] = hash
	}
	return result
}
