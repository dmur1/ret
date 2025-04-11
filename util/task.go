package util

import (
	"encoding/json"
	"log"
	"os"
	"ret/config"
	"ret/data"
	"ret/theme"
)

func GetCurrentTask() data.Task {
	var task data.Task

	jsonData, err := os.ReadFile(config.TaskFileName)
	if err != nil {
		return task
	}

	err = json.Unmarshal(jsonData, &task)
	if err != nil {
		return task
	}

	return task
}

func SetCurrentTask(task *data.Task) {
	jsonData, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		log.Fatalf("💥 "+theme.ColorRed+"error"+theme.ColorReset+": %v\n", err)
	}

	err = os.WriteFile(config.TaskFileName, jsonData, 0644)
	if err != nil {
		log.Fatalf("💥 "+theme.ColorRed+"error"+theme.ColorReset+": %v\n", err)
	}
}

func GetCurrentTaskName() string {
	task := GetCurrentTask()
	return task.Name
}

func SetCurrentTaskName(name string) {
	task := GetCurrentTask()
	task.Name = name
	SetCurrentTask(&task)
}

func GetCurrentTaskCategory() string {
	task := GetCurrentTask()
	return task.Category
}

func SetCurrentTaskCategory(category string) {
	task := GetCurrentTask()
	task.Category = category
	SetCurrentTask(&task)
}
