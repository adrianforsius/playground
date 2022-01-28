package main

import (
	"container/heap"
	"log"
)

func main() {

	pq := make(PriorityQueue, 0)
	q := &pq
	heap.Init(q)
	heap.Push(q, &node{lowest: 0})
	heap.Push(q, &node{lowest: 1})
	heap.Push(q, &node{lowest: 5})
	heap.Push(q, &node{lowest: 8})
	heap.Push(q, &node{lowest: 3})

	log.Println("popped", heap.Pop(q).(*node))
	heap.Push(q, &node{lowest: 2})
	for _, i := range *q {
		log.Printf("item %+v", i)
	}

	log.Println("popped", heap.Pop(q).(*node))
	for _, i := range *q {
		log.Printf("item %+v", i)
	}
}

type node struct {
	pos     []int
	visited bool
	weight  int
	lowest  int
	parent  *node
	index   int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].lowest < pq[j].lowest
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
