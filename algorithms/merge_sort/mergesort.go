package merge_sort

import (
  "math"
)

func MergeSort(values []int) {
  helper := make([]int, len(values))
  mergeSort(values, helper, 0, len(values) - 1)
}

func mergeSort(values []int, helper []int, low int, high int) {
  if low < high {
    middle := int(math.Floor(float64(low + high) / 2))

    mergeSort(values, helper, low, middle)
    mergeSort(values, helper, middle + 1, high)

    merge(values, helper, low, middle, high)
  }
}

func merge(values []int, helper []int, low int, middle int, high int) {
  // Copy both halves into a helper array
  for i := low; i <= high; i++ {
      helper[i] = values[i]
  }

  helperLeft := low
  helperRight := middle + 1
  current := low

  // Iterate through the helper array. Compare the left and right half,
  // copying back the smaller element from the two halves into the original array.
  for helperLeft <= middle && helperRight <= high {
    if helper[helperLeft] <= helper[helperRight] {
      values[current] = helper[helperLeft]
      helperLeft++
    } else {
      values[current] = helper[helperRight]
      helperRight++
    }

    current++
  }

  remaining := middle - helperLeft
  for i := 0; i <= remaining; i++ {
    values[current + i] = helper[helperLeft + i]
  }
}
