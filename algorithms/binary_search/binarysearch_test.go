package binary_search

import (
  "fmt"
  "testing"
)

func TestBinarySearch(t *testing.T) {
  fmt.Printf("RESULT: %d\n", BinarySearch(0, []int{0}))
  fmt.Printf("RESULT: %d\n", BinarySearch(1, []int{0, 1, 2, 3, 4, 5}))
  fmt.Printf("RESULT: %d\n", BinarySearch(4, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
  fmt.Printf("RESULT: %d\n", BinarySearch(25, []int{11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239}))
  fmt.Printf("RESULT: %d\n", BinarySearch(99, []int{11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239}))
  fmt.Printf("RESULT: %d\n", BinarySearch(3255, []int{11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239}))
  fmt.Printf("RESULT: %d\n", BinarySearch(11011, []int{11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239}))
  fmt.Printf("RESULT: %d\n", BinarySearch(101239, []int{11, 25, 48, 52, 78, 99, 101, 173, 256, 567, 1024, 1564, 3255, 4387, 11011, 13245, 23456, 101239}))
}
