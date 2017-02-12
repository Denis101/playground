package quick_sort

import (
  "math"
)

func QuickSort(values []int) {
  quickSort(values, 0, len(values) - 1)
}

func quickSort(values []int, left int, right int) {
  index := partition(values, left, right)

  if left < index - 1 {
    quickSort(values, left, index - 1)
  }

  if index < right {
    quickSort(values, index, right)
  }
}

func partition(values []int, left int, right int) int {
  pivot := values[int(math.Floor(float64(left + right) / 2))]

  for left <= right {
    for values[left] < pivot {
      left++
    }

    for values[right] > pivot {
      right--
    }

    if left <= right {
      values[left], values[right] = values[right], values[left]

      left++
      right--
    }
  }

  return left
}
