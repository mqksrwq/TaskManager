package main

import (
	. "TaskManager/manager"
	. "TaskManager/ui"
)

func main() {
	m := NewManager()
	ui := NewUI(m)
	ui.Run()
}
