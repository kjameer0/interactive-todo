package todo

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func SaveToFile(saveLocation string, insertionOrder []string, tasks map[string]*Task) {
	s := SaveData{}
	s.Tasks = tasks
	s.InsertionOrder = insertionOrder
	taskJson, err := json.Marshal(s)
	if err != nil {
		log.Fatal("failed to convert tasks to JSON")
	}
	err = os.WriteFile(saveLocation, taskJson, 0644)
	if err != nil {
		log.Fatal("failed to write to file ", saveLocation)
	}
}

func readTasksFromFile(saveLocation string, insertionOrder []string, tasks map[string]*Task) ([]string, map[string]*Task) {
	data, err := os.ReadFile(saveLocation)
	s := SaveData{}
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("Task file does not exist")
		}
		log.Fatal("failed to read from save location", err)
	}
	json.Unmarshal(data, &s)
	finalInsertionOrder := s.InsertionOrder
	var taskMap map[string]*Task
	if len(s.Tasks) > 0 {
		taskMap = s.Tasks
	}
	return finalInsertionOrder, taskMap
}
