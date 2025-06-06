package main

import (
	"github.com/kjameer0/interactive-todo/todo"
	"github.com/rivo/tview"
)

type ui struct {
	app                *tview.Application
	optionsMenu        *tview.List
	output             *tview.Flex
	messageContainer   *tview.TextView
	globalEventManager *globalEventManager
}

func NewUi() *ui {
	uiGlobalEventManager := NewGlobalEventManager()
	ui := &ui{
		app:                tview.NewApplication(),
		optionsMenu:        tview.NewList(),
		output:             tview.NewFlex(),
		messageContainer:   tview.NewTextView().SetText("Message"),
		globalEventManager: uiGlobalEventManager,
	}
	ui.registerEvents()
	return ui
}

func main() {
	ui := NewUi()
	configPath := "./config.json"
	taskStoragePath := "./tasks.json"
	taskManager := todo.NewApp(configPath, taskStoragePath)

	createDefaultOutputMenu(ui, taskManager)
	//TODO: how do i refactor this with global event manager
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
