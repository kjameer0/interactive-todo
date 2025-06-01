package main

import "github.com/kjameer0/interactive-todo/todo"

func initializeMainMenu(ui *ui) {
	ui.output.Clear()
	ui.optionsMenu.Clear()
	ui.clearAllEvents()
	ui.setDoEventsRun(true)
}

func navigateToMainMenu(ui *ui, taskManager *todo.App) {
	initializeMainMenu(ui)
	generateMainOptionsMenu(ui, taskManager)
	createListTaskOutputMenu(ui, taskManager)
}

func generateMainOptionsMenu(ui *ui, taskManager *todo.App) {

	var handlers []*handler = []*handler{
		newHandler("Add task", '0', func() {
			ui.setDoEventsRun(false)
			navigateToAddTaskMenu(ui, taskManager)
		}),
		newHandler("Delete tasks", '1', func() {}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}

func createListTaskOutputMenu(ui *ui, taskManager *todo.App) {
	ui.output.Clear()

	tasks := taskManager.ListInsertionOrder(false, false)
	table := generateListTaskOutputTable(ui, taskManager, tasks)
	table.SetFixed(1, 3)

	shortcutKeys := createShortCutKeys(table, len(tasks))
	for idx, key := range shortcutKeys {
		row := table.GetCell(idx+1, 0)
		row.SetText(string(key))
	}
	ui.output.AddItem(table, 0, 2, false)
}
