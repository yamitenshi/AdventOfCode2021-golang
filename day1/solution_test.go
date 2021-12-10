package day1

import (
	"reflect"
	"testing"
)

func TestSonarSweep(t *testing.T) {
	type args struct {
		readings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Empty list_of_ints case", args{[]int{}}, 0},
		{
			"Example case from description",
			args{[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SonarSweep(tt.args.readings); got != tt.want {
				t.Errorf("SonarSweep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlidingSonarSweep(t *testing.T) {
	type args struct {
		readings []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Empty list_of_ints case", args{[]int{}}, 0},
		{
			"Example case from description",
			args{[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			}},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SlidingSonarSweep(tt.args.readings); got != tt.want {
				t.Errorf("SlidingSonarSweep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_emptySlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		want slidingWindow
	}{
		{"Initialisation test", slidingWindow{-1, -1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := emptySlidingWindow(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("emptySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slidingWindow_isFull(t *testing.T) {
	tests := []struct {
		name   string
		window slidingWindow
		want   bool
	}{
		{"Empty sliding window", emptySlidingWindow(), false},
		{"Partially filled sliding window", slidingWindow{-1, 1, 2}, false},
		{"Filled sliding window", slidingWindow{1, 2, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.window.isFull(); got != tt.want {
				t.Errorf("isFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slidingWindow_push(t *testing.T) {
	type args struct {
		measurement int
	}
	tests := []struct {
		name   string
		window slidingWindow
		args   args
		want   slidingWindow
	}{
		{
			"Push onto empty sliding window",
			emptySlidingWindow(),
			args{1},
			slidingWindow{-1, -1, 1},
		},
		{
			"Push onto full sliding window",
			slidingWindow{1, 2, 3},
			args{5},
			slidingWindow{2, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.window.push(tt.args.measurement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_slidingWindow_sum(t *testing.T) {
	tests := []struct {
		name   string
		window slidingWindow
		want   int
	}{
		{"Empty sliding window", emptySlidingWindow(), 0},
		{"Partially filled sliding window", slidingWindow{-1, 1, 2}, 3},
		{"Full sliding window", slidingWindow{1, 2, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.window.sum(); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
