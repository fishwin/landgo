package leetcode

import "testing"

func Test_compressString(t *testing.T) {
	type args struct {
		S string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"test 1",
			args:args{S:"aabcccccaaa"},
			want: "a2b1c5a3",
		},
		{
			name:"test 2",
			args:args{S:"abbccd"},
			want: "abbccd",
		},
		{
			name:"test 3",
			args:args{S:"bb"},
			want: "bb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressString(tt.args.S); got != tt.want {
				t.Errorf("CompressString() = %v, want %v", got, tt.want)
			}
		})
	}
}