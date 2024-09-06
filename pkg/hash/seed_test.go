package hash

import (
	"testing"
)

func TestSeed(t *testing.T) {
	type args struct {
		input any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestSeed",
			args: args{
				input: "test",
			},
			want: 0,
		},
		{
			name: "TestSeed2",
			args: args{
				input: struct{}{},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Seed(tt.args.input); got > 2147483647 || got < -2147483648 {
				t.Errorf("Seed() = %v", got)
			}
		})
	}
}
