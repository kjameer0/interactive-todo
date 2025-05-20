package main

// Manage key presses across the application
// this file should hold functions that can update the global event handlers
// add event
// overwrite event
// delete event
// what forms will events take?
const charTotal = 62

type globalEventManager struct {
	keyEventMap map[rune]func()
}

// number of events that can be registered
func (a *ui) NewGlobalEventManager() *globalEventManager {
	return &globalEventManager{keyEventMap: make(map[rune]func(), 62)}
}
