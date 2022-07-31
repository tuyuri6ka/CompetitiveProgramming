package main

import "testing"

func TestDequeue(t *testing.T) {
	queue := []Queue{Queue{name: "p1", time: 150}, Queue{name: "p2", time: 100}}
	element := dequeue(queue)

	if element.name != "p1" {
		t.Error(`miss`)
	}
	if element.time != 150 {
		t.Error(`miss`)
	}
}

func TestEnqueue(t *testing.T) {
	queue := []Queue{Queue{name: "p1", time: 150}, Queue{name: "p2", time: 100}}
	element := Queue{name: "p3", time: 200}

	result := enqueue(queue, element)

	if result[0].time != 150 {
		t.Error(`miss`)
	}
	if result[2].time != 200 {
		t.Error(`miss`)
	}
}
