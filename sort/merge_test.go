package sort

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	MergeSort(arr, Asc)
	t.Log(arr)
	MergeSort(arr, Desc)
	t.Log(arr)

	arr = []int{1, 100, 32, 41111, -5}
	MergeSort(arr, Asc)
	t.Log(arr)
	MergeSort(arr, Desc)
	t.Log(arr)

	arr = []int{}
	MergeSort(arr, Asc)
	t.Log(arr)
	MergeSort(arr, Desc)
	t.Log(arr)

	arr = []int{1}
	MergeSort(arr, Asc)
	t.Log(arr)
	MergeSort(arr, Desc)
	t.Log(arr)
}
