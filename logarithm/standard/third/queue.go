package third

import "errors"

/**
使用数组实现的循环队列（静态队列）
*/
type queue struct {
	data  []int
	start int //队首
	end   int //队尾
	size  int //队伍的长度
}

func NewQueue(size int) (*queue, error) {
	if size <= 0 {
		return nil, errors.New("size must be positive integer")
	}
	return &queue{
		data:  make([]int, size+1),
		start: 0,
		end:   0,
		size:  size + 1,
	}, nil
}

func (queue *queue) isEmpty() bool {
	return queue.start == queue.end
}

func (queue *queue) isFull() bool {
	return queue.start == (queue.end+1)%queue.size
}

/**
入队
*/
func (queue *queue) Enqueue(item int) error {
	if queue.isFull() {
		return errors.New("queue is full")
	}
	queue.data[queue.end] = item
	queue.end = (queue.end + 1) % queue.size
	return nil
}

/**
出队
*/
func (queue *queue) Dequeue() (int, error) {
	if queue.isEmpty() {
		return 0, errors.New("queue is empty")
	}
	data := queue.data[queue.start]
	queue.start = (queue.start + 1) % queue.size
	return data, nil
}
