package stack

import (
  "testing"
)

func TestPushPop (t *testing.T) {
  expected := 1
  stackTest := NewStack()

  stackTest.Push(expected)
  actual, _ := stackTest.Pop()

  if actual != expected {
    t.Errorf("Popped value %d doesn't match pushed value %d", actual, expected)
  }
}

func TestMin (t *testing.T) {
  expected := 1
  stackTest := NewStack()

  stackTest.Push(1)
  stackTest.Push(2)
  stackTest.Push(3)
  actual, _ := stackTest.Min()

  if actual != expected {
    t.Errorf("Min value %d doesn't match expected value %d", actual, expected)
  }
}

func TestMax (t *testing.T) {
  expected := 3
  stackTest := NewStack()

  stackTest.Push(1)
  stackTest.Push(2)
  stackTest.Push(3)
  actual, _ := stackTest.Max()

  if actual != expected {
    t.Errorf("Max value %d doesn't match expected value %d", actual, expected)
  }
}

func TestMinAfterPop (t *testing.T) {
  expected := 2
  stackTest := NewStack()

  stackTest.Push(3)
  stackTest.Push(2)
  stackTest.Push(1)
  stackTest.Pop()
  actual, _ := stackTest.Min()

  if actual != expected {
    t.Errorf("Min value %d doesn't match expected value %d", actual, expected)
  }
}
