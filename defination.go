package main

import "gopkg.in/bendahl/uinput.v1"

const (
	KUp = 0
	KDown = 1
	KPress = 2
)

//Whole App Context
type AppContext struct{
	keyboard uinput.Keyboard
}

//---------Key Action --------------
type KAction struct {
	Key int
	Action int
}


//----------------- Hot Key --------------
type KHotkey struct {
	Count int
	Actions []KAction;
}