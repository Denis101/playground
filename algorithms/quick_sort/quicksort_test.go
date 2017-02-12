package quick_sort

import (
  "fmt"
  "testing"
)

func TestQuickSort(t *testing.T) {
  values := []int{3, 2, 5, 8, 9, 10, 1}

  fmt.Printf("BEFORE: %+v", values)
  QuickSort(values)
  fmt.Printf(", AFTER: %+v\n", values)
}
