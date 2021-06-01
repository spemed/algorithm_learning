package sort

import "testing"

func TestInsertionSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	InsertionSort(arr, Asc)
	t.Log(arr)
	InsertionSort(arr, Desc)
	t.Log(arr)

	arr = []int{1, 100, 32, 41111, -5, 21232323, 22323, 111}
	InsertionSort(arr, Asc)
	t.Log(arr)
	InsertionSort(arr, Desc)
	t.Log(arr)

	arr = []int{}
	InsertionSort(arr, Asc)
	t.Log(arr)
	InsertionSort(arr, Desc)
	t.Log(arr)

	arr = []int{1}
	InsertionSort(arr, Asc)
	t.Log(arr)
	InsertionSort(arr, Desc)
	t.Log(arr)
}
