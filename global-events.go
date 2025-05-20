package main

import (
	"errors"

	"github.com/gdamore/tcell/v2"
)

// Manage key presses across the application
// this file should hold functions that can update the global event handlers
// add event
// overwrite event
// delete event
// what forms will events take?
const charTotal = 62

type globalEventManager struct {
	KeyEventMap map[rune]func()
	DoEventsRun bool
}

// number of events that can be registered
func (a *ui) NewGlobalEventManager() *globalEventManager {
	return &globalEventManager{KeyEventMap: make(map[rune]func(), 62)}
}

func (ui *ui) addGlobalEvent(key rune, event func()) error {
	if _, ok := ui.globalEventManager.KeyEventMap[key]; ok {
		return errors.New("event already exists")
	}
	ui.globalEventManager.KeyEventMap[key] = event
	return nil
}

func (ui *ui) removeGlobalEvent(key rune) {
	delete(ui.globalEventManager.KeyEventMap, key)
}

func (ui *ui) updateGlobalEvent(key rune, event func()) {
	ui.removeGlobalEvent(key)
	ui.addGlobalEvent(key, event)
}

func (ui *ui) getEvent(key rune) (func(), error) {
	if _, ok := ui.globalEventManager.KeyEventMap[key]; !ok {
		return nil, errors.New("event does not exist")
	}
	return ui.globalEventManager.KeyEventMap[key], nil
}

// prevent events from triggering(if the user focuses a form)
func (ui *ui) setDoEventsRun(doEventsRun bool) {
	ui.globalEventManager.DoEventsRun = doEventsRun
}

func (ui *ui) getDoEventsRun() bool {
	return ui.globalEventManager.DoEventsRun
}

func (ui *ui) registerEvents() {
	ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if !ui.getDoEventsRun() {
			return event
		}
		registeredEvent, err := ui.getEvent(event.Rune())
		if err != nil {
			return nil
		}
		registeredEvent()
		return nil
	})
}
