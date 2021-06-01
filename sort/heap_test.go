package sort

import "testing"

func TestNewMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap()
	maxHeap.Insert(1)
	maxHeap.Insert(100)
	maxHeap.Insert(50)
	maxHeap.Insert(40)
	maxHeap.Insert(23)
	maxHeap.Insert(1)
	maxHeap.Insert(4444)
	maxHeap.Insert(22221111)
	maxHeap.Insert(123123123)
	maxHeap.Insert(1000000000)
	maxHeap.Insert(232)
	for {
		val, err := maxHeap.Pop()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(val)
	}
}

func TestNewMinHeap(t *testing.T) {
	minHeap := NewMinHeap()
	minHeap.Insert(1)
	minHeap.Insert(100)
	minHeap.Insert(50)
	minHeap.Insert(40)
	minHeap.Insert(23)
	minHeap.Insert(1)
	minHeap.Insert(4444)
	minHeap.Insert(22221111)
	minHeap.Insert(232)
	minHeap.Insert(1)
	minHeap.Insert(4444)
	minHeap.Insert(22221111)
	minHeap.Insert(123123123)
	minHeap.Insert(1000000000)
	minHeap.Insert(232)
	for {
		val, err := minHeap.Pop()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(val)
	}
}

func TestHeapify(t *testing.T) {
	arr := GeneratedRandomArray(10, 20)
	maxHeap := Heapify(arr, MinHeap)
	t.Log(arr)

	for {
		val, err := maxHeap.Pop()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(val)
	}
}

func TestHeap_Sort(t *testing.T) {
	arr := GeneratedRandomArray(20, 200)
	minHeap := Heapify(arr, MinHeap)
	t.Log(minHeap.Sort())

	arr = GeneratedRandomArray(20, 200)
	maxHeap := Heapify(arr, MaxHeap)
	t.Log(maxHeap.Sort())
}

func TestHeap_FindTopK(t *testing.T) {
	arr := GeneratedRandomArray(20, 200)
	t.Log(arr)
	minHeap := Heapify(arr, MinHeap)
	t.Log(minHeap.FindTopK(3))

	arr = GeneratedRandomArray(20, 200)
	t.Log(arr)
	maxHeap := Heapify(arr, MaxHeap)
	t.Log(maxHeap.FindTopK(3))
}
