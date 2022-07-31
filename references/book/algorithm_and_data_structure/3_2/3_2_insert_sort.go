package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	list := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&list[i])
	}

	result := insetSort(list, N)

	fmt.Print("result: ", result)
}

func insetSort(list []int, N int) int {
	count := 0

	// 挿入ソート
	// イメージは比較した値をごっそり横にスライドさせるイメージとなる。
	for i := 0; i < N; i++ {
		// スライスは入れ替えるので、移動させたい値を保存しておく。
		value := list[i]

		// 比較対象をずらしていくため、ｉとは別に取得しておく。
		j := i - 1
		// 比較して大きかったら隣にスライドさせていく。
		for j >= 0 && list[j] > value {
			list[j+1] = list[j]
			//比較対象をずらす。
			j--
		}
		list[j+1] = value
		count++
		output(list, N)
	}

	return count
}

func output(r []int, num int) {
	for i := 0; i < num; i++ {
		fmt.Printf("%d", r[i])
		if i != num {
			fmt.Printf(" ")
		}
	}
	fmt.Print("\n")
}
