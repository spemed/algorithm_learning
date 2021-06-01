package sort

import "testing"

func TestSwitchSort(t *testing.T) {
	arr := []int{1,2,3,4,5}
	SwitchSort(arr,Asc)
	t.Log(arr)
	SwitchSort(arr,Desc)
	t.Log(arr)

	arr = []int{1,100,32,41111,-5}
	SwitchSort(arr,Asc)
	t.Log(arr)
	SwitchSort(arr,Desc)
	t.Log(arr)

	arr = []int{}
	SwitchSort(arr,Asc)
	t.Log(arr)
	SwitchSort(arr,Desc)
	t.Log(arr)

	arr = []int{1}
	SwitchSort(arr,Asc)
	t.Log(arr)
	SwitchSort(arr,Desc)
	t.Log(arr)
}
