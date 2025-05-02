package main

import (
	"fmt"

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
	ui := ui{
		app:              tview.NewApplication(),
		optionsMenu:      tview.NewList(),
		output:           tview.NewFlex(),
		messageContainer: tview.NewTextView(),
	}
	_ = ui
	configPath := "./config.json"
	taskManager := todo.NewApp(configPath)
	_ = taskManager
	fmt.Println(taskManager.Config)
	//load config
	//leftoff: i added the piece where I can actually get the application config to load, I need to
		// make sure save data can save
		//write a function to generate the UI components 
}
