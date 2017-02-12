package merge_sort

import (
  "fmt"
  "testing"
)

func TestMergeSort(t *testing.T) {
  values := []int{3, 2, 5, 8, 9, 10, 1}

  fmt.Printf("BEFORE: %+v", values)
  MergeSort(values)
  fmt.Printf(", AFTER: %+v\n", values)
}
