package third

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {
	stack, err := NewStaticStack(3)
	assert.NoError(t, err, "failed to new stack")
	err = stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")
	err = stack.Push(4)
	assert.NoError(t, err, "failed to push item 4")
}

func TestStack_Pop(t *testing.T) {
	stack, err := NewStaticStack(3)
	assert.NoError(t, err, "failed to new stack")
	err = stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")

	item, err := stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 3)

	item, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 2)

	item, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 1)

	_, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
}

/**
测试弹出为空值
*/
func TestActiveStack_Pop_Empty(t *testing.T) {
	stack := NewActiveStack()
	_, err := stack.Pop()
	assert.NoError(t, err, "failed to pop")
}

/**
测试push
*/
func TestActiveStack_Push(t *testing.T) {
	stack := NewActiveStack()
	err := stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")
}

func TestActiveStack_PopAfterPush(t *testing.T) {

	stack := NewActiveStack()
	err := stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")

	item, err := stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 3)

	item, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 2)

	item, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 1)

	_, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
}

func TestActiveStack_TopAfterPush(t *testing.T) {

	stack := NewActiveStack()
	err := stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")

	item, err := stack.Top()
	assert.NoError(t, err, "failed to get top")
	assert.Equal(t, item, 3)

	_, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")

	item, err = stack.Top()
	assert.NoError(t, err, "failed to get top")
	assert.Equal(t, item, 2)

	_, err = stack.Pop()

	item, err = stack.Top()
	assert.NoError(t, err, "failed to get top")
	assert.Equal(t, item, 1)
}

func TestNewMinStack(t *testing.T) {
	stack := NewMinStack()
	err := stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(-1)
	assert.NoError(t, err, "failed to push item -1")
	err = stack.Push(10085)
	assert.NoError(t, err, "failed to push item 10085")
	err = stack.Push(223)
	assert.NoError(t, err, "failed to push item 223")
	err = stack.Push(3341)
	assert.NoError(t, err, "failed to push item 3341")

	item, err := stack.GetMin()
	assert.NoError(t, err, "failed to get min")
	assert.Equal(t, -1, item)

	popVal, err := stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, 3341, popVal)

	item, err = stack.GetMin()
	assert.NoError(t, err, "failed to get min")
	assert.Equal(t, -1, item)

	popVal, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, 223, popVal)

	item, err = stack.GetMin()
	assert.NoError(t, err, "failed to get min")
	assert.Equal(t, -1, item)

	popVal, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, 10085, popVal)

	item, err = stack.GetMin()
	assert.NoError(t, err, "failed to get min")
	assert.Equal(t, -1, item)

	popVal, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, -1, popVal)

	item, err = stack.GetMin()
	assert.NoError(t, err, "failed to get min")
	assert.Equal(t, 1, item)
}

/**
测试弹出为空值
*/
func TestTwoQueueStack_Pop_Empty(t *testing.T) {
	stack := NewTwoQueueStack()
	_, err := stack.Pop()
	assert.NoError(t, err, "failed to pop")
}

/**
测试push
*/
func TestTwoQueueStack_Push(t *testing.T) {
	stack := NewTwoQueueStack()
	err := stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")
}

func TestTwoQueueStack_PopAfterPush(t *testing.T) {

	stack := NewTwoQueueStack()
	err := stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")

	item, err := stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 3)

	item, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 2)

	item, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")
	assert.Equal(t, item, 1)

	//
	//_, err = stack.Pop()
	//assert.NoError(t, err, "failed to pop")
}

func TestTwoQueueStack_TopAfterPush(t *testing.T) {

	stack := NewTwoQueueStack()
	err := stack.Push(1)
	assert.NoError(t, err, "failed to push item 1")
	err = stack.Push(2)
	assert.NoError(t, err, "failed to push item 2")
	err = stack.Push(3)
	assert.NoError(t, err, "failed to push item 3")

	item, err := stack.Top()
	assert.NoError(t, err, "failed to get top")
	assert.Equal(t, item, 3)

	_, err = stack.Pop()
	assert.NoError(t, err, "failed to pop")

	item, err = stack.Top()
	assert.NoError(t, err, "failed to get top")
	assert.Equal(t, item, 2)

	_, err = stack.Pop()

	item, err = stack.Top()
	assert.NoError(t, err, "failed to get top")
	assert.Equal(t, item, 1)
}
