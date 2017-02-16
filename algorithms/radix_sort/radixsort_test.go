package radix_sort

import (
  "fmt"
  "testing"
)

func TestRadixSort(t *testing.T) {
  values := []int{552, 32, 48, 7, 2, 68, 987, 1, 46}

  fmt.Printf("BEFORE: %+v", values)
  RadixSort(values)
  fmt.Printf(", AFTER: %+v", values)
}
