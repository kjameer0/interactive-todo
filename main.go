package main

import (
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

type ui struct {
	app              *tview.Application
	optionsMenu      *tview.List
	output           *tview.Flex
	messageContainer *tview.TextView
}

func main() {
	// TODO: write event core architecture
	ui := &ui{
		app:              tview.NewApplication(),
		optionsMenu:      tview.NewList(),
		output:           tview.NewFlex(),
		messageContainer: tview.NewTextView().SetText("Message"),
	}
	configPath := "./config.json"
	taskStoragePath := "./tasks.json"
	taskManager := todo.NewApp(configPath, taskStoragePath)

	createDefaultOutputMenu(ui, taskManager)

	listTaskOption := newHandler("List Tasks", '0', listTaskHandler(ui, taskManager))
	deleteTaskOption := newHandler("Delete Tasks", '1', listTaskHandler(ui, taskManager))
	mainOptionsMenu := createOptions(ui, []*handler{listTaskOption, deleteTaskOption})
	ui.optionsMenu = mainOptionsMenu

	wrapper := tview.NewFlex().SetDirection(tview.FlexColumnCSS)

	wrapper.AddItem(ui.messageContainer, 3, 1, false)
	wrapper.AddItem(ui.optionsMenu, 0, 2, true)
	layout := tview.NewFlex().
		AddItem(wrapper, 0, 1, false).
		AddItem(ui.output, 0, 3, false)
	if err := ui.app.SetRoot(layout, true).EnableMouse(true).SetFocus(ui.optionsMenu).Run(); err != nil {
		// todo.SaveToFile()
		panic(err)
	}
}

//how do i handle the events? Should I handle events at the application level? technically we already have an event driven architecture, but I want to generalize some of the behavior.
