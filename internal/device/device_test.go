package device_test

import (
	"reflect"
	"testing"

	"github.com/gustavocd/gotask/internal/device"
)

var tasksTests = []struct {
	name       string
	cap        int
	foreground []device.Task
	background []device.Task
	want       [][]int
}{
	{
		"test 1",
		7,
		[]device.Task{{1, 6}, {2, 2}, {3, 4}},
		[]device.Task{{1, 2}},
		[][]int{{3, 1}},
	},
	{
		"test 2",
		10,
		[]device.Task{{1, 5}, {2, 7}, {3, 10}, {4, 3}},
		[]device.Task{{1, 5}, {2, 4}, {3, 3}, {4, 2}},
		[][]int{{1, 1}, {2, 3}},
	},
	{
		"test 3",
		20,
		[]device.Task{{1, 9}, {2, 15}, {3, 8}},
		[]device.Task{{1, 11}, {2, 8}, {3, 12}},
		[][]int{{3, 3}, {1, 1}},
	},
	{
		"test 4",
		20,
		[]device.Task{{1, 7}, {2, 14}, {3, 8}},
		[]device.Task{{1, 14}, {2, 5}, {3, 10}},
		[][]int{{2, 2}},
	},
}

func TestProduceCombinations(t *testing.T) {
	for _, tt := range tasksTests {
		t.Run(tt.name, func(t *testing.T) {
			got := device.ProduceCombinations(tt.cap, tt.foreground, tt.background)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
