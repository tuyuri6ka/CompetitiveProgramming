package main

import (
	"fmt"
	"strconv"
)

type Card struct {
	suit  string
	value int
}

func main() {
	// 標準入力を取得する
	var N int
	fmt.Scan(&N)
	input := inputSource(N)

	// 問題に適した形に整形する
	processingInput := inputProcessing(input)

	result := bubbleSort(processingInput)

	fmt.Println(result)

}

func inputSource(N int) []string {
	input := make([]string, N, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&input[i])
	}
	return input
}

func inputProcessing(input []string) []Card {
	result := make([]Card, len(input))
	for i, v := range input {
		result[i].suit = v[:1]
		result[i].value, _ = strconv.Atoi(v[1:])
	}
	return result
}

func bubbleSort(card []Card) []Card {
	result := card
	for i := 0; i < len(result); i++ {
		for j := len(result) - 1; j > i; j-- {
			if result[j-1].value > result[j].value {
				big := result[j-1].value
				small := result[j].value

				result[j-1].value = small
				result[j].value = big
			}
		}
	}

	return result
}
