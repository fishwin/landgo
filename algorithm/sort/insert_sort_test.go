package sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test selection sort",
			args: args{
				src: []int{
					6,
					4,
					3,
					5,
					7,
					7,
					1,
					0,
					2,
					9,
				},
			},
			want: []int{
				0,
				1,
				2,
				3,
				4,
				5,
				6,
				7,
				7,
				9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			} else {
				t.Logf("InsertSort() = %v", got)
			}
		})
	}
}