package pqueue

import (
	"a-star/internal/model"

	pq "github.com/jupp0r/go-priority-queue"
)

// PriorityQueue wraps pq from github.com/jupp0r/go-priority-queue
// Since jupp0r uses hashes we cannot use it directly on slice-based type `node`
// Thus PrioritQueue converts nodes into string
type PriorityQueue struct {
	wrapee pq.PriorityQueue
}

func New() PriorityQueue {
	wr := pq.New()
	return PriorityQueue{wrapee: wr}
}

func (p *PriorityQueue) Pop() model.Node {
	nodeInterface, err := p.wrapee.Pop()
	if err != nil {
		panic("Error while popping node")
	}
	nodeStr, ok := nodeInterface.(string)
	if !ok {
		panic("Failed to assert type Node")
	}
	node := model.ParseNode(nodeStr)
	return node
}

func (p *PriorityQueue) Insert(node model.Node, priority float64) {
	p.wrapee.Insert(node.String(), -priority)
}
