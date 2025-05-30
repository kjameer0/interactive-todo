# interactive-todo
what is the workflow that goes into an event:
1. event is triggered
2. new components are created
3. clear up old events
4. transfer focus

If someone presses a key then i am basically spinning up a page. So i basiclaly have to create a bunch of pages and navigate to them. Each Global Event should trigger the navigation. So if I have
```go
func createListTaskMenu(ui *ui, taskManager *todo.App) {
	ui.optionsMenu.Clear()
	ui.optionsMenu.AddItem("Return to main menu", "", '0', func() {
		//generate main menu
		generateMainOptionsMenu(ui, taskManager)
		createDefaultOutputMenu(ui, taskManager)
	})
}
```
This function should a bunch of events that will navigate to different pages. each of the functions called by the events should create the required options menu and output menu. But one output menu can have many submenus. Like the dashboard for a single task can have
