package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

type handler struct {
	Label    string
	Shortcut rune
	Action   func()
}

// create a new handler method than can be passed to a tview list
func newHandler(label string, shortcut rune, action func()) *handler {
	return &handler{Label: label, Shortcut: shortcut, Action: action}
}

// populate a list with handler items
func createOptions(ui *ui, handlers []*handler) *tview.List {
	list := tview.NewList()

	for _, hander := range handlers {
		list.AddItem(hander.Label, "", hander.Shortcut, hander.Action)
	}
	return list
}

// function for creating a handler function for when the user wants to list out all of their tasks. should be registered witha a global event manager
func listTaskHandler(ui *ui, taskManager *todo.App) func() {
	return func() {
		//generate new list options(return to main menu)
		createListTaskMenu(ui, taskManager)
		//generate new output menu
		tasks := taskManager.ListInsertionOrder(false, false)
		table := generateListTaskOutputMenu(ui, taskManager, tasks)
		shortcutKeys := createShortCutKeys(table, len(tasks))
		for idx, key := range shortcutKeys {
			row := table.GetCell(idx+1, 0)
			row.SetText(string(key))
		}
		ui.output.Clear().AddItem(table, 0, 2, true)
	}
}

func createListTaskMenu(ui *ui, taskManager *todo.App) {
	ui.optionsMenu.Clear()
	ui.optionsMenu.AddItem("Return to main menu", "", '0', func() {
		//generate main menu
		generateMainOptionsMenu(ui, taskManager)
		createDefaultOutputMenu(ui, taskManager)
	})
}

//item from options is selected
//new ui is generated
//focus is moved

// can i at least generate the main menu, allow the user to press 0, and render the list of individual tasks
func generateMainOptionsMenu(ui *ui, taskManager *todo.App) {
	ui.optionsMenu.Clear()
	var handlers []handler = []handler{
		*newHandler("List tasks", '0', listTaskHandler(ui, taskManager)),
		*newHandler("Add task", '1', func() {}),
		*newHandler("Delete tasks", '2', func() {}),
	}
	for _, handler := range handlers {
		ui.optionsMenu.AddItem(handler.Label, "", handler.Shortcut, handler.Action)
	}
}
func createDefaultOutputMenu(ui *ui, taskManager *todo.App) {
	ui.output.Clear()
	tasks := taskManager.ListInsertionOrder(false, false)
	table := createTaskTable(ui, taskManager, tasks)
	ui.output.AddItem(table, 0, 1, false)
}

func createTaskTable(ui *ui, taskManager *todo.App, taskList []*todo.Task) *tview.Table {
	taskTable := tview.NewTable().SetSelectable(false, false)
	taskTable.SetBorders(true)

	nameCell := tview.NewTableCell("Name")
	statusCell := tview.NewTableCell("Completion Status")
	statusCell.SetAlign(tview.AlignRight)

	taskTable.SetCell(0, 0, nameCell)
	taskTable.SetCell(0, 1, statusCell)

	for rowNum, t := range taskList {
		cell := tview.NewTableCell(t.Name)
		complete := "✅"
		if t.IsComplete() {
			complete = "❌"
		}
		completionCell := tview.NewTableCell(complete).SetAlign(tview.AlignCenter)
		taskTable.SetCell(rowNum+1, 0, cell)
		taskTable.SetCell(rowNum+1, 1, completionCell)
	}
	return taskTable
}

// list of single tasks that can be selected to open single task menu
func generateListTaskOutputMenu(ui *ui, taskManager *todo.App, taskList []*todo.Task) *tview.Table {
	taskTable := tview.NewTable().SetSelectable(false, false)
	taskTable.SetBorders(true)
	keyCell := tview.NewTableCell("Key")
	nameCell := tview.NewTableCell("Name")
	statusCell := tview.NewTableCell("Completion Status")
	statusCell.SetAlign(tview.AlignRight)

	taskTable.SetCell(0, 0, keyCell)
	taskTable.SetCell(0, 1, nameCell)
	taskTable.SetCell(0, 2, statusCell)

	for rowNum, t := range taskList {
		cell := tview.NewTableCell(t.Name)
		complete := "✅"
		if t.IsComplete() {
			complete = "❌"
		}
		keyNameCell := tview.NewTableCell("").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter)
		completionCell := tview.NewTableCell(complete).SetAlign(tview.AlignCenter)
		taskTable.SetCell(rowNum+1, 0, keyNameCell)
		taskTable.SetCell(rowNum+1, 1, cell)
		taskTable.SetCell(rowNum+1, 2, completionCell)
	}
	return taskTable
}

func createShortCutKeys(table *tview.Table, n int) []rune {
	keyList := make([]rune, 0, n)
	r := 'a'
	for i := 0; i < n; i++ {
		keyList = append(keyList, r)
		r += 1
		if r == 'z'+1 {
			r = 'A'
		}
	}
	return keyList
}

