package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	BubbleSort(arr, Asc)
	t.Log(arr)
	BubbleSort(arr, Desc)
	t.Log(arr)

	arr = []int{1, 100, 32, 41111, -5}
	BubbleSort(arr, Asc)
	t.Log(arr)
	BubbleSort(arr, Desc)
	t.Log(arr)

	arr = []int{}
	BubbleSort(arr, Asc)
	t.Log(arr)
	BubbleSort(arr, Desc)
	t.Log(arr)

	arr = []int{1}
	BubbleSort(arr, Asc)
	t.Log(arr)
	BubbleSort(arr, Desc)
	t.Log(arr)
}
