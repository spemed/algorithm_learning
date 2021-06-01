package sort

/**
	选择排序的实现
		原理：
			1.确定需要比较的轮数，3个数只需要比较2轮，2个数只需要比较1轮，因此设外部循环条件 i从[0,len(arr)-1)
			2.内部循环从i+1开始循环至数组的最后一个元素，设置一个局部变量k记录最小值/最大值元素出现的位置，循环开始之前预设当前i所指向的元素为最小值/最大值。
			3.内部循环得到的元素依次和k指向的元素作比较，如果当前元素比k指向的元素小/大，就用当前元素的下标覆盖k的值（实际上是将交换元素的操作滞后了）
			4.内部循环结束时，如果当前循环变量i的值和记录最小/最大变量的下标k的值不一致，则交换则两个下标上的数组元素，此时将最小/最大值移动到了数组开头
			5.此时i的值自增，如果i的当前值已经到达了数组的最后一个元素，则退出循环（只剩下一个元素了没有比较的意义）
			6.否则在除去已排序元素的剩余区间继续做1~4操作
		最好时间复杂度：O(n^2）
		最坏时间复杂度: O(n^2)
		平均时间复杂度: O(n^2)
	    是否是稳定的排序：是（涉及两两比较的排序算法一定是稳定的排序）

	和排序算法做比较
	1.都是稳定的排序算法
	2.无论什么情况都需要O(n^2)的时间复杂度
	3.通过动态记录下标以及比较的方式减少了移动数组元素的开销，在数据量较少，待排序元素较为无序的情况下比冒泡算法表现优秀。
 */

func SwitchSort(arr[] int,compare Compare) {
	for i:=0;i<len(arr)-1;i++ {
		k := i
		for j:=i+1;j<len(arr);j++ {
			if compare(arr[k],arr[j]) {
				k = j
			}
		}
		if k != i {
			arr[i],arr[k] = arr[k],arr[i]
		}
	}
}
