package second

/**
荷兰国旗问题
	要求：给定一个数num和一个数组arr，要求小于num的数都分布在数组的左边，大于num的数都分布在数组右边，
		等于num的数必须是左右区域的分界线
	双指针法,最终返回等值区间的左闭区间和右闭区间
*/
func NetherlandsFlag(num int, arr []int) (start, end int) {
	partition := num
	left, right := -1, len(arr)-1
	currentIndex := 0
	for currentIndex < right {
		if partition == arr[currentIndex] {
			currentIndex++
			continue
		}
		if arr[currentIndex] < partition {
			left++
			arr[left], arr[currentIndex] = arr[currentIndex], arr[left]
			currentIndex++
			continue
		}
		if arr[currentIndex] > partition {
			right--
			arr[currentIndex], arr[right] = arr[right], arr[currentIndex]
		}
	}
	return left + 1, right - 1
}
