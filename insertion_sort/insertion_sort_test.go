package insertion_sort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct{
		list []int
		want []int
	}{
		{
			list: []int{3, 7, 1, 20, 44, 14, 190, 99},
			want: []int{1, 3, 7, 14, 20, 44, 99, 190},
		},
		{
			list: []int{56, 345, 23, 18, 2},
			want: []int{2, 18, 23, 56, 345},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := InsertionSort(tt.list)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertionSort() want %#v got %#v", tt.want, got)
			}
		})
	}
}
