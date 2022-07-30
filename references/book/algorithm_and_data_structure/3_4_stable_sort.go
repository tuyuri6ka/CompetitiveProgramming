package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	input := make([]string, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&input[i])
	}

	fmt.Println(input)
	// ハッシュのスライス
	hash := make([]map[string]int, N, N)
	for i, _ := range input {
		tmp := make(map[string]int)
		tmp[input[i][:1]], _ = strconv.Atoi(input[i][1:])
		hash[i] = tmp
	}
	fmt.Println(hash)

}
