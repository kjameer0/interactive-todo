package todo

import (
	"fmt"
	"log"
	"time"

	"github.com/aidarkhanov/nanoid"
)

type App struct {
	Tasks          map[string]*Task `json:"tasks"`
	InsertionOrder []string         `json:"insertionOrder"`
	saveLocation   string
	configPath     string
	Config         *Config
}

func NewApp(configPath string, saveLocation string) *App {
	tasks := make(map[string]*Task, 100)
	a := &App{Tasks: tasks, configPath: configPath, saveLocation: saveLocation}

	insertionOrder, taskMap := readTasksFromFile(saveLocation)
	a.InsertionOrder = insertionOrder
	a.Tasks = taskMap

	c, err := a.LoadConfig()
	if err != nil {
		log.Fatal("config path not provided, ended application")
	}
	a.Config = c
	return a
}

func NewTask(name string, beginDate time.Time) *Task {
	if name == "" {
		log.Fatal("a task must have a name")
	}
	var taskId string
	taskId, err := nanoid.Generate(nanoid.DefaultAlphabet, 20)
	if err != nil {
		log.Fatal(err)
	}
	t := &Task{Id: taskId, Name: name, BeginDate: beginDate}
	return t
}
func (a *App) GetTaskById(id string) (*Task, error) {
	if t, ok := a.Tasks[id]; ok {
		return t, nil
	}
	return nil, fmt.Errorf("no task witht the id %s", id)
}
func (a *App) ListInsertionOrder(showComplete bool, showFutureTasks bool) []*Task {
	tasks := make([]*Task, 0, len(a.InsertionOrder))
	for _, t := range a.InsertionOrder {
		curTask := a.Tasks[t]
		if !showComplete && curTask.IsComplete() {
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
type Task struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	CompletionDate time.Time  `json:"completionDate"`
	BeginDate      time.Time  `json:"beginDate"`
	Subtasks       []*subtask `json:"subtasks"`
}
type subtask struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func (t *Task) IsComplete() bool {
	return !t.CompletionDate.IsZero()
}

func (t *Task) String() string {
	var completed string
	if !t.IsComplete() {
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

func (a *App) AddTask(taskText string, beginTime time.Time) {
	addedTask := NewTask(taskText, beginTime)
	a.Tasks[addedTask.Id] = addedTask
	a.InsertionOrder = append(a.InsertionOrder, addedTask.Id)
	SaveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
}

func (a *App) RemoveTask(taskId string) bool {
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
	SaveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
	return true
}

func (a *App) RemoveAllTasks() {
	a.InsertionOrder = []string{}
	clear(a.Tasks)
	SaveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
}

func (a *App) ListTasks() []string {
	tasks := []string{}
	for _, taskId := range a.InsertionOrder {
		if taskId == "" {
			continue
		}
		curTask := a.Tasks[taskId]
		//show a task if it not complete or if show complete and task
		if !a.Config.ShowComplete && curTask.IsComplete() {
			continue
		}
		if time.Now().Compare(curTask.BeginDate) == -1 {
			continue
		}
		var completed string
		if !curTask.IsComplete() {
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

func (a *App) UpdateTask(t *Task) {
	var zeroTime time.Time
	if !t.IsComplete() {
		t.CompletionDate = time.Now()
	} else {
		t.CompletionDate = zeroTime
	}
	SaveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
}

func (a *App) ToggleTaskCompletion(t *Task) *Task {
	var zeroTime time.Time
	if !t.IsComplete() {
		t.CompletionDate = time.Now()
	} else {
		t.CompletionDate = zeroTime
	}
	return t
}

func (a *App) UpdateTaskInfo(t *Task) {
	SaveToFile(a.saveLocation, a.InsertionOrder, a.Tasks)
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

func generateDaysList(t time.Time, numDays int) []string {
	dateList := []string{}
	for i := 0; i < numDays; i++ {
		addedDay := addDayToDate(t, i)
		dateList = append(dateList, monthDayYear(addedDay))
	}
	return dateList
}

type SaveData struct {
	Tasks          map[string]*Task `json:"tasks"`
	InsertionOrder []string         `json:"insertionOrder"`
}
