package storage

import "syscall/js"

func Something() {
	localStorage := js.Global().Get("localStorage")
}
