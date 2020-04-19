package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	var (
		nums   []int
		result []int
	)

	nums = []int{6, 5, 4, 3, 2, 1}
	result = []int{1, 2, 3, 4, 5, 6}
	MergeSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Errorf("got [%v], want [%v]", nums, result)
	}

	nums = []int{0, 15, 4, 15, -100, -100}
	result = []int{-100, -100, 0, 4, 15, 15}
	MergeSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Errorf("got [%v], want [%v]", nums, result)
	}

	nums = make([]int, 0)
	result = []int{}
	MergeSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Errorf("got [%v], want [%v]", nums, result)
	}

	nums = nil
	result = nil
	MergeSort(nums)
	if !reflect.DeepEqual(nums, result) {
		t.Errorf("got [%v], want [%v]", nums, result)
	}
}
