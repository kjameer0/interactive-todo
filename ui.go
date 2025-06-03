package main

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
)

type handler struct {
	Label    string
	Shortcut rune
	Action   func()
}

// create a new handler method than can be passed to a tview list
func newHandler(label string, shortcut rune, action func()) *handler {
	labelWithShortcut := fmt.Sprintf("%s) %s", keyRuneToLabel(shortcut), label)
	return &handler{Label: labelWithShortcut, Shortcut: shortcut, Action: action}
}

func keyRuneToLabel(key rune) string {
	switch key {
	// ESC button
	case rune(0):
		return "esc"
	default:
		return string(key)
	}
}

func updateOptions(ui *ui, handlers []*handler, list *tview.List) *tview.List {
	list.Clear()
	var zeroValueRune rune
	for _, handler := range handlers {
		list.AddItem(handler.Label, "", zeroValueRune, nil)
		err := ui.addGlobalEvent(handler.Shortcut, handler.Action)
		if err != nil {
			log.Fatal(err)
		}
	}
	return list
}

func createShortCutKeys(table *tview.Table, n int) []rune {
	keyList := make([]rune, 0, n)
	r := 'a'
	for i := 0; i < n; i++ {
		keyList = append(keyList, r)
		r += 1
		if r == 'z'+1 {
			r = 'A'
		}
	}
	return keyList
}

func (ui *ui) resetUI() {
	ui.optionsMenu.Clear()
	ui.output.Clear()
	ui.clearAllEvents()
}
