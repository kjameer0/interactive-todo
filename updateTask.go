package main

import (
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

func initializeUpdateTaskMenu(ui *ui) {
	ui.resetUI()
	ui.setDoEventsRun(true)
}
func initializeUpdateTaskForm(ui *ui) {
	ui.resetUI()
	ui.setDoEventsRun(false)
}

// form needs to be closed on any event that navigates away
func createUpdateTaskOutputFormMenu(ui *ui, taskManager *todo.App, t *todo.Task) *tview.Form {
	taskName := t.Name

	form := tview.NewForm()
	form.SetTitle("Update new task(Press Esc to cancel)")
	form.SetBorder(true)
	form.AddInputField("Task name", t.Name, 25, nil, func(text string) {
		taskName = text
	})

	form.AddCheckbox("Completed", t.IsComplete(), func(checked bool) {
		taskManager.ToggleTaskCompletion(t)
	})
	form.AddButton("Submit", func() {
		if len(taskName) == 0 {
			ui.messageChannel <- "Error: Task name cannot be empty"
			navigateToUpdateTaskForm(ui, taskManager, t)
			return
		}
		//update all task fields here
		t.Name = taskName

		taskManager.UpdateTaskInfo(t)
		ui.messageChannel <- "Task successfully updated"
		navigateToMainMenu(ui, taskManager)
	})
	form.SetCancelFunc(closeForm(ui, taskManager))
	return form
}

func createUpdateTaskOutputSelectMenu(ui *ui, taskManager *todo.App, showComplete bool) {
	table := createListTaskOutputMenu(ui, taskManager, showComplete)
	registerSelectUpdateEvents(ui, taskManager, table)
}

func generateUpdateTaskOptionsMenu(ui *ui, taskManager *todo.App) {
	var handlers []*handler = []*handler{
		newHandler("Return to Main menu", rune(0), func() {
			navigateToMainMenu(ui, taskManager)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}

func registerSelectUpdateEvents(ui *ui, taskManager *todo.App, table *tview.Table) {
	for rowIdx := 1; rowIdx < table.GetRowCount(); rowIdx++ {
		keyNameCell := table.GetCell(rowIdx, 0)
		idCell := table.GetCell(rowIdx, 3)
		keyRune := rune([]byte(keyNameCell.Text)[0])

		ui.addGlobalEvent(keyRune, func() {
			//navigate to update menu with task
			t, err := taskManager.GetTaskById(idCell.Text)
			if err != nil {
				ui.messageChannel <- err.Error()
			}
			navigateToUpdateTaskForm(ui, taskManager, t)
		})
	}

}

func navigateToUpdateTaskForm(ui *ui, taskManager *todo.App, t *todo.Task) {
	initializeUpdateTaskForm(ui)
	form := createUpdateTaskOutputFormMenu(ui, taskManager, t)
	ui.output.AddItem(form, 0, 2, true)
	generateUpdateTaskOptionsMenu(ui, taskManager)
	//focus the form so user can immediately start typing
	ui.app.SetFocus(ui.output)
}

func navigateToUpdateTaskSelectTable(ui *ui, taskManager *todo.App, showComplete bool) {
	initializeUpdateTaskMenu(ui)
	generateUpdateTaskOptionsMenu(ui, taskManager)
	createUpdateTaskOutputSelectMenu(ui, taskManager, showComplete)
	ui.app.SetFocus(ui.output)
}
