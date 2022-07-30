package main

import (
	"fmt"
	"strconv"
)

const N = 10000

var top int
var stack [N]int

// stackは機能として,
// 1.一番上の要素を取り出すpop
// 2.一番上に要素を追加するpush
// 3.スタックが空かを調べるisEmpty
// 4.スタックが満杯かを調べるisFull
// が基本的に存在する。
// また、topの位置がどこかを管理するインデックスが存在する。

// 今回は逆ポーランド記法で与えられた数式の計算結果を出力する。
func main() {
	top = 0
	i := 0
	for i < 1 {
		var scan string
		fmt.Scan(&scan)
		if scan == "+" {
			a := pop(stack)
			b := pop(stack)
			push(a + b)
		} else if scan == "-" {
			b := pop(stack)
			a := pop(stack)
			push(a - b)
		} else if scan == "*" {
			a := pop(stack)
			b := pop(stack)
			push(a * b)
		} else if scan == "" {
			i++
		} else {
			input, _ := strconv.Atoi(scan)
			push(input)
		}
	}
	fmt.Println(stack[1])
}

func pop(stack [N]int) int {
	a := stack[top]
	top--

	return a
}

func push(element int) [N]int {
	top++
	stack[top] = element

	return stack
}
