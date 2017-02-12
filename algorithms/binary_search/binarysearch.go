package binary_search

import (
  "math"
)

func BinarySearch(target int, values []int) int {
  min := 0
  max := len(values)

  for min <= max {
    current := int(math.Floor(float64(min + max) / 2))

    if values[current] < target {
      min = current + 1
    } else if values[current] > target {
      max = current - 1
    } else {
      return current
    }
  }

  return -1
}
