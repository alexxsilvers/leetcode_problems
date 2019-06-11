package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct{
		list []int
		item int
		want int
	}{
		{
			list: []int{1, 32, 45, 67, 88, 140, 754},
			item: 88,
			want: 4,
		},
		{
			list: []int{1, 3, 5, 7, 9},
			item: 3,
			want: 1,
		},
		{
			list: []int{1, 3, 5, 7, 9},
			item: -1,
			want: -1,
		},
		{
			list: []int{1},
			item: 44,
			want: -1,
		},
		{
			list: []int{},
			item: 10,
			want: -1,
		},
		{
			list: []int{10, 1, 2, 10},
			item: 10,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := BinarySearch(tt.list, tt.item)
			if got != tt.want {
				t.Errorf("BinarySearch() not passed. List %#v, item %d, got %d, want %d", tt.list, tt.item, got, tt.want)
			}
		})
	}
}
