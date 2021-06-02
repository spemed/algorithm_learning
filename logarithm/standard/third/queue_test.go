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
