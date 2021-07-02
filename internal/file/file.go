// Package file knows how to read an input file and parse it to a device configuration.
package file

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gustavocd/gotask/internal/device"
)

// DeviceConfiguration represets a device's configuration parameters such as capacity, foreground and background tasks.
type DeviceConfiguration struct {
	Capacity int
	Tasks    [][]device.Task
}

// ReadDeviceInputs reads devices' inputs from a file and transform them into DeviceConfiguration struct.
func ReadDeviceInputs(name string) ([]DeviceConfiguration, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("could not open file %v", err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var content []string

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	defer file.Close()

	var result [][]string
	re := regexp.MustCompile("[0-9]+")
	for i := range content {
		if i%3 == 0 {
			fTask := re.FindAllString(content[i+1], -1)
			bTask := re.FindAllString(content[i+2], -1)
			result = append(result, []string{content[i], strings.Join(fTask, " "), strings.Join(bTask, " ")})
		}
	}

	var dc []DeviceConfiguration
	for i := range result {
		tasks, err := generateDeviceConfiguration(result[i][1], result[i][2])
		if err != nil {
			return nil, err
		}

		capacity, err := strconv.Atoi(result[i][0])
		if err != nil {
			return nil, err
		}

		dc = append(dc, DeviceConfiguration{capacity, tasks})
	}

	return dc, nil
}

func generateDeviceConfiguration(foreground, background string) ([][]device.Task, error) {
	var tasks [][]device.Task

	fTasks, err := generateTasks(foreground)
	if err != nil {
		return nil, err
	}

	bTasks, err := generateTasks(background)
	if err != nil {
		return nil, err
	}

	tasks = append(tasks, fTasks, bTasks)

	return tasks, nil
}

func generateTasks(tasksStr string) ([]device.Task, error) {
	var task device.Task
	var tasks []device.Task

	combinations := strings.Split(tasksStr, " ")
	var i int
	for i < len(combinations)-1 {
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
		i++
	}

	return tasks, nil
}
