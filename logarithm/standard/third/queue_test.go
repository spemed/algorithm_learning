package third

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Enqueue(t *testing.T) {
	queue, err := NewQueue(3)
	assert.NoError(t, err, "failed to new queue")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to Enqueue item 1")
	err = queue.Enqueue(2)
	assert.NoError(t, err, "failed to Enqueue item 2")
	err = queue.Enqueue(3)
	assert.NoError(t, err, "failed to Enqueue item 3")
	err = queue.Enqueue(4)
	assert.NoError(t, err, "failed to Enqueue item 4")
}

func TestQueue_Dequeue(t *testing.T) {
	queue, err := NewQueue(3)
	assert.NoError(t, err, "failed to new queue")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to Enqueue item 1")
	err = queue.Enqueue(2)
	assert.NoError(t, err, "failed to Enqueue item 2")
	err = queue.Enqueue(3)
	assert.NoError(t, err, "failed to Enqueue item 3")

	item, err := queue.Dequeue()
	assert.NoError(t, err, "failed to Dequeue")
	assert.Equal(t, 1, item)

	item, err = queue.Dequeue()
	assert.NoError(t, err, "failed to Dequeue")
	assert.Equal(t, 2, item)

	item, err = queue.Dequeue()
	assert.NoError(t, err, "failed to Dequeue")
	assert.Equal(t, 3, item)

	_, err = queue.Dequeue()
	assert.NoError(t, err, "failed to Dequeue")
}

func TestActiveQueue_IsEmpty(t *testing.T) {
	queue := NewActiveQueue()
	assert.Equal(t, true, queue.IsEmpty())
}

func TestActiveQueue_IsFull(t *testing.T) {
	queue := NewActiveQueue()
	assert.Equal(t, false, queue.IsFull())
}

func TestActiveQueue_Enqueue(t *testing.T) {
	queue := NewActiveQueue()
	err := queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
}

func TestActiveQueue_Dequeue(t *testing.T) {
	queue := NewActiveQueue()

	err := queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")

	err = queue.Enqueue(2)
	assert.NoError(t, err, "failed to enqueue item 2")

	err = queue.Enqueue(3)
	assert.NoError(t, err, "failed to enqueue item 3")

	err = queue.Enqueue(4)
	assert.NoError(t, err, "failed to enqueue item 4")

	err = queue.Enqueue(5)
	assert.NoError(t, err, "failed to enqueue item 5")

	item, err := queue.Dequeue()
	assert.NoError(t, err, "failed to dequeue")
	assert.Equal(t, 1, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 2, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 3, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 4, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 5, item)
	assert.NoError(t, err, "failed to dequeue")

	_, err = queue.Dequeue()
	assert.NoError(t, err, "failed to dequeue")
}

func TestTwoStackQueue_IsEmpty(t *testing.T) {
	queue := NewTwoStackQueue()
	assert.Equal(t, true, queue.IsEmpty())
}

func TestTwoStackQueue_IsFull(t *testing.T) {
	queue := NewTwoStackQueue()
	assert.Equal(t, false, queue.IsFull())
}

func TestTwoStackQueue_Enqueue(t *testing.T) {
	queue := NewTwoStackQueue()
	err := queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
	err = queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")
}

func TestTwoStackQueue_Dequeue(t *testing.T) {
	queue := NewTwoStackQueue()

	err := queue.Enqueue(1)
	assert.NoError(t, err, "failed to enqueue item 1")

	err = queue.Enqueue(2)
	assert.NoError(t, err, "failed to enqueue item 2")

	err = queue.Enqueue(3)
	assert.NoError(t, err, "failed to enqueue item 3")

	err = queue.Enqueue(4)
	assert.NoError(t, err, "failed to enqueue item 4")

	err = queue.Enqueue(5)
	assert.NoError(t, err, "failed to enqueue item 5")

	item, err := queue.Dequeue()
	assert.NoError(t, err, "failed to dequeue")
	assert.Equal(t, 1, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 2, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 3, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 4, item)
	assert.NoError(t, err, "failed to dequeue")

	item, err = queue.Dequeue()
	assert.Equal(t, 5, item)
	assert.NoError(t, err, "failed to dequeue")

	_, err = queue.Dequeue()
	assert.NoError(t, err, "failed to dequeue")
}
