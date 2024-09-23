package queue

import (
	"fmt"
	"time"
)

var ErrorPriorityQueueOnCapacity = fmt.Errorf("la cola está llena")

type QueueItem interface {
	Priority() int
	Arrival() time.Time
}

type PriorityQueue struct {
	Items    []QueueItem `json:"items"`
	Capacity uint        `json:"capacity"`
}

func (pq PriorityQueue) Len() int { return len(pq.Items) }

func (pq PriorityQueue) Less(i, j int) bool {
	a, b := pq.Items[i], pq.Items[j]
	if a.Priority() == b.Priority() {
		// Si las prioridades son iguales, se da prioridad al que llegó primero
		return a.Arrival().Before(b.Arrival())
	}
	return a.Priority() < b.Priority() // Ordenamiento por prioridad
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(QueueItem)
	pq.Items = append(pq.Items, item)
}

func (pq *PriorityQueue) Pop() any {
	if len(pq.Items) == 0 {
		return nil
	}
	n := len(pq.Items)
	item := pq.Items[n-1]
	pq.Items = pq.Items[0 : n-1]
	return &item
}

func (pq PriorityQueue) CanPush() error {
	if len(pq.Items) >= int(pq.Capacity) {
		return ErrorPriorityQueueOnCapacity
	}
	return nil
}
