package queue

import (
  "testing"
)

// Should be split into multiple tests but im lazy
func TestQueue(t *testing.T) {
  queue := NewQueue()

  queue.Queue(1)
  queue.Queue(3)
  queue.Queue(2)

  min, _ := queue.Min()
  max, _ := queue.Max()

  if min != 1 {
    t.Errorf("Min is %d when should be 1", min)
  }

  if max != 3 {
    t.Errorf("Max is %d when should be 3", max)
  }

  value, _ := queue.Dequeue()

  if value != 1 {
    t.Errorf("Dequeue should return 1, returned %d", value)
  }

  value, _ = queue.Dequeue()

  if value != 3 {
    t.Errorf("Dequeue should return 3, returned %d", value)
  }

}
