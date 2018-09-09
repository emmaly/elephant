package util

import "syscall/js"

// OnUnload runs a function when the browser starts to unload
func OnUnload(fxn func()) {
	js.Global().Get("addEventListener").Invoke("beforeunload", js.NewEventCallback(0, func(_ js.Value) {
		fxn()
	}))
}

// WaitUntilUnload blocks until the browser starts to unload
func WaitUntilUnload() {
	var quit = make(chan struct{})
	OnUnload(func() {
		close(quit)
	})
	<-quit
}
