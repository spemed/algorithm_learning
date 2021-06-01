package sort

import (
	"errors"
)

const (
	MaxHeap = 1
	MinHeap = 2
)

type Heap struct {
	data     []int
	heapType int
	length   int
	Compare
}

func newHeap(heapType int) Heap {
	var compare Compare
	if heapType == MaxHeap {
		compare = Desc
	}
	if heapType == MinHeap {
		compare = Asc
	}
	return Heap{
		data:     make([]int, 10),
		heapType: heapType,
		Compare:  compare,
	}
}

func NewMaxHeap() Heap {
	return newHeap(MaxHeap)
}

func NewMinHeap() Heap {
	return newHeap(MinHeap)
}

func (heap *Heap) Insert(value int) {
	//动态扩容
	if heap.length >= len(heap.data) {
		newData := make([]int, 2*heap.length)
		for i := 0; i < heap.length; i++ {
			newData[i] = heap.data[i]
		}
		heap.data = newData
	}
	heap.data[heap.length] = value
	heap.length++
	heap.shiftUp()
}

func (heap *Heap) shiftUp() {
	//取得最后一个叶子节点的坐标
	lastIndex := heap.length - 1
	//取得父节点的坐标
	fatherIndex := getFatherIndexBySonIndex(lastIndex)
	for fatherIndex >= 0 {
		if heap.Compare(heap.data[fatherIndex], heap.data[lastIndex]) {
			heap.data[lastIndex], heap.data[fatherIndex] = heap.data[fatherIndex], heap.data[lastIndex]
			lastIndex = fatherIndex
			fatherIndex = getFatherIndexBySonIndex(lastIndex)
		} else {
			break
		}
	}
}

func (heap *Heap) Top() (int, error) {
	if heap.length == 0 {
		return 0, errors.New("empty heap")
	}
	return heap.data[0], nil
}

func (heap *Heap) Pop() (int, error) {
	if heap.length == 0 {
		return 0, errors.New("empty heap")
	}
	top := heap.data[0]
	heap.data[0] = heap.data[heap.length-1]
	heap.length--
	heap.shiftDown()
	return top, nil
}

func (heap *Heap) shiftDown() {
	//取得最后一个叶子节点的坐标
	shiftDown(heap.data, 0, heap.length-1, heap.Compare)
	return
}

func getFatherIndexBySonIndex(sonIndex int) int {
	if sonIndex%2 == 1 {
		return (sonIndex - 1) >> 1
	} else {
		return (sonIndex - 2) >> 1
	}
}

func shiftDown(arr []int, start, end int, compare Compare) {
	//取得最后一个叶子节点的坐标
	startIndex, lastIndex := start, end
	j := startIndex*2 + 1

	for j <= lastIndex {
		if j+1 <= end && compare(arr[j], arr[j+1]) {
			j = j + 1
		}
		if !compare(arr[startIndex], arr[j]) {
			break
		}
		arr[startIndex], arr[j] = arr[j], arr[startIndex]
		startIndex = j
		j = startIndex*2 + 1
	}
	return
}

/**
将数组原地置换成堆
*/
func Heapify(arr []int, heapType int) Heap {
	lastIndex := len(arr) - 1
	var compare Compare
	if heapType == MaxHeap {
		compare = Desc
	}
	if heapType == MinHeap {
		compare = Asc
	}
	curParentIndex := (lastIndex - 1) >> 1
	for curParentIndex >= 0 {
		shiftDown(arr, curParentIndex, lastIndex, compare)
		curParentIndex--
	}
	return Heap{
		data:     arr,
		heapType: heapType,
		length:   len(arr),
		Compare:  compare,
	}
}

/**
堆排序
*/
func (heap *Heap) Sort() []int {
	return heap.FindTopK(heap.length)
}

/**
@param sortNumber int 求前k位
*/
func (heap *Heap) FindTopK(k int) []int {
	if heap.length == 0 {
		return nil
	}
	tempArr := DeepCopySlice(heap.data)
	reserveArr := make([]int, k)
	for i := 0; i < k; i++ {
		if i >= heap.length {
			break
		}
		compareIndex := heap.length - 1 - i
		tempArr[compareIndex], tempArr[0] = tempArr[0], tempArr[compareIndex]
		shiftDown(tempArr, 0, compareIndex-1, heap.Compare)
		reserveArr[i] = tempArr[compareIndex]
	}
	return reserveArr
}

/**
通过大根堆+小跟堆 实现从实时输入的流中获取中位数
原理：
双堆的heapSize之差超过1，heap1-heap2 > 1，则弹最大堆/最小堆的第一个元素到对向堆
*/
