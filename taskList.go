package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

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

func createListTaskMenu(ui *ui, taskManager *todo.App) {
	ui.optionsMenu.Clear()
	ui.optionsMenu.AddItem("Return to main menu", "", '0', func() {
		//generate main menu
		generateMainOptionsMenu(ui, taskManager)
		createDefaultOutputMenu(ui, taskManager)
	})
}
