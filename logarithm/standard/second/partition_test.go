package second

import (
	"testing"

	"algorithm/lib"
)

func TestPartition(t *testing.T) {
	arr := lib.GeneratedRandomArray(40, 10)
	Partition(20, arr)
	t.Log(arr)
}
