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

func Test_checkLevelsWithRetry(t *testing.T) {

	tests := []struct {
		name    string
		levels  []int
		wantErr bool
	}{
		{
			name:    "ascending",
			levels:  []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "descending",
			levels:  []int{3, 2, 1},
			wantErr: false,
		},
		{
			name:    "stops ascending",
			levels:  []int{1, 2, 3, 3},
			wantErr: false,
		},
		{
			name:    "stops ascending but unrecoverable",
			levels:  []int{1, 2, 3, 3, 3},
			wantErr: true,
		},
		{
			name:    "example 1",
			levels:  []int{7, 6, 4, 2, 1},
			wantErr: false,
		},
		{
			name:    "example 2",
			levels:  []int{1, 2, 7, 8, 9},
			wantErr: true,
		},
		{
			name:    "example 3",
			levels:  []int{9, 7, 6, 2, 1},
			wantErr: true,
		},
		{
			name:    "example 4",
			levels:  []int{1, 3, 2, 4, 5},
			wantErr: false,
		},
		{
			name:    "example 5",
			levels:  []int{8, 6, 4, 4, 1},
			wantErr: false,
		},
		{
			name:    "example 6",
			levels:  []int{1, 3, 6, 7, 9},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkLevelsWithRetry(tt.levels); (err != nil) != tt.wantErr {
				t.Errorf("checkLevelsWithRetry() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_rebuildLevel(t *testing.T) {
	source := []int{1, 2, 3}
	dest := make([]int, 2)
	rebuildLevel(&source, &dest, 1)
	if dest[0] != 1 || dest[1] != 3 {
		t.Errorf("rebuildLevel() = %v", dest)
	}

	rebuildLevel(&source, &dest, 0)
	if dest[0] != 2 || dest[1] != 3 {
		t.Errorf("rebuildLevel() = %v", dest)
	}

	rebuildLevel(&source, &dest, 2)
	if dest[0] != 1 || dest[1] != 2 {
		t.Errorf("rebuildLevel() = %v", dest)
	}
}
