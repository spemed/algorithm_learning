package third

import "errors"

type Stack interface {
	IsEmpty() bool
	IsFull() bool
	Push(item int) error
	Pop() (int, error)
	Top() (int, error)
}

/**
使用数组实现栈(静态栈)
*/

type staticStack struct {
	data    []int
	size    int
	current int
}

func (stack *staticStack) IsEmpty() bool {
	return stack.current == -1
}

func (stack *staticStack) IsFull() bool {
	nextPointer := stack.current + 1
	return stack.size == nextPointer
}

func (stack *staticStack) Push(item int) error {
	if stack.IsFull() {
		return errors.New("stack is full")
	}
	stack.current++
	stack.data[stack.current] = item
	return nil
}

func (stack *staticStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	data := stack.data[stack.current]
	stack.current--
	return data, nil
}

func (stack *staticStack) Top() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return stack.data[stack.current], nil
}

func NewStaticStack(size int) (Stack, error) {
	if size <= 0 {
		return nil, errors.New("size must be positive integer")
	}
	return &staticStack{
		data:    make([]int, size),
		size:    size,
		current: -1,
	}, nil
}

/**
动态栈的链表节点
*/
type activeStackNode struct {
	value int
	next  *activeStackNode
}

/**
使用链表实现的动态栈
*/
type activeStack struct {
	top *activeStackNode //栈顶指针
}

func NewActiveStack() *activeStack {
	return &activeStack{
		top: new(activeStackNode),
	}
}

func (*activeStack) IsFull() bool {
	return false
}

func (stack *activeStack) IsEmpty() bool {
	return stack.top.next == nil
}

func (stack *activeStack) Push(item int) error {
	node := new(activeStackNode)
	node.value = item
	//取得第一个有效节点
	firstNode := stack.top.next
	stack.top.next = node
	node.next = firstNode
	return nil
}

func (stack *activeStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	//top本身只是一个哨兵节点，不存储任何有效值，top的next元素为栈顶的第一个元素
	realNode := stack.top.next
	stack.top.next = realNode.next
	return realNode.value, nil
}

func (stack *activeStack) Top() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return stack.top.next.value, nil
}

/**
给你的栈新增功能，新增一个getMin的api，要求以o(1)的时间复杂度取得栈中的最小元素
*/
type minStack struct {
	data Stack
	min  Stack
}

func (*minStack) IsFull() bool {
	return false
}

func (stack *minStack) IsEmpty() bool {
	return stack.data.IsEmpty()
}

func (stack *minStack) Push(item int) error {
	//如果minStack栈为空,则同时入双栈
	if stack.IsEmpty() {
		if err := stack.data.Push(item); err != nil {
			return err
		}
		if err := stack.min.Push(item); err != nil {
			return err
		}
	}
	//如果minStack栈不为空，则取得min栈的top值进行比较
	minTop, err := stack.min.Top()
	if err != nil {
		return err
	}
	//如果当前值比minTop小,则将当前值入min栈
	//如果当前值比minTop大，则把minTop再入min栈
	value := minTop
	if item < minTop {
		value = item
	}
	//最小值入min栈
	if err = stack.min.Push(value); err != nil {
		return err
	}
	//实际值入data栈
	return stack.data.Push(item)
}

func (stack *minStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	//data栈弹出数据
	item, err := stack.data.Pop()
	if err != nil {
		return 0, err
	}
	//min栈跟随弹出数据
	_, err = stack.min.Pop()
	if err != nil {
		return 0, err
	}
	return item, nil
}

func (stack *minStack) Top() (int, error) {
	return stack.data.Top()
}

/**
获取当前栈中的最小值
*/
func (stack *minStack) GetMin() (int, error) {
	return stack.min.Top()
}

/**
工厂创建minStack的实例
*/
func NewMinStack() *minStack {
	return &minStack{
		data: NewActiveStack(),
		min:  NewActiveStack(),
	}
}

type twoQueueStack struct {
	dataQueue Queue
	helpQueue Queue
}

func NewTwoQueueStack() *twoQueueStack {
	return &twoQueueStack{
		dataQueue: NewActiveQueue(),
		helpQueue: NewActiveQueue(),
	}
}

func (*twoQueueStack) IsFull() bool {
	return false
}

func (stack *twoQueueStack) IsEmpty() bool {
	return stack.dataQueue.IsEmpty() && stack.helpQueue.IsEmpty()
}

//入栈只入data队列
func (stack *twoQueueStack) Push(item int) error {
	return stack.dataQueue.Enqueue(item)
}

//出栈时把除队首元素以外的其他元素都迁移到help队列中
func (stack *twoQueueStack) Pop() (int, error) {
	if stack.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	//用于记录队尾最后一个元素的值
	lastVal := 0
	for {
		//弹出队首元素
		item, _ := stack.dataQueue.Dequeue()
		lastVal = item
		//data队列为空,说明前n-1个元素已经入help队列了
		if stack.dataQueue.IsEmpty() {
			break
		}
		_ = stack.helpQueue.Enqueue(item)
	}
	//交换data和help的引用,下次pop时进行同样的流程
	stack.dataQueue, stack.helpQueue = stack.helpQueue, stack.dataQueue
	return lastVal, nil
}

/**
非常蠢的办法,仅供娱乐哈
*/
func (stack *twoQueueStack) Top() (int, error) {
	return 0, errors.New("no support")
}
