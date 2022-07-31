package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	input := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&input[i])
	}

	list := make([]int, len(input), cap(input))
	copy(list, input)

	result := selectionSort(list, N)

	for i, _ := range result {
		fmt.Print(result[i])
		if i != N {
			fmt.Print(" ")
		}
	}
}

func selectionSort(result []int, N int) []int {
	for i := 0; i < N; i++ {
		mini := i
		for j := i; j < N; j++ {
			if result[mini] > result[j] {
				mini = j
			}
		}
		small := result[mini]
		big := result[i]

		// i 以降の最小値と入れ替える
		result[i] = small
		result[mini] = big

		fmt.Println(result)
	}

	return result
}

func output(r []int, N int) {
	for i := 0; i < N; i++ {
		fmt.Printf("%d", r[i])
		if i != N {
			fmt.Printf(" ")
		}
	}
}
