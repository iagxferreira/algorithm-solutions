package leetcode

import (
	"container/heap"
	"sort"
)

type eventsQueue []int

func (eq eventsQueue) isEmpty() bool {
	return eq.Len() == 0
}

func (eq eventsQueue) Len() int {
	return len(eq)
}

// Min-heap
func (eq eventsQueue) Less(i, j int) bool {
	return eq[i] < eq[j]
}

func (eq eventsQueue) Swap(i, j int) {
	eq[i], eq[j] = eq[j], eq[i]
}

func (eq *eventsQueue) Push(num interface{}) {
	*eq = append(*eq, num.(int))
}

func (eq *eventsQueue) Pop() interface{} {
	len := len(*eq)
	num := (*eq)[len-1]
	*eq = (*eq)[:len-1]
	return num
}

func (eq eventsQueue) Peek() interface{} {
	return eq[0]
}

func maxEvents(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	eventsQueue := new(eventsQueue)
	eventsIndex, out := 0, 0

	for currentDay := 1; currentDay <= 100000; currentDay++ {

		for !eventsQueue.isEmpty() && eventsQueue.Peek().(int) < currentDay {
			heap.Pop(eventsQueue)
		}

		for eventsIndex < len(events) && events[eventsIndex][0] == currentDay {
			heap.Push(eventsQueue, events[eventsIndex][1])
			eventsIndex++
		}

		if !eventsQueue.isEmpty() {
			heap.Pop(eventsQueue)
			out++
		}
	}

	return out
}
