package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

func initializeDeleteMenu(ui *ui) {
	ui.resetUI()
	ui.setDoEventsRun(true)
}

func navigateToDeleteMenu(ui *ui, taskManager *todo.App) {
	initializeDeleteMenu(ui)
	generateDeleteOptionsMenu(ui, taskManager)
	table := createListTaskOutputMenu(ui, taskManager)
	registerDeletionEvents(ui, taskManager, table)
	ui.app.SetFocus(ui.output)
}
func registerDeletionEvents(ui *ui, taskManager *todo.App, table *tview.Table) {
	//TODO: figure out how to not run into a memory problem by spreading out values in the array
	idsToDelete := make([]string, table.GetRowCount())
	ui.addGlobalEvent(rune(tcell.KeyEnter), submitDeletionFunc(ui, taskManager, &idsToDelete))
	for rowIdx := 1; rowIdx < table.GetRowCount(); rowIdx++ {
		keyNameCell := table.GetCell(rowIdx, 0)
		taskNameCell := table.GetCell(rowIdx, 1)
		idCell := table.GetCell(rowIdx, 3)
		keyRune := rune([]byte(keyNameCell.Text)[0])
		ui.addGlobalEvent(keyRune, func() {
			if indexOf(idCell.Text, idsToDelete) == -1 {
				idsToDelete[rowIdx] = idCell.Text

				selectedColor := tcell.ColorRed

				keyNameCell.SetTextColor(selectedColor)
				taskNameCell.SetTextColor(selectedColor)
				idCell.SetTextColor(selectedColor)
			} else {
				idsToDelete[rowIdx] = ""

				keyNameCell.SetTextColor(tcell.ColorWhite)
				taskNameCell.SetTextColor(tcell.ColorWhite)
				idCell.SetTextColor(tcell.ColorWhite)
			}
		})
	}

}

func submitDeletionFunc(ui *ui, taskManager *todo.App, taskIdsToDelete *[]string) func() {
	return func() {
		modal := tview.NewModal().
			SetText("Do you want to quit the application?").
			AddButtons([]string{"Quit", "Cancel"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Quit" {
					ui.app.Stop()
				}
			})
		ui.output.AddItem(modal, 0, 1, true)
	}
}

func generateDeleteOptionsMenu(ui *ui, taskManager *todo.App) {
	var handlers []*handler = []*handler{
		newHandler("Return to Main menu", rune(27), func() {
			navigateToMainMenu(ui, taskManager)
		}),
	}
	updateOptions(ui, handlers, ui.optionsMenu)
}

// create menu
// menu has letters
// for each letter register event to delete task
// grab id from row
// add to array
// when user confirms use array to delete all ids
func indexOf[T string](itemToFind T, list []T) int {
	for idx, elem := range list {
		if elem == itemToFind {
			return idx
		}
	}
	return -1
}

func densifyStringArray(list []string) []string {
	densifiedArray := make([]string, 0, len(list)/2)
	for _, item := range list {
		if item != "" {
			densifiedArray = append(densifiedArray, item)
		}
	}
	return densifiedArray

}
