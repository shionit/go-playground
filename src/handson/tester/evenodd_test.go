package tester

import "testing"

func TestIsEven(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "0 even",
			args: args{n: 0},
			want: true,
		},
		{
			name: "1 odd",
			args: args{n: 1},
			want: false,
		},
		{
			name: "2 even",
			args: args{n: 2},
			want: true,
		},
		{
			name: "3 odd",
			args: args{n: 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := IsEven(tt.args.n); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
