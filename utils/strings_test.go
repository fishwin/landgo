package utils

import (
	"testing"
)

func TestRandomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestRandomString",
			args: args{n: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandomString(tt.args.n); len(got) <= 0 {
				t.Errorf("got string is empty")
			} else {
				t.Logf("got: %s", got)
			}
		})
	}
}

func TestSliceContainsString(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name:"exist",
			args:args{
				slice: []string{"abc", "cba", "nba", "ppt", "world"},
				item:  "world",
			},
			want: true,
		},
		{
			name:"not exist",
			args:args{
				slice: []string{"abc", "cba", "nba", "ppt", "world"},
				item:  "word",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceContainsString(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("SliceContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}