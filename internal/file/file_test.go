package file_test

import (
	"reflect"
	"testing"

	"github.com/gustavocd/gotask/internal/device"
	"github.com/gustavocd/gotask/internal/file"
)

var tasksTests = []struct {
	name       string
	cap        int
	foreground []device.Task
	background []device.Task
}{
	{
		"test 1",
		10,
		[]device.Task{
			{ID: 1, ResourceConsuption: 5},
			{ID: 2, ResourceConsuption: 7},
			{ID: 3, ResourceConsuption: 10},
			{ID: 4, ResourceConsuption: 3},
		},
		[]device.Task{
			{ID: 1, ResourceConsuption: 5},
			{ID: 2, ResourceConsuption: 4},
			{ID: 3, ResourceConsuption: 3},
			{ID: 4, ResourceConsuption: 2},
		},
	},
	{
		"test 2",
		20,
		[]device.Task{
			{ID: 1, ResourceConsuption: 9},
			{ID: 2, ResourceConsuption: 15},
			{ID: 3, ResourceConsuption: 8},
		},
		[]device.Task{
			{ID: 1, ResourceConsuption: 11},
			{ID: 2, ResourceConsuption: 8},
			{ID: 3, ResourceConsuption: 12},
		},
	},
	{
		"test 3",
		20,
		[]device.Task{
			{ID: 1, ResourceConsuption: 7},
			{ID: 2, ResourceConsuption: 14},
			{ID: 3, ResourceConsuption: 8},
		},
		[]device.Task{
			{ID: 1, ResourceConsuption: 14},
			{ID: 2, ResourceConsuption: 5},
			{ID: 3, ResourceConsuption: 10},
		},
	},
}

func TestReadDeviceInputs(t *testing.T) {
	inputs, err := file.ReadDeviceInputs("../../challenge.in")
	if err != nil {
		t.Fatal(err)
	}

	for i, tt := range tasksTests {
		input := inputs[i]
		fTasks := input.Tasks[0]
		bTasks := input.Tasks[1]
		t.Run(tt.name, func(t *testing.T) {
			got := input.Capacity
			if got != tt.cap {
				t.Errorf("got %v, want %v", got, tt.cap)
			}

			if !reflect.DeepEqual(fTasks, tt.foreground) {
				t.Errorf("got %v, want %v", fTasks, tt.foreground)
			}

			if !reflect.DeepEqual(bTasks, tt.background) {
				t.Errorf("got %v, want %v", bTasks, tt.background)
			}
		})
	}
}
