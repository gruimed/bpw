package main

import (
	"context"
	"testing"
)

func Test_rollonce(t *testing.T) {
	type args struct {
		ctx  context.Context
		load string
	}
	tests := []struct {
		name string
		args args
		min  int
		max  int
	}{
		{
			name: "Test1",
			args: args{
				ctx:  context.Background(),
				load: "C",
			},
			min: 1,
			max: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rollonce(tt.args.ctx, tt.args.load); got > tt.max && got < tt.min {
				t.Errorf("rollonce() = %v, max %v", got, tt.max)
			}
		})
	}
}

func Benchmark_rollonce(b *testing.B) {
	ctx := context.Background()
	for b.Loop() {
		rollonce(ctx, "C")
	}
}
