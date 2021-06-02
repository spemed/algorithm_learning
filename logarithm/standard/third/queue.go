package third

import "errors"

type Queue interface {
	IsEmpty() bool
	IsFull() bool
	Enqueue(item int) error
	Dequeue() (int, error)
	Front() (int, error)
}

/**
使用数组实现的循环队列（静态队列）
*/
type staticQueue struct {
	data  []int
	start int //队首
	end   int //队尾
	size  int //队伍的长度
}

func NewQueue(size int) (*staticQueue, error) {
	if size <= 0 {
		return nil, errors.New("size must be positive integer")
	}
	return &staticQueue{
		data:  make([]int, size+1),
		start: 0,
		end:   0,
		size:  size + 1,
	}, nil
}

func (queue *staticQueue) IsEmpty() bool {
	return queue.start == queue.end
}

func (queue *staticQueue) IsFull() bool {
	return queue.start == (queue.end+1)%queue.size
}

/**
入队
*/
func (queue *staticQueue) Enqueue(item int) error {
	if queue.IsFull() {
		return errors.New("queue is full")
	}
	queue.data[queue.end] = item
	queue.end = (queue.end + 1) % queue.size
	return nil
}

func (queue *staticQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return queue.data[queue.start], nil
}

/**
出队
*/
func (queue *staticQueue) Dequeue() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	data := queue.data[queue.start]
	queue.start = (queue.start + 1) % queue.size
	return data, nil
}

type activeQueueNode struct {
	value int
	next  *activeQueueNode
}

type activeQueue struct {
	front *activeQueueNode
	end   *activeQueueNode
}

func NewActiveQueue() *activeQueue {
	node := new(activeQueueNode)
	return &activeQueue{
		front: node,
		end:   node,
	}
}

func (queue *activeQueue) IsEmpty() bool {
	return queue.front == queue.end
}

func (queue *activeQueue) IsFull() bool {
	return false
}

/**
入队
*/
func (queue *activeQueue) Enqueue(item int) error {
	//入队从队尾入
	node := new(activeQueueNode)
	node.value = item
	queue.end.next = node
	queue.end = node
	return nil
}

/**
出队
*/
func (queue *activeQueue) Dequeue() (int, error) {
	//出队从队首出
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	//取得队首的第一个有效元素
	validVal := queue.front.next
	queue.front.next = validVal.next
	//队列取空了,队尾指针挪至队首指针
	if queue.front.next == nil {
		queue.end = queue.front
	}
	return validVal.value, nil
}

func (queue *activeQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return queue.front.next.value, nil
}

/**
双栈模拟队列
*/
type towStackQueue struct {
	inStack  Stack
	outStack Stack
}

func NewTwoStackQueue() *towStackQueue {
	return &towStackQueue{
		inStack:  NewActiveStack(),
		outStack: NewActiveStack(),
	}
}

func (queue *towStackQueue) IsFull() bool {
	return false
}

func (queue *towStackQueue) IsEmpty() bool {
	return queue.inStack.IsEmpty() && queue.outStack.IsEmpty()
}

/**
入队只入inStack
*/
func (queue *towStackQueue) Enqueue(item int) error {
	return queue.inStack.Push(item)
}

/**
出队只出outStack
*/
func (queue *towStackQueue) Dequeue() (int, error) {
	//队空返回错误
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	//outStack从inStack迁移数据到outStack中
	if queue.outStack.IsEmpty() {
		queue.transDataFromInStack2OutStack()
	}
	return queue.outStack.Pop()
}

func (queue *towStackQueue) Front() (int, error) {
	if queue.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	if queue.outStack.IsEmpty() {
		queue.transDataFromInStack2OutStack()
	}
	return queue.outStack.Top()
}

/**
从inStack迁移数据到outStack
*/
func (queue *towStackQueue) transDataFromInStack2OutStack() {
	for !queue.inStack.IsEmpty() {
		item, _ := queue.inStack.Pop()
		_ = queue.outStack.Push(item)
	}
}
