package radix_sort

func RadixSort(values []int) {
  max := getMax(values)

  for exp := 1; max / exp > 0; exp *= 10 {
    countSort(values, exp)
  }
}

func countSort(values []int, exp int) {
  output := make([]int, len(values))
  i := 0

  count := make([]int, 10)

  for i := 0; i < len(values); i++ {
    count[(values[i]/exp) % 10]++
  }

  for i := 1; i < 10; i++ {
    count[i] += count[i - 1]
  }

  for i = len(values) - 1; i >= 0; i-- {
    output[count[(values[i] / exp) % 10] - 1] = values[i]
    count[(values[i] / exp) % 10]--
  }

  for i := 0; i < len(values); i++ {
    values[i] = output[i]
  }
}


func getMax(values []int) int {
  max := values[0]

  for i := 0; i < len(values); i++ {
    if values[i] > max {
      max = values[i]
    }
  }

  return max
}
