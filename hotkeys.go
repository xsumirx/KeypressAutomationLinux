package main

import "gopkg.in/bendahl/uinput.v1"

//----------------- Global Hotkey list which runs in Loop ------
var hotkeys = []KHotkey{

	//Hotket 1 - LeftCtrl + A
	KHotkey {
		Count : 3, 	//No. Action need to be taken in this Hotkey
		Actions : []KAction{
			//Press Left Control Key
			{
				Key:uinput.KeyLeftctrl,
				Action:KDown,
			},
			//Press and Release A Key
			{
				Key:uinput.KeyA,
				Action:KPress,
			},

			//Release Left Control Key
			{
				Key:uinput.KeyLeftctrl,
				Action:KUp,
			},
		},
	},

	//Hotkey 2 - LeftCtrl + C
	KHotkey {
		Count : 3, 	//No. Action need to be taken in this Hotkey
		Actions : []KAction{
			//Press Left Control Key
			{
				Key:uinput.KeyLeftctrl,
				Action:KDown,
			},
			//Press and Release C Key
			{
				Key:uinput.KeyC,
				Action:KPress,
			},

			//Release Left Control Key
			{
				Key:uinput.KeyLeftctrl,
				Action:KUp,
			},
		},
	},
}