// Package device knows how to produce the IDs of the combination of tasks that yields an optimally configured device.
package device

import (
	"fmt"
	"sort"
	"strings"
)

// Task represents a single task to be optimized.
type Task struct {
	ID                  int
	ResourceConsumption int
}

// Configuration represents a device's configuration parameters such as capacity, foreground and background tasks.
type Configuration struct {
	Capacity int
	Tasks    [][]Task
}

// ProduceCombinations produces the IDs for an optimally configured device.
func ProduceCombinations(capacity int, foreground, background []Task) [][]int {
	var result [][]int
	var potentialIDs [][]int
	sortedForeground := sortByResource(foreground)
	sortedBackground := sortByResource(background)

	for _, fTask := range sortedForeground {
		for _, bTask := range sortedBackground {
			consumption := fTask.ResourceConsumption + bTask.ResourceConsumption
			if consumption == capacity {
				result = append(result, []int{fTask.ID, bTask.ID})
			} else if consumption < capacity {
				potentialIDs = [][]int{{fTask.ID, bTask.ID}}
			}
		}
	}

	if len(potentialIDs) > 0 && len(result) == 0 {
		return potentialIDs
	}

	return result
}

func sortByResource(tasks []Task) []Task {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ResourceConsumption < tasks[j].ResourceConsumption
	})
	return tasks
}

// FormatCombinations formats a set of combinations with the needed format.
func FormatCombinations(combinations [][][]int) string {
	builder := strings.Builder{}

	for _, combination := range combinations {
		for i, comb := range combination {
			var c string
			if i == len(combination)-1 {
				c = fmt.Sprintf("(%d, %d)", comb[0], comb[1])
			} else {
				c = fmt.Sprintf("(%d, %d), ", comb[0], comb[1])
			}
			builder.WriteString(c)
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
