package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

func initializeDeleteMenu(ui *ui) {
	ui.resetUI()
	ui.setDoEventsRun(true)
}

func navigateToDeleteMenu(ui *ui, taskManager *todo.App, page int) {
	initializeDeleteMenu(ui)
	generateDeleteOptionsMenu(ui, taskManager, page)
	table := createListTaskOutputMenu(ui, taskManager, false, page)
	registerDeletionEvents(ui, taskManager, table)
	ui.app.SetFocus(ui.output)
}
func registerDeletionEvents(ui *ui, taskManager *todo.App, table *tview.Table) {
	idsToDelete := make([]string, table.GetRowCount())
	markedForDeletionCount := 0
	ui.addGlobalEvent(rune(tcell.KeyEnter), submitDeletionFunc(ui, taskManager, &idsToDelete, &markedForDeletionCount))
	for rowIdx := 1; rowIdx < table.GetRowCount(); rowIdx++ {
		keyNameCell := table.GetCell(rowIdx, 0)
		taskNameCell := table.GetCell(rowIdx, 1)
		idCell := table.GetCell(rowIdx, 3)
		keyRune := rune([]byte(keyNameCell.Text)[0])
		//toggle whether or not a task will be deleted
		ui.addGlobalEvent(keyRune, func() {
			if indexOf(idCell.Text, idsToDelete) == -1 {
				markedForDeletionCount++

				idsToDelete[rowIdx] = idCell.Text

				selectedColor := tcell.ColorRed

				keyNameCell.SetTextColor(selectedColor)
				taskNameCell.SetTextColor(selectedColor)
				idCell.SetTextColor(selectedColor)
			} else {
				markedForDeletionCount--
				idsToDelete[rowIdx] = ""

				deselectColor := tcell.ColorWhite

				keyNameCell.SetTextColor(deselectColor)
				taskNameCell.SetTextColor(deselectColor)
				idCell.SetTextColor(deselectColor)
			}
		})
	}

}

func submitDeletionFunc(ui *ui, taskManager *todo.App, taskIdsToDelete *[]string, numMarkedForDeletion *int) func() {
	return func() {
		//add count to deletion
		modal := tview.NewTextView().
			SetText(fmt.Sprintf("Do you want to delete %d task(s)\n a: Confirm Deletion b: Go back", *numMarkedForDeletion))

		ui.output.Clear().
			AddItem(modal, 0, 100, true)
		ui.clearAllEvents()
		ui.addGlobalEvent('a', func() {
			for _, id := range *taskIdsToDelete {
				if id != "" {
					taskManager.RemoveTask(id)
				}
			}
			navigateToMainMenu(ui, taskManager, 1)
		})

		ui.addGlobalEvent('b', func() {
			navigateToDeleteMenu(ui, taskManager, 1)
		})
	}
}

func generateDeleteOptionsMenu(ui *ui, taskManager *todo.App, page int) {
	var handlers []*handler = []*handler{
		newHandler("Return to Main menu", rune(0), func() {
			navigateToMainMenu(ui, taskManager, 1)
		}),
		newHandler("Next Page", '1', func() {
			navigateToDeleteMenu(ui, taskManager, page+1)
		}),
		newHandler("Previous Page", '2', func() {
			navigateToDeleteMenu(ui, taskManager, page-1)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}

func indexOf[T string](itemToFind T, list []T) int {
	for idx, elem := range list {
		if elem == itemToFind {
			return idx
		}
	}
	return -1
}
