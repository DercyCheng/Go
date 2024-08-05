package queue

import (
	"GTL"
	"GTL/internal/slice"
	"errors"
)

var (
	ErrOutOfCapacity = errors.New("GTL: 超出最大容量")
	ErrEmptyQueue    = errors.New("GTL:空队列")
)

// 小顶堆
// <=0 无界，自动扩容
// >0 有界，固定容量
type PriorityQueue[T any] struct {
	compare  GTL.Comparator[T]
	capacity int
	data     []T
}

func (p *PriorityQueue[T]) Len() int {
	return len(p.data) - 1
}

func (p *PriorityQueue[T]) Cap() int {
	return p.capacity
}

func (p *PriorityQueue[T]) IsBoundless() bool {
	return p.capacity <= 0
}

func (p *PriorityQueue[T]) IsFull() bool {
	return p.capacity > 0 && len(p.data)-1 == p.capacity
}

func (p *PriorityQueue[T]) IsEmpty() bool {
	return len(p.data) < 2
}
func (p *PriorityQueue[T]) Peek() (T, error) {
	if p.IsEmpty() {
		var t T
		return t, ErrEmptyQueue
	}
	return p.data[1], nil
}

func (p *PriorityQueue[T]) Enqueue(t T) error {
	if p.IsFull() {
		return ErrOutOfCapacity
	}
	p.data = append(p.data, t)
	node, parent := len(p.data)-1, (len(p.data)-1)/2
	for parent > 0 && p.compare(p.data[node], p.data[parent]) < 0 {
		p.data[parent], p.data[node] = p.data[node], p.data[parent]
		node = parent
		//TODO: bug?
		parent /= 2
	}
	return nil
}

func (p *PriorityQueue[T]) Dequeue() (T, error) {
	if p.IsEmpty() {
		var t T
		return t, ErrEmptyQueue
	}
	pop := p.data[1]
	p.data[1] = p.data[len(p.data)-1]
	p.data = p.data[:len(p.data)-1]
	p.shrinkIfNecessary()
	p.heapify(p.data, len(p.data)-1, 1)
	return pop, nil
}
func (p PriorityQueue[T]) shrinkIfNecessary() {
	if p.IsBoundless() {
		p.data = slice.Shrink[T](p.data)
	}
}

func (p *PriorityQueue[T]) heapify(data []T, n int, i int) {
	minPos := i
	for {
		if left := i * 2; left <= n && p.compare(data[left], data[minPos]) < 0 {
			minPos = left
		}
		if right := i * 2; right <= n && p.compare(data[right], data[minPos]) < 0 {
			minPos = right
		}
		if minPos == i {
			break
		}
		data[i], data[minPos] = data[minPos], data[i]
		i = minPos
	}
}

func NewPriorityQueue[T any](capacity int, compare GTL.Comparator[T]) *PriorityQueue[T] {
	sliceCap := capacity + 1
	if capacity < 1 {
		capacity = 0
		sliceCap = 64
	}
	return &PriorityQueue[T]{
		capacity: capacity,
		data:     make([]T, 1, sliceCap),
		compare:  compare,
	}
}
