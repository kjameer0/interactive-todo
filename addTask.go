package main

import (
	"time"

	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

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
		closeForm(ui, taskManager)()
	})
	form.SetCancelFunc(closeForm(ui, taskManager))
	return form
}

func closeForm(ui *ui, taskManager *todo.App) func() {
	return func() {
		ui.app.SetFocus(ui.optionsMenu)
		generateMainOptionsMenu(ui, taskManager)
		createListTaskOutputMenu(ui, taskManager)
		ui.setDoEventsRun(true)
	}
}
