// Package file knows how to read an input file and parse it to a device configuration.
package file

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gustavocd/gotask/internal/device"
)

// ReadDeviceInputs reads devices' inputs from a file and transform them into DeviceConfiguration struct.
func ReadDeviceInputs(name string) ([]device.Configuration, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var content []string

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	defer file.Close()

	extractedTasks := extractTasksFromInputs(content)
	dc, err := generateDeviceConfiguration(extractedTasks)
	if err != nil {
		return nil, err
	}

	return dc, nil
}

func extractTasksFromInputs(content []string) [][]string {
	var extractedTasks [][]string
	re := regexp.MustCompile("[0-9]+")
	for i := range content {
		if i%3 == 0 {
			fTask := re.FindAllString(content[i+1], -1)
			bTask := re.FindAllString(content[i+2], -1)
			extractedTasks = append(extractedTasks, []string{content[i], strings.Join(fTask, " "), strings.Join(bTask, " ")})
		}
	}

	return extractedTasks
}

func generateDeviceConfiguration(extractedTasks [][]string) ([]device.Configuration, error) {
	var dc []device.Configuration

	for i := range extractedTasks {
		var tasks [][]device.Task

		fTasks, err := parseTasks(extractedTasks[i][1])
		if err != nil {
			return nil, err
		}

		bTasks, err := parseTasks(extractedTasks[i][2])
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, fTasks, bTasks)

		capacity, err := strconv.Atoi(extractedTasks[i][0])
		if err != nil {
			return nil, err
		}

		dc = append(dc, device.Configuration{Capacity: capacity, Tasks: tasks})
	}

	return dc, nil
}

func parseTasks(tasksStr string) ([]device.Task, error) {
	var task device.Task
	var tasks []device.Task

	combinations := strings.Split(tasksStr, " ")
	for i := range combinations {
		if i%2 == 0 {
			tID, err := strconv.Atoi(combinations[i])
			if err != nil {
				return nil, err
			}
			task.ID = tID
			tRc, err := strconv.Atoi(combinations[i+1])
			if err != nil {
				return nil, err
			}
			task.ResourceConsuption = tRc
			tasks = append(tasks, task)
		}
	}

	return tasks, nil
}
