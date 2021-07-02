// Package device knows how to produce the IDs of the combination of tasks that yields an optimally configured device.
package device

import (
	"fmt"
	"sort"
	"strings"
)

// Task represents a single task to be optimized.
type Task struct {
	ID                 int
	ResourceConsuption int
}

// ProduceCombinations produces the IDS for an optimally configured device.
func ProduceCombinations(capacity int, foreground, background []Task) [][]int {
	var result [][]int
	var potentialIDs [][]int
	sortedForeground := sortByResource(foreground)
	sortedBackground := sortByResource(background)

	for _, fTask := range sortedForeground {
		for _, bTask := range sortedBackground {
			consuption := fTask.ResourceConsuption + bTask.ResourceConsuption
			if consuption == capacity {
				result = append(result, []int{fTask.ID, bTask.ID})
			} else if consuption < capacity {
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
		return tasks[i].ResourceConsuption < tasks[j].ResourceConsuption
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
				c = fmt.Sprintf("(%d, %d) ", comb[0], comb[1])
			} else {
				c = fmt.Sprintf("(%d, %d), ", comb[0], comb[1])
			}
			builder.WriteString(c)
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
