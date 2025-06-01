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

	navigateToMainMenu(ui, taskManager)

	wrapper := tview.NewFlex().SetDirection(tview.FlexColumnCSS)
	wrapper.AddItem(ui.messageContainer, 3, 1, false)
	wrapper.AddItem(ui.optionsMenu, 0, 3, false)
	layout := tview.NewFlex().
		AddItem(wrapper, 0, 2, false).
		AddItem(ui.output, 0, 4, false)
	if err := ui.app.SetRoot(layout, true).EnableMouse(true).Run(); err != nil {
		// todo.SaveToFile()
		panic(err)
	}
}
