package sort

import (
	"math/rand"
	"time"
)

const (
	LESS  = -1
	EQUAL = 0
	GREAT = 1
)

type Compare func(a, b int) bool

func Asc(left, right int) bool {
	return left > right
}

func Desc(left, right int) bool {
	return left < right
}

/**
生成随机数组
*/
func GeneratedRandomArray(size, value int) []int {
	arr := make([]int, size)
	for i := range arr {
		n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(value)
		arr[i] = n
	}
	return arr
}

func DeepCopySlice(arr []int) []int {
	temp := make([]int, len(arr))
	for i := range arr {
		temp[i] = arr[i]
	}
	return temp
}

func EqualSlice(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr1[i] {
			return false
		}
	}
	return true
}
