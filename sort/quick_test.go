package sort

import (
	"fmt"
	"testing"
	"time"

	"algorithm/lib"
)

func TestQuickSort(t *testing.T) {
	arr := []int{1, 2}
	arr = GeneratedRandomArray(10000, 200)
	QuickSort(arr, Asc)
	t.Log(arr)
	QuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{1, 100, 32, 41111, -5}
	QuickSort(arr, Asc)
	t.Log(arr)
	QuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{}
	QuickSort(arr, Asc)
	t.Log(arr)
	QuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{1}
	QuickSort(arr, Asc)
	t.Log(arr)
	QuickSort(arr, Desc)
	t.Log(arr)
}

func TestDoubleWayQuickSort(t *testing.T) {
	arr := []int{1, 2}
	arr = GeneratedRandomArray(20, 200)
	DoubleWayQuickSort(arr, Asc)
	t.Log(arr)
	DoubleWayQuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{1, 100, 32, 41111, -5}
	DoubleWayQuickSort(arr, Asc)
	t.Log(arr)
	DoubleWayQuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{}
	DoubleWayQuickSort(arr, Asc)
	t.Log(arr)
	DoubleWayQuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{1}
	DoubleWayQuickSort(arr, Asc)
	t.Log(arr)
	DoubleWayQuickSort(arr, Desc)
	t.Log(arr)
}

func TestThirdWayQuickSort(t *testing.T) {
	arr := []int{1, 2}
	arr = GeneratedRandomArray(20, 200)
	ThreeWayQuickSort(arr, Asc)
	t.Log(arr)
	ThreeWayQuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{1, 100, 32, 41111, -5}
	ThreeWayQuickSort(arr, Asc)
	t.Log(arr)
	ThreeWayQuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{}
	ThreeWayQuickSort(arr, Asc)
	t.Log(arr)
	ThreeWayQuickSort(arr, Desc)
	t.Log(arr)

	arr = []int{1}
	ThreeWayQuickSort(arr, Asc)
	t.Log(arr)
	ThreeWayQuickSort(arr, Desc)
	t.Log(arr)
}

func TestQuickSortBench(t *testing.T) {
	arr := lib.GeneratedRandomArray(10000, 10000)
	MergeSort(arr, Asc)
	//arr1 := DeepCopySlice(arr)
	//arr2 := DeepCopySlice(arr)
	printUsingTime(arr, Asc, QuickSort, QuickSortRandom, "QuickSort", "QuickSortRandom")
	//printUsingTime(arr, Asc, QuickSort, DoubleWayQuickSort, "QuickSort", "DoubleWayQuickSort")
	//printUsingTime(arr1, Asc, QuickSort, ThreeWayQuickSort, "QuickSort", "ThreeWayQuickSort")
	//printUsingTime(arr2, Asc, DoubleWayQuickSort, ThreeWayQuickSort, "DoubleWayQuickSort", "ThreeWayQuickSort")
}

func printUsingTime(arr []int, compare Compare, f1, f2 func(arr []int, compare Compare), f1Name, f2Name string) {
	arr1 := DeepCopySlice(arr)
	start := time.Now()
	f1(arr, compare)
	fmt.Printf("%s using time=%v\n", f1Name, time.Since(start))
	start = time.Now()
	f2(arr1, compare)
	fmt.Printf("%s using time=%v\n", f2Name, time.Since(start))
	fmt.Println(EqualSlice(arr, arr1))
}
