package merge_sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "with negative numbers",
			nums: []int{0, 15, 4, -100},
			want: []int{-100, 0, 4, 15},
		},
		{
			name: "empty input array",
			nums: []int{},
			want: []int{},
		},
		{
			name: "same number is returned more than once",
			nums: []int{47, 10, 7, 100, 55, 47},
			want: []int{7, 10, 47, 47, 55, 100},
		},
		{
			name: "corner case for input with 1 num",
			nums: []int{456},
			want: []int{456},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort(tt.nums)

			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("want [%v], got [%v]", tt.want, tt.nums)
			}
		})
	}
}
