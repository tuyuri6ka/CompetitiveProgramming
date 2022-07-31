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

	result_by_bubble_sort := bubbleSort(processingInput)
	output_result(true, result_by_bubble_sort)

	result_by_selection_sort := selectionSort(processingInput)
	isStable := isStable(result_by_bubble_sort, result_by_selection_sort)
	output_result(isStable, result_by_selection_sort)
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

// bubbleSortは常に安定なソート（入力の順番を変えない）
func bubbleSort(card []Card) []Card {
	// 引数を直接変更しないようにディープコピーする
	result := make([]Card, len(card), cap(card))
	copy(result, card)

	for i := 0; i < len(result); i++ {
		for j := len(result) - 1; j > i; j-- {
			if result[j-1].value > result[j].value {
				big := result[j-1]
				small := result[j]

				result[j-1] = small
				result[j] = big
			}
		}
	}

	return result
}

// selectionSortは不安定なソート（入力の順番が変わりうる）
func selectionSort(card []Card) []Card {
	// 引数を直接変更しないようにディープコピーする
	result := make([]Card, len(card), cap(card))
	copy(result, card)

	for i := 0; i < len(result); i++ {
		minj := i
		for j := i; j < len(result); j++ {
			if result[minj].value > result[j].value {
				minj = j
			}
		}
		tmp := result[i]
		result[i] = result[minj]
		result[minj] = tmp
	}

	return result
}

func isStable(bubble_card []Card, selection_card []Card) bool {
	var flag bool
	for i, v := range bubble_card {
		if v.suit == selection_card[i].suit && v.value == selection_card[i].value {
			// 何もしない
		} else {
			flag = false
		}
	}

	return flag
}

func output_result(is_ok bool, result []Card) bool {
	if is_ok == true {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
	fmt.Println(result)

	return true
}
