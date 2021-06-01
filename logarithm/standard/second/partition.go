package second

/**
算法要求：
	给定一个随机输入num，和一个数组arr。要求比num小的元素都聚集在数组的左侧，比num大的元素都聚集在数组的右侧
思路：
	1.minIndex从-1开始，[0,minIndex]代表所有小于num的元素所在的区间。startIndex从0开始，向右滑动直到达到数组边界
	2.每当发现当前的数值比number大，则startIndex++。
	3.每当发现当前的数值小于number,则将minIndex的下一位同startIndex交换
      因[0,minIndex]总是小于num,因此minIndex+1所指的元素必定大于num,
      并将minIndex自增，增加最小区间的范围
*/

func Partition(number int, arr []int) {
	minIndex := -1
	for startIndex := 0; startIndex < len(arr); startIndex++ {
		if arr[startIndex] >= number {
			continue
		}
		minIndex++
		arr[minIndex], arr[startIndex] = arr[startIndex], arr[minIndex]
	}
}
