package main

import "time"

//all message box related functionality

//TODO: set up channel to receive text events
//TODO: allow ability to tell duration
//TODO:
//TODO:

func messageReceiver(ui *ui) {
	var timer *time.Timer
	for msg := range ui.messageChannel {
		ui.messageContainer.SetText(msg)

		if timer != nil {
			timer.Reset(3 * time.Second)
		} else {
			timer = time.NewTimer(3 * time.Second)
		}

		go func() {
			<-timer.C
			ui.messageContainer.SetText("")
			ui.app.Draw()
		}()
	}
}
