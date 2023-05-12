//go:build js && wasm

package vee

import "syscall/js"

var (
	global                js.Value
	promiseConstructor    js.Value
	errorConstructor      js.Value
	Uint8ArrayConstructor js.Value
	Undefined             js.Value
	Global                js.Value
)

func init() {
	global = js.Global()
	promiseConstructor = global.Get("Promise")
	errorConstructor = global.Get("Error")
	Uint8ArrayConstructor = global.Get("Uint8Array")
	Undefined = js.Undefined()
	global.Set("Vee", make(map[string]any))
	Global = global.Get("Vee")
}

// Keep thread alive
func KeepAlive() {
	<-make(chan struct{})
}
