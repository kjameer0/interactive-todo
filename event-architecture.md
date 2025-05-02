# Event architecture

The idea with this architecture is that the UI is made of three components:

1. Message box
2. output window
3. side bar

The tview library is already pretty event driven in the sense that you can register listeners so that events ca trigger side effects. The issue that I'm having is that I want to trigger effects across the entire app, and have the event behavior colocated. An event happens when there is a ui event. if someone submits a form or focus changes, that may be a reason for a ui change. I want to make a factory for UI elements for this particular tool.

Pressing the focus button for subtasks should change:

1. the side bar; it needs subtask specific list items
2. the subtask menu itself; it needs to change to having checkboxes that can be checked to change completion status


