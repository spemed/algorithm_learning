package sort

func MergeSort(arr []int, compare Compare) {
	if len(arr) <= 1 {
		return
	}
	mergeSort(arr, 0, len(arr)-1, compare)
}

func mergeSort(arr []int, left, right int, compare Compare) {
	if left == right {
		return
	}
	mid := left + (right-left)>>1
	//分治算法，左半区间排序
	mergeSort(arr, left, mid, compare)
	//分治算法，右半区间排序
	mergeSort(arr, mid+1, right, compare)
	//对左右区间进行比较
	merge(arr, left, mid, right, compare)
}

func merge(arr []int, left, mid, right int, compare Compare) {
	leftPos, rightPos := left, mid+1
	tempArr := make([]int, right-left+1)
	tempIndex := 0
	for leftPos <= mid && rightPos <= right {
		if !compare(arr[leftPos], arr[rightPos]) {
			tempArr[tempIndex] = arr[leftPos]
			tempIndex++
			leftPos++
		} else {
			tempArr[tempIndex] = arr[rightPos]
			tempIndex++
			rightPos++
		}
	}

	for leftPos <= mid {
		tempArr[tempIndex] = arr[leftPos]
		tempIndex++
		leftPos++
	}

	for rightPos <= right {
		tempArr[tempIndex] = arr[rightPos]
		tempIndex++
		rightPos++
	}

	for i := 0; i < len(tempArr); i++ {
		arr[i+left] = tempArr[i]
	}
}
