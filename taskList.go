package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

// list of single tasks that can be selected to open single task menu
func generateListTaskOutputTable(ui *ui, taskManager *todo.App, taskList []*todo.Task) *tview.Table {
	taskTable := tview.NewTable().SetSelectable(false, false)
	taskTable.SetBorders(true)
	keyCell := tview.NewTableCell("Key")
	nameCell := tview.NewTableCell("Name")
	statusCell := tview.NewTableCell("Completed")
	idCell := tview.NewTableCell("ID")
	statusCell.SetAlign(tview.AlignRight)
	taskTable.SetCell(0, 0, keyCell)
	taskTable.SetCell(0, 1, nameCell)
	taskTable.SetCell(0, 2, statusCell)
	taskTable.SetCell(0, 3, idCell)

	for rowNum, t := range taskList {
		cell := tview.NewTableCell(t.Name)
		complete := "❌"
		if t.IsComplete() {
			complete = "✅"
		}
		keyNameCell := tview.NewTableCell("").SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignCenter)
		completionCell := tview.NewTableCell(complete).SetAlign(tview.AlignCenter)
		idTextCell := tview.NewTableCell(t.Id).
			SetAlign(tview.AlignCenter).
			SetMaxWidth(10)
		taskTable.SetCell(rowNum+1, 0, keyNameCell)
		taskTable.SetCell(rowNum+1, 1, cell)
		taskTable.SetCell(rowNum+1, 2, completionCell)
		taskTable.SetCell(rowNum+1, 3, idTextCell)
	}
	return taskTable
}
