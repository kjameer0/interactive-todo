package main

import (
	"math"
	"strconv"

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

func navigateToMainMenu(ui *ui, taskManager *todo.App, page int) {
	initializeMainMenu(ui)
	generateMainOptionsMenu(ui, taskManager, page)
	createListTaskOutputMenu(ui, taskManager, false, page)
	ui.app.SetFocus(ui.output)
}

func generateMainOptionsMenu(ui *ui, taskManager *todo.App, page int) {
	var handlers []*handler = []*handler{
		newHandler("Add task", '0', func() {
			navigateToAddTaskMenu(ui, taskManager)
		}),
		newHandler("Delete tasks", '1', func() {
			navigateToDeleteMenu(ui, taskManager, 1)
		}),
		newHandler("Update incomplete tasks", '2', func() {
			navigateToUpdateTaskSelectTable(ui, taskManager, false, 1)
		}),
		newHandler("Update any task", '3', func() {
			navigateToUpdateTaskSelectTable(ui, taskManager, true, 1)
		}),
		newHandler("Next Page", '4', func() {
			navigateToMainMenu(ui, taskManager, page+1)
		}),
		newHandler("Previous Page", '5', func() {
			navigateToMainMenu(ui, taskManager, page-1)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}

func createListTaskOutputMenu(ui *ui, taskManager *todo.App, showComplete bool, page int) *tview.Table {
	tasks := taskManager.ListInsertionOrder(showComplete, false)

	pageLength := 10
	pageLimit := int(math.Ceil(float64(len(tasks)) / float64(pageLength)))
	page = calculatePage(page, pageLimit)

	pageStart := ((page) * pageLength) - pageLength
	pageEnd := int(math.Min(float64(pageStart+pageLength), float64(len(tasks))))

	AppendToFile(strconv.Itoa(pageStart) + " " + strconv.Itoa(pageEnd) + "\n")
	table := generateListTaskOutputTable(ui, taskManager, tasks[pageStart:pageEnd])
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

func calculatePage(page int, pageLimit int) int {
	if page == 0 {
		return pageLimit
	}
	if page > pageLimit {
		return page%pageLimit + 1
	}
	if page < 0 {
		page = page % pageLimit
		return pageLimit + page
	}
	return page
}
