package main

import "testing"

func Test_checkLevel(t *testing.T) {

	tests := []struct {
		name   string
		levels []int
		want   bool
	}{
		{
			name:   "ascending",
			levels: []int{1, 2, 3},
			want:   true,
		},
		{
			name:   "descending",
			levels: []int{3, 2, 1},
			want:   true,
		},
		{
			name:   "stops ascending",
			levels: []int{1, 2, 3, 3},
			want:   false,
		},
		{
			name:   "stops descending",
			levels: []int{3, 2, 1, 3},
			want:   false,
		},
		{
			name:   "ascending diff up to 3",
			levels: []int{1, 2, 3, 6},
			want:   true,
		},
		{
			name:   "descending diff up to 3",
			levels: []int{6, 3, 2, 1},
			want:   true,
		},
		{
			name:   "ascending diff greater than 3",
			levels: []int{1, 2, 3, 7},
			want:   false,
		},
		{
			name:   "descending diff greater than 3",
			levels: []int{7, 3, 2, 1},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLevel(tt.levels); tt.want != (got == -1) {
				t.Errorf("checkLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
