package sort

/**
插入排序的实现
	原理：
		1.当前数组的游标i从1开始，每次同[0,i-1]做比较，只要当前元素小于区间中的最后元素，那么就交换i和该元素的值
		2.如果当前元素大于等于区间中的最后一个元素，则不进行内循环且游标i自增(即是认为[0,1]已经排序好，接下来进行[0,2]的排序)
		3.一旦交换了元素，则内循环中需要继续做比较，直到i所指的元素插入到了合适的位置
*/
func InsertionSort(arr []int, compare Compare) {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && compare(arr[j], arr[j+1]); j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}
