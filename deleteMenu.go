package main

import (
	"github.com/kjameer0/interactive-todo/todo"
)

func initializeDeleteMenu(ui *ui) {
	ui.resetUI()
	ui.setDoEventsRun(true)
}

func navigateToDeleteMenu(ui *ui, taskManager *todo.App) {
	initializeDeleteMenu(ui)
	generateDeleteOptionsMenu(ui, taskManager)
	createListTaskOutputMenu(ui, taskManager)
	ui.app.SetFocus(ui.output)
}

func generateDeleteOptionsMenu(ui *ui, taskManager *todo.App) {
	var handlers []*handler = []*handler{
		newHandler("Return to Main menu", rune(27), func() {
			navigateToMainMenu(ui, taskManager)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}

//we ha
