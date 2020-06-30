package leetcode

import "testing"

func Test_findString(t *testing.T) {
	type args struct {
		words []string
		s     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test 1",
			args:args{
				words: []string{"at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""},
				s:     "ball",
			},
			want: 4,
		},
		{
			name:"test 2",
			args:args{
				words: []string{"ball", "", "", "", "", "", "", "", "", "","", "", ""},
				s:     "ball",
			},
			want: 0,
		},
		{
			name:"test 3",
			args:args{
				words: []string{"", "", "", "", "", "", "", "", "", "","", "", "ball"},
				s:     "ball",
			},
			want: 12,
		},
		{
			name:"test 4",
			args:args{
				words: []string{"", "", "", "", "", "", "", "", "", "","", "", ""},
				s:     "ball",
			},
			want: -1,
		},
		{
			name:"test 5",
			args:args{
				words: []string{"at", "ba", "cba", "dbb", "efrt", "ews", "fds", "gtre", "hIJDS", "iOU","jkkds", "kiwl", "lmdsk"},
				s:     "efrt",
			},
			want: 4,
		},
		{
			name:"test 6",
			args:args{
				words: []string{"at", "ba", "cba", "dbb", "efrt", "ews", "fds", "gtre", "hIJDS", "iOU","jkkds", "kiwl", "lmdsk"},
				s:     "dddd",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindString(tt.args.words, tt.args.s); got != tt.want {
				t.Errorf("FindString() = %v, want %v", got, tt.want)
			}
		})
	}
}