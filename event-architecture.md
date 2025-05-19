# Event Architecture

The idea with this architecture is that the UI is made of three components:

1. Message box
2. Output window
3. Side bar

The `tview` library is already pretty event-driven in the sense that you can register listeners so that events can trigger side effects. The issue that I'm having is that I want to trigger effects across the entire app, and have the event behavior colocated.

An event happens when there is a UI event. If someone submits a form or focus changes, that may be a reason for a UI change. I want to make a factory for UI elements for this particular tool.

Pressing the focus button for subtasks should change:

1. The side bar – it needs subtask-specific list items
2. The subtask menu itself – it needs to change to having checkboxes that can be checked to change completion status

## What is the workflow for event-driven architecture?

If I have an input capture:

```go
cells.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
 if event.Rune() == rune(tcell.KeyEnter) {
  if selectedTask == 0 {
   return event
  }
  t, ok := taskMap[selectedTask]
  if ok {
   ui.messageContainer.
    SetText(t.Name + " \nwas selected").
    SetTextColor(tcell.ColorGreen)

   ui.optionsMenu.Clear()
   optionsHandlers := generateSingleTaskOptionsHandlers(ui, app)
   for _, opt := range optionsHandlers {
    action := opt.Action
    ui.optionsMenu.AddItem(opt.Label, "", opt.Shortcut, action)
   }
  // make single task ui
   layout := createTaskDashBoard(ui, app, t)
   ui.output.Clear()
   ui.output.AddItem(layout, 0, 1, true)

   ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
    switch event.Rune() {
    case '0':
     resetMainMenu(ui, app)
    }
    return event
   })
  }
 } else {
  selectedTask = event.Rune()
 }
 return event
})
```
So if I have this event what do I have access to?
I'm triggering multiple ui changes so

