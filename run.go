package main

func (action *KAction) run(ctx *AppContext){
	switch action.Action {
	case KUp:
		ctx.keyboard.KeyUp(action.Key)
	case KDown:
		ctx.keyboard.KeyDown(action.Key)
	case KPress:
		ctx.keyboard.KeyPress(action.Key)
	}
}



func (hotkey *KHotkey) run(ctx *AppContext){
	for i,action := range hotkey.Actions{
		if i >= hotkey.Count {
			continue
		} 
		action.run(ctx)
	}
}