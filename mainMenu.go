package main

import (
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

// every component declares its starting state
func initializeMainMenu(ui *ui) {
	// clear elements from previous menu
	ui.resetUI()
	// declare whether or not global events should run on this page, should only be false to capture user input in forms/text fields for now.
	ui.setDoEventsRun(true)
}

func navigateToMainMenu(ui *ui, taskManager *todo.App) {
	initializeMainMenu(ui)
	generateMainOptionsMenu(ui, taskManager)
	createListTaskOutputMenu(ui, taskManager, false)
	ui.app.SetFocus(ui.output)
}

func generateMainOptionsMenu(ui *ui, taskManager *todo.App) {
	var handlers []*handler = []*handler{
		newHandler("Add task", '0', func() {
			navigateToAddTaskMenu(ui, taskManager)
		}),
		newHandler("Delete tasks", '1', func() {
			navigateToDeleteMenu(ui, taskManager)
		}),
		newHandler("Update incomplete tasks", '2', func() {
			navigateToUpdateTaskSelectTable(ui, taskManager, false)
		}),
		newHandler("Update any task", '3', func() {
			navigateToUpdateTaskSelectTable(ui, taskManager, true)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}

func createListTaskOutputMenu(ui *ui, taskManager *todo.App, showComplete bool) *tview.Table {
	tasks := taskManager.ListInsertionOrder(showComplete, false)
	table := generateListTaskOutputTable(ui, taskManager, tasks)
	table.SetFixed(1, 4)
	table.SetEvaluateAllRows(true)

	shortcutKeys := createShortCutKeys(table, len(tasks))
	for idx, key := range shortcutKeys {
		row := table.GetCell(idx+1, 0)
		row.SetText(string(key))
	}
	ui.output.AddItem(table, 0, 2, true)
	return table
}
