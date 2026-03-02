package main

import (
	. "TaskManager/manager"
	. "TaskManager/ui"
)

// main - точка входа в приложение
func main() {
	m := NewManager()
	ui := NewUI(m)
	ui.Run()
}
