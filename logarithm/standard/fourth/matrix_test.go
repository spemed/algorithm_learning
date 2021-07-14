package fourth

import "testing"

func TestPrintEdge1(t *testing.T) {
	printEdge1([][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	})

	printEdge1([][]int{
		{1, 2},
		{6, 3},
		{5, 4},
	})
}
