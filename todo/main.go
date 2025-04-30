package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aidarkhanov/nanoid"
)

type app struct {
	Tasks          map[string]*task `json:"tasks"`
	InsertionOrder []string         `json:"insertionOrder"`
	saveLocation   string
	configPath     string
	config         *Config
}

func newApp() *app {
	tasks := make(map[string]*task, 100)
	return &app{Tasks: tasks}
}
func newTask(name string, beginDate time.Time) *task {
	if name == "" {
		log.Fatal("a task must have a name")
	}
	var taskId string
	taskId, err := nanoid.Generate(nanoid.DefaultAlphabet, 20)
	if err != nil {
		log.Fatal(err)
	}
	t := &task{Id: taskId, Name: name, BeginDate: beginDate}
	return t
}
func (a *app) listInsertionOrder(showComplete bool, showFutureTasks bool) []*task {
	tasks := make([]*task, 0, len(a.InsertionOrder))
	for _, t := range a.InsertionOrder {
		curTask := a.Tasks[t]
		if !showComplete && curTask.isComplete() {
			continue
		}
		if time.Now().Compare(curTask.BeginDate) == -1 && showFutureTasks {
			continue
		}
		tasks = append(tasks, curTask)
	}
	return tasks
}

// if completion date is zero value then the task is incomplete
type task struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	CompletionDate time.Time `json:"completionDate"`
	BeginDate      time.Time `json:"beginDate"`
}

func (t *task) isComplete() bool {
	return !t.CompletionDate.IsZero()
}

func (t *task) String() string {
	var completed string
	if !t.isComplete() {
		completed = "❌"
	} else {
		completed = "✅"
	}
	var completionDate string
	if t.CompletionDate.IsZero() {
		completionDate = ""
	} else {
		completionDate = monthDayYear(t.CompletionDate)
	}
	return fmt.Sprintf("%s %s %s", t.Name, completed, completionDate)
}

func addTask(a *app, taskText string, beginTime time.Time) {
	addedTask := newTask(taskText, beginTime)
	a.Tasks[addedTask.Id] = addedTask
	a.InsertionOrder = append(a.InsertionOrder, addedTask.Id)
	saveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
}

func removeTask(a *app, taskId string) bool {
	if _, ok := a.Tasks[taskId]; !ok {
		fmt.Println("hi")
		return false
	}
	delete(a.Tasks, taskId)
	//remove deleted id from insertion order
	filteredInsertionOrder := []string{}
	for _, id := range a.InsertionOrder {
		if id == taskId {
			continue
		}
		filteredInsertionOrder = append(filteredInsertionOrder, id)
	}
	a.InsertionOrder = filteredInsertionOrder
	saveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
	return true
}

func removeAllTasks(a *app) {
	a.InsertionOrder = []string{}
	clear(a.Tasks)
	saveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
}

func listTasks(a *app) []string {
	tasks := []string{}
	for _, taskId := range a.InsertionOrder {
		if taskId == "" {
			continue
		}
		curTask := a.Tasks[taskId]
		//show a task if it not complete or if show complete and task
		if !a.config.ShowComplete && curTask.isComplete() {
			continue
		}
		if time.Now().Compare(curTask.BeginDate) == -1 {
			continue
		}
		var completed string
		if !curTask.isComplete() {
			completed = "❌"
		} else {
			completed = "✅"
		}
		t := monthDayYear(curTask.CompletionDate)
		if curTask.CompletionDate.IsZero() {
			t = ""
		}
		tasks = append(tasks, fmt.Sprintf("%s %s %s", curTask.Name, completed, t))
	}
	return tasks
}

func updateTask(a *app, t *task) {
	var zeroTime time.Time
	if !t.isComplete() {
		t.CompletionDate = time.Now()
	} else {
		t.CompletionDate = zeroTime
	}
	saveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
}

func monthDayYear(t time.Time) string {
	y, m, d := t.Date()
	return fmt.Sprintf("%v %v %v %v", t.Weekday().String(), m.String()[:3], d, y)
}

func addDayToDate(t time.Time, days int) time.Time {
	if days <= 0 {
		days = 0
	}
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	aheadTime := midnight.Add(time.Hour * 24 * time.Duration(days))
	return aheadTime
}

type saveData struct {
	Tasks          map[string]*task `json:"tasks"`
	InsertionOrder []string         `json:"insertionOrder"`
}

func saveToFile(saveLocation string, insertionOrder []string, tasks map[string]*task) {
	s := saveData{}
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

func readTasksFromFile(saveLocation string, insertionOrder []string, tasks map[string]*task) ([]string, map[string]*task) {
	data, err := os.ReadFile(saveLocation)
	s := saveData{}
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Fatal("Task file does not exist")
		}
		log.Fatal("failed to read from save location", err)
	}
	json.Unmarshal(data, &s)
	finalInsertionOrder := s.InsertionOrder
	var taskMap map[string]*task
	if len(s.Tasks) > 0 {
		taskMap = s.Tasks
	}
	return finalInsertionOrder, taskMap
}
