package main

import "fmt"

type Queue struct {
	name string
	time int
}

var head int
var tail int

const LEN = 10000

/*
input
5 100
p1 150
p2 80
p3 200
p4 350
p5 20

output
p2 180
p5 400
p1 450
p3 550
p4 800
*/
func main() {
	var N int
	fmt.Scan(&N)
	var q int
	fmt.Scan(&q)
	queues := inputSources(N)

	roundRobin(queues, q)
}

func inputSources(N int) []Queue {
	input := make([]Queue, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&input[i].name)
		fmt.Scan(&input[i].time)
	}
	return input
}

func roundRobin(queue []Queue, q int) {
	head = 0
	tail = len(queue)
	elaps := 0
	var element Queue
	for head != tail {
		element = dequeue(queue)

		min := min(q, element.time)
		element.time -= min
		elaps += min

		if element.time > 0 {
			queue = enqueue(queue, element)
		} else {
			// Answer用の出力
			fmt.Printf("%s %d\n", element.name, elaps)
		}
	}
}

func enqueue(queue []Queue, element Queue) []Queue {
	// appendは内部で新しくメモリを取り直す
	result := append(queue, element)
	tail = (tail + 1) % LEN

	return result
}

func dequeue(queue []Queue) Queue {
	element := queue[head]
	queue = queue[1 : len(queue)-1]
	head = (head + 1) % LEN

	return element
}

func min(q int, element int) int {
	var result int

	if q < element {
		result = q
	} else {
		result = element
	}

	return result
}
