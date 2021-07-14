package third

import (
	"errors"
	"math"
)

/**
给定一个数组，求出该数组排序后，相邻两数最大的差值，且算法的时间复杂度必须为o(n)
	思路：利用桶排序的思路，若数组存在n个元素，则创建n+1个桶。
		遍历一遍数组得到最小和最大两个元素，如果min==max，则说明该数组都是等值元素，最大差值为0
		否则则让min和max分别占据第一个和最后一个桶
		range = (max-min)/(n+1) 得到每个桶负责的range
		再次遍历一遍数组，将当前数组的元素除以range得到应该放置在哪个桶内
		每个桶只记录经过这个桶的最大和最小元素
		因为存在N+1个桶，所以填充完所有数据后，必定有一个桶会是空桶。
		(
			已知在桶1内，存在x，则x的取值范围为[x,x+range]，取x1,x2自桶1，易证明 -range <= x1-x2 <= range
			假设该桶的下一个桶桶2为空桶，2桶的下一个桶3的取值范围为[x+2*range,x+3*range],取x1来自桶3，x2来自桶1，易证明 range <= x1-x2 <= 3 * range
 			因此可以证明，当存在空桶的情况下，排序后的两数相减得到的最大差值只会在两个不同的桶之间产生，而不会来自同一个桶
			这就是我们需要N+1个桶的原因，为了构造一个空桶
		)
        最后再次遍历我们的桶数组，只要当前桶为空桶则跳过，非空桶则计算当前桶的最小值和上一个桶的最大值，将差值更新到全局变量中（如果该差值一直在增大）
*/

type bucketItem struct {
	min        int
	max        int
	hasElement bool
}

/**
往bucket中追加元素
*/
func (item *bucketItem) add(data int) {
	item.hasElement = true
	if item.min > data {
		item.min = data
	}
	if item.max < data {
		item.max = data
	}
}

/**
返回下一个桶的最小值减去当前桶的最大值的差值
*/
func (item *bucketItem) compare(next bucketItem) int {
	return next.min - item.max
}

func newBucketItem() bucketItem {
	return bucketItem{
		min:        math.MaxInt32,
		max:        math.MinInt32,
		hasElement: false,
	}
}

type bucket struct {
	arr        []bucketItem
	zeroResult bool
}

func calculate(arr []int) (int, error) {

	if len(arr) == 0 {
		return 0, errors.New("arr could not be empty")
	}

	size := len(arr)
	bucketArr := make([]bucketItem, size+1)
	//初始化桶
	for i := range bucketArr {
		bucketArr[i] = newBucketItem()
	}

	min, max := math.MaxInt32, math.MinInt32
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	if min == max {
		return 0, nil
	}

	bucketArr[0].add(min)
	bucketArr[size].add(max)
	rangeVal := (max - min) / (size + 1)
	for _, v := range arr {
		index := (v - min) / (rangeVal + 1)
		bucketArr[index].add(v)
	}

	lastIndex := 0
	global := math.MinInt32
	for currentIndex := 1; currentIndex < len(bucketArr); currentIndex++ {
		if !bucketArr[currentIndex].hasElement {
			continue
		}
		temp := bucketArr[lastIndex].compare(bucketArr[currentIndex])
		if temp > global {
			global = temp
		}
		lastIndex = currentIndex
	}
	return global, nil
}
