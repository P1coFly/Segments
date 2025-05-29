package main

import (
	"testing"
)

func TestMinPoints(t *testing.T) {
	tests := []struct {
		name     string
		segments []segment
		want     int
	}{
		{
			name:     "empty",
			segments: []segment{},
			want:     0,
		},
		{
			name:     "single segment",
			segments: []segment{{Start: 1, End: 2}},
			want:     1,
		},
		{
			name: "non-overlapping segments",
			segments: []segment{
				{Start: 1, End: 2},
				{Start: 3, End: 4},
			},
			want: 2,
		},
		{
			name: "overlapping segments",
			segments: []segment{
				{Start: 1, End: 2},
				{Start: 2, End: 3},
			},
			want: 1,
		},
		{
			name: "nested segments",
			segments: []segment{
				{Start: 1, End: 5},
				{Start: 2, End: 3},
				{Start: 4, End: 5},
			},
			want: 2,
		},
		{
			name: "consecutive segments",
			segments: []segment{
				{Start: 1, End: 4},
				{Start: 2, End: 3},
				{Start: 3, End: 5},
			},
			want: 1,
		},
		{
			name: "single point segment",
			segments: []segment{
				{Start: 1, End: 1},
			},
			want: 1,
		},
		{
			name: "reverse order",
			segments: []segment{
				{Start: 3, End: 4},
				{Start: 1, End: 2},
			},
			want: 2,
		},
		{
			name: "complex overlap",
			segments: []segment{
				{Start: 1, End: 3},
				{Start: 2, End: 4},
				{Start: 5, End: 6},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPoints(tt.segments); got != tt.want {
				t.Errorf("minPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
