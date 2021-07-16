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
			{ID: 1, ResourceConsumption: 5},
			{ID: 2, ResourceConsumption: 7},
			{ID: 3, ResourceConsumption: 10},
			{ID: 4, ResourceConsumption: 3},
		},
		[]device.Task{
			{ID: 1, ResourceConsumption: 5},
			{ID: 2, ResourceConsumption: 4},
			{ID: 3, ResourceConsumption: 3},
			{ID: 4, ResourceConsumption: 2},
		},
	},
	{
		"test 2",
		20,
		[]device.Task{
			{ID: 1, ResourceConsumption: 9},
			{ID: 2, ResourceConsumption: 15},
			{ID: 3, ResourceConsumption: 8},
		},
		[]device.Task{
			{ID: 1, ResourceConsumption: 11},
			{ID: 2, ResourceConsumption: 8},
			{ID: 3, ResourceConsumption: 12},
		},
	},
	{
		"test 3",
		20,
		[]device.Task{
			{ID: 1, ResourceConsumption: 7},
			{ID: 2, ResourceConsumption: 14},
			{ID: 3, ResourceConsumption: 8},
		},
		[]device.Task{
			{ID: 1, ResourceConsumption: 14},
			{ID: 2, ResourceConsumption: 5},
			{ID: 3, ResourceConsumption: 10},
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
