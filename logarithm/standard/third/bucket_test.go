package third

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"algorithm/lib"
	"algorithm/sort"
)

func TestCalculate(t *testing.T) {
	arr := lib.GeneratedRandomArray(20000, 200)
	arr1 := lib.DeepCopySlice(arr)
	source, err := calculate(arr)
	assert.NoError(t, err, "failed to calculate")
	assert.Equal(t, source, compare(arr1))
	t.Log(arr)
}

func compare(arr []int) int {
	sort.QuickSort(arr, sort.Asc)
	global := math.MinInt32
	for i := 1; i < len(arr); i++ {
		temp := arr[i] - arr[i-1]
		if temp > global {
			global = temp
		}
	}
	return global
}
