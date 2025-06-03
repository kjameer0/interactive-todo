package main

import (
	"time"

	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

func initializeAddTaskMenu(ui *ui) {
	ui.resetUI()
	ui.setDoEventsRun(false)
}

func navigateToAddTaskMenu(ui *ui, taskManager *todo.App) {
	initializeAddTaskMenu(ui)

	form := createAddTaskOutputMenu(ui, taskManager)
	ui.output.AddItem(form, 0, 2, true)
	generateAddTaskOptionsMenu(ui, taskManager)
	//focus the form so user can immediately start typing
	ui.app.SetFocus(ui.output)
}

// form to add a new task
// form needs to be closed on any event that navigates away
func createAddTaskOutputMenu(ui *ui, taskManager *todo.App) *tview.Form {
	taskName := ""

	form := tview.NewForm()
	form.SetTitle("Add a new task(Press Esc to cancel)")
	form.SetBorder(true)
	form.AddInputField("Task name", "", 25, nil, func(text string) {
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

func generateAddTaskOptionsMenu(ui *ui, taskManager *todo.App) {
	var handlers []*handler = []*handler{
		newHandler("Return to Main menu", rune(0), func() {
			navigateToMainMenu(ui, taskManager)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}
func closeForm(ui *ui, taskManager *todo.App) func() {
	return func() {
		navigateToMainMenu(ui, taskManager)
	}
}
