package sort

import (
	"math/rand"
	"time"
)

/**
单路快排的原理
	1.选定数组中的第一个元素作为partition
	2.数组的滑动游标leftPos从第二个元素开始，设置结束游标为rightPos
	3.循环启动，只要leftPost <= rightPos,则循环保持
	4.判断leftPos指向的元素是否比partition小,是的话将leftPos++
	5.如果leftPos指向的元素比partition大，则置换leftPos所指向的元素和rightPost所指向的元素，rightPos--【这一步是为了把比partition大的元素挪到数组后边】
	6.循环4.5，直到3的循环条件被打破，此时rightPos比leftPos小1，将partition所在的下标同rightPos做交换。此时rightPos右侧的数都比partition大，rightPos左侧的数都比partition小
	7.返回partition，对[left,partition),[partition+1,right]两个分区各自进行快排
*/
func QuickSort(arr []int, compare Compare) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1, compare)
}

func quickSort(arr []int, left, right int, compare Compare) {
	if left >= right {
		return
	}
	//进行分区排序
	partition := handle(arr, left, right, compare)
	quickSort(arr, left, partition-1, compare)
	quickSort(arr, partition+1, right, compare)

}

func handle(arr []int, start, end int, compare Compare) int {
	leftPos, rightPos := start+1, end
	compareValue := arr[start]
	for leftPos <= rightPos {
		if compare(arr[leftPos], compareValue) {
			arr[leftPos], arr[rightPos] = arr[rightPos], arr[leftPos]
			rightPos--
		} else {
			leftPos++
		}
	}
	arr[start], arr[rightPos] = arr[rightPos], arr[start]
	return rightPos
}

/**
二路快排，左右指针，同时遍历分别向左右逼近减少循环次数
*/
func DoubleWayQuickSort(arr []int, compare Compare) {
	if len(arr) <= 1 {
		return
	}
	doubleWayQuickSort(arr, 0, len(arr)-1, compare)
}

/**
递归子任务，所以一定要有退出条件
*/
func doubleWayQuickSort(arr []int, start, end int, compare Compare) {
	//下标出现负数，或者当前分配到的处理任务中只需要处理数组的一个元素时，递归任务可以结束了
	if start >= end {
		return
	}
	mid := doubleWayPartition(arr, start, end, compare)
	doubleWayQuickSort(arr, start, mid-1, compare)
	doubleWayQuickSort(arr, mid+1, end, compare)
}

func doubleWayPartition(arr []int, start, end int, compare Compare) int {
	partition := arr[start]
	left, right := start+1, end
	for {

		for left <= end && !compare(arr[left], partition) {
			left++
		}

		for right >= start && compare(arr[right], partition) {
			right--
		}

		if left > right {
			break
		}

		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	arr[right], arr[start] = arr[start], arr[right]
	return right
}

/**
三路快排+随机数优化
随机选择的partition通过概率统计得到了一个较为稳定的数学期望，使得快排的分区时间复杂度能保证在o(NlogN)
三路快排的双指针减少了剩余分区的数组长度，使得处理单个的分区逻辑时常数项耗时更短(因为统计出等值区间使得左右子区间的比较更短了)
*/
//fixme 算法有bug，记得改一下
//[]int{99, 47, 59, 65}
//todo 找到错误了，以后写这种带递归的算法，要取子任务的边界时一定只能从父任务传递的实际参数中去演算，而不能直接依赖预设值
//todo 比如
/**
   func threeWayPartition(arr []int, start, end int, compare Compare) (int, int) {
		partition := arr[end]
		less, more := start-1, end
	}
	less和more的取值一定只能来自实际参数start和end的运算，而不是自己默认的数组开头(0)或者数组结尾(len(arr)-1)
*/
func ThreeWayQuickSort(arr []int, compare Compare) {
	if len(arr) <= 1 {
		return
	}
	threeWayQuickSort(arr, 0, len(arr)-1, compare)
}

/**
递归子任务，所以一定要有退出条件
*/
func threeWayQuickSort(arr []int, start, end int, compare Compare) {
	//下标出现负数，或者当前分配到的处理任务中只需要处理数组的一个元素时，递归任务可以结束了
	if start >= end {
		return
	}
	left, right := threeWayPartition(arr, start, end, compare)
	threeWayQuickSort(arr, start, left-1, compare)
	threeWayQuickSort(arr, right+1, end, compare)
}

func threeWayPartition(arr []int, start, end int, compare Compare) (int, int) {
	//todo 需要用随机数生成函数置换数组最后一个下标
	//暂取最后一个元素作为partition
	partition := arr[end]
	//设置左右指针，左指针在当前数组之外
	left, right := start-1, end
	//设置遍历游标
	currentIndex := start
	//只要游标尚未到达right，则循环进行
	for currentIndex < right {
		//当前元素和partition相等，则游标自增
		if arr[currentIndex] == partition {
			currentIndex++
			continue
		}
		if !compare(arr[currentIndex], partition) {
			left++
			arr[left], arr[currentIndex] = arr[currentIndex], arr[left]
			currentIndex++
			continue
		}
		right--
		arr[currentIndex], arr[right] = arr[right], arr[currentIndex]
	}

	arr[currentIndex], arr[end] = arr[end], arr[currentIndex]
	return left + 1, right
}

func QuickSortRandom(arr []int, compare Compare) {
	if len(arr) <= 1 {
		return
	}
	quickSortRandom(arr, 0, len(arr)-1, compare)
}

func quickSortRandom(arr []int, left, right int, compare Compare) {
	if left >= right {
		return
	}
	//进行分区排序
	partition := handleRandom(arr, left, right, compare)
	quickSortRandom(arr, left, partition-1, compare)
	quickSortRandom(arr, partition+1, right, compare)
}

func handleRandom(arr []int, start, end int, compare Compare) int {
	leftPos, rightPos := start+1, end
	index := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(end + 1)
	arr[index], arr[start] = arr[start], arr[index]
	compareValue := arr[start]
	for leftPos <= rightPos {
		if compare(arr[leftPos], compareValue) {
			arr[leftPos], arr[rightPos] = arr[rightPos], arr[leftPos]
			rightPos--
		} else {
			leftPos++
		}
	}
	arr[start], arr[rightPos] = arr[rightPos], arr[start]
	return rightPos
}
