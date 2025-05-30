package main

import (
	"fmt"
	"log"

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
	labelWithShortcut := fmt.Sprintf("%c) %s", shortcut, label)
	return &handler{Label: labelWithShortcut, Shortcut: shortcut, Action: action}
}

// populate a list with handler items
func createOptions(ui *ui, handlers []*handler) *tview.List {
	list := tview.NewList()
	var zeroValueRune rune
	for _, handler := range handlers {
		list.AddItem(handler.Label, "", zeroValueRune, nil)
		err := ui.addGlobalEvent(handler.Shortcut, handler.Action)
		if err != nil {
			log.Fatal(err)
		}
	}
	return list
}

func updateOptions(ui *ui, handlers []*handler, list *tview.List) *tview.List {
	list.Clear()
	var zeroValueRune rune
	for _, handler := range handlers {
		list.AddItem(handler.Label, "", zeroValueRune, nil)
		err := ui.addGlobalEvent(handler.Shortcut, handler.Action)
		if err != nil {
			log.Fatal(err)
		}
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

//item from options is selected
//new ui is generated
//focus is moved

// can i at least generate the main menu, allow the user to press 0, and render the list of individual tasks
func generateMainOptionsMenu(ui *ui, taskManager *todo.App) {
	ui.optionsMenu.Clear()
	var handlers []*handler = []*handler{
		newHandler("Add task", '0', func() {}),
		newHandler("Delete tasks", '1', func() {}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}
func createDefaultOutputMenu(ui *ui, taskManager *todo.App) {
	ui.output.Clear()
	tasks := taskManager.ListInsertionOrder(false, false)
	table := createTaskTable(ui, taskManager, tasks)
	ui.output.AddItem(table, 0, 1, false)
}
func createListTaskOutputMenu(ui *ui, taskManager *todo.App) {
	ui.output.Clear()
	tasks := taskManager.ListInsertionOrder(false, false)
	table := generateListTaskOutputMenu(ui, taskManager, tasks)
	shortcutKeys := createShortCutKeys(table, len(tasks))
	for idx, key := range shortcutKeys {
		row := table.GetCell(idx+1, 0)
		row.SetText(string(key))
	}
	ui.output.AddItem(table, 0, 2, false)
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
