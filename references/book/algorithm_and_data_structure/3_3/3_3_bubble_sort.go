package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	input := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&input[i])
	}

	result := bubbleSort(N, input)

	for i, _ := range result {
		fmt.Print(result[i])
		if i != N {
			fmt.Print(" ")
		}
	}
}

func bubbleSort(N int, input []int) []int {
	list := make([]int, len(input), cap(input))
	copy(list, input)

	// 一回ループが回れば1つソート済みとなるので、ソート済みの行をiで管理すると計算数が減る
	flag := 1
	for i := 0; flag == 1; i++ {
		flag = 0
		for j := N - 1; j >= i+1; j-- {
			if list[j] < list[j-1] {
				small := list[j]
				big := list[j-1]

				list[j] = big
				list[j-1] = small

				flag = 1
			}
		}

		fmt.Println(list)
	}

	return list
}
