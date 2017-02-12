package queue

import (
  "fmt"
)

type queueNode struct {
  data int
  min int
  max int
  next *queueNode
}

type Queue struct {
  head *queueNode
  tail *queueNode
}

func NewQueue() *Queue {
  return &Queue{}
}

func (q *Queue) Queue(value int) error {
  node := &queueNode{
    data: value,
    min: value,
    max: value,
  }

  if q.head == nil {
    q.head = node

    if q.tail == nil {
      q.tail = q.head
    }

    return nil
  }

  if value > q.head.min {
    node.min = q.head.min
  }

  if value < q.head.max {
    node.max = q.head.max
  }

  tmp := q.head
  tmp.next = node
  q.head = node

  if q.tail == nil {
    q.tail = q.head
  }

  return nil
}

func (q *Queue) Dequeue() (int, error) {
  if q.tail == nil {
    return -1, fmt.Errorf("can't dequeue from an empty queue")
  }

  tmp := q.tail
  q.tail = tmp.next
  return tmp.data, nil
}

func (q *Queue) Min() (int, error) {
  if q.head == nil {
    return -1, fmt.Errorf("can't get the min value of an empty queue")
  }

  return q.head.min, nil
}

func (q *Queue) Max() (int, error) {
  if q.head == nil {
    return -1, fmt.Errorf("can't get the max value of an empty queue")
  }

  return q.head.max, nil
}
