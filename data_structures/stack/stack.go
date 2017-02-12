package stack

import (
  "fmt"
)

type stackNode struct {
  data int
  min int
  max int
  next *stackNode
}

type Stack struct {
  head *stackNode
}

func NewStack() *Stack {
  return &Stack{}
}


func (s *Stack) Push(value int) error {
  node := &stackNode{
    data: value,
    min: value,
    max: value,
  }

  if s.head == nil {
    s.head = node
    return nil
  }

  if value > s.head.min {
    node.min = s.head.min
  }

  if value < s.head.max {
    node.max = s.head.max
  }

  tmp := s.head
  node.next = tmp
  s.head = node
  return nil
}

func (s *Stack) Pop() (int, error) {
  if s.head == nil {
    return -1, fmt.Errorf("can't pop from an empty stack")
  }

  tmp := s.head
  s.head = tmp.next
  return tmp.data, nil
}

func (s *Stack) Min() (int, error) {
  if s.head == nil {
    return -1, fmt.Errorf("can't get the min value of an empty stack")
  }

  return s.head.min, nil
}

func (s *Stack) Max() (int, error) {
  if s.head == nil {
    return -1, fmt.Errorf("can't get the max value of an empty stack")
  }

  return s.head.max, nil
}
