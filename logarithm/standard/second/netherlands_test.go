package second

import (
	"testing"

	"algorithm/lib"
)

func TestNetherlandsFlag(t *testing.T) {
	arr := lib.GeneratedRandomArray(10, 40)
	t.Log(NetherlandsFlag(8, arr))
	t.Log(arr)
}
