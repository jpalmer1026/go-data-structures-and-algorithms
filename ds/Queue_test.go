package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_isEmpty(t *testing.T) {
	q := NewQueue(10)

	assert.False(t, q.isEmpty())
	q.Dequeue()
	assert.True(t, q.isEmpty())
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue(10)

	q.Enqueue(20)
	q.Enqueue(30)

	assert.Equal(t, 10, q.Peek().value)
}

func TestQueue_Dequeue(t *testing.T) {
	q := NewQueue(10)

	assert.Equal(t, 10, q.first.value)

	q.Enqueue(20)
	q.Enqueue(30)

	q.Dequeue()
	assert.Equal(t, 20, q.first.value)

	q.Dequeue()
	assert.Equal(t, 30, q.first.value)

	q.Dequeue()
	assert.Nil(t, q.first)
}
