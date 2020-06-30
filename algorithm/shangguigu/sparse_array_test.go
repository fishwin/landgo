package shangguigu

import (
	"reflect"
	"testing"
)

func TestSparseArrayCompress(t *testing.T) {
	type args struct {
		src [][]int
	}
	tests := []struct {
		name string
		args args
		want [][3]int
	}{
		{
			name: "TestSparseArrayCompress",
			args: args{
				src: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 2, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 2, 3, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				},
			},
			want: [][3]int{
				{10, 10, 4},
				{1, 1, 1},
				{2, 2, 2},
				{3, 2, 2},
				{3, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SparseArrayCompress(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SparseArrayCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSparseArrayDecompress(t *testing.T) {
	type args struct {
		src [][3]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 2",
			args: args{
				src: [][3]int{
					{10, 10, 4},
					{1, 1, 1},
					{2, 2, 2},
					{3, 2, 2},
					{3, 3, 3},
				},
			},
			want: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 2, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 2, 3, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SparseArrayDecompress(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SparseArrayDecompress() = %v, want %v", got, tt.want)
			}
		})
	}
}
