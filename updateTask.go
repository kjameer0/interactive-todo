package main

import (
	"time"

	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

func initializeUpdateTaskMenu(ui *ui) {
	ui.resetUI()
	ui.setDoEventsRun(false)
}

func navigateToUpdateTaskMenu(ui *ui, taskManager *todo.App, t *todo.Task) {
	initializeUpdateTaskMenu(ui)
	// generate options
	// open select menu
	table := createListTaskOutputMenu(ui, taskManager)
	var selectedTask *todo.Task
	// extract choice
	// open form with prefilled values for task
}

// form to add a new task
// form needs to be closed on any event that navigates away
func createUpdateTaskOutputFormMenu(ui *ui, taskManager *todo.App, t *todo.Task) *tview.Form {
	taskName := ""

	form := tview.NewForm()
	form.SetTitle("Add a new task(Press Esc to cancel)")
	form.SetBorder(true)
	form.AddInputField("Task name", t.Name, 25, nil, func(text string) {
		taskName = text
	})
	form.AddButton("Submit", func() {
		if len(taskName) == 0 {
			//TODO: add the message handler functionality in order to display error message
			taskName = "invalid"
		}
		taskManager.AddTask(taskName, time.Now())
		navigateToMainMenu(ui, taskManager)
	})
	form.SetCancelFunc(closeForm(ui, taskManager))
	return form
}

func createUpdateTaskOutputSelectMenu(ui *ui, taskManager *todo.App) {

}

func generateUpdateTaskOptionsMenu(ui *ui, taskManager *todo.App) {
	var handlers []*handler = []*handler{
		newHandler("Return to Main menu", rune(0), func() {
			navigateToMainMenu(ui, taskManager)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}
