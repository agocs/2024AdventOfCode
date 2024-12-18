package main

import "testing"

func Test_findArguments(t *testing.T) {
	tests := []struct {
		name  string
		args  string
		want  int
		want1 int
	}{
		{
			"mul(1,2)",
			"mul(1,2)",
			1,
			2,
		},
		{
			"mul(123,2)",
			"mul(123,2)",
			123,
			2,
		},
		{
			"mul(1,234)",
			"mul(1,234)",
			1,
			234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findArguments(tt.args)
			if got != tt.want {
				t.Errorf("findArguments() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findArguments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
