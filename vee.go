package main

import (
	"bytes"
	"syscall/js"

	"github.com/jamesliu96/faux"
	"github.com/jamesliu96/vee/vee"
)

func main() {
	// const faux: (data: Bytes, key: Bytes) => Promise<Bytes>;
	vee.Global.Set("faux", vee.PromiseOf(func(this js.Value, args []js.Value) any {
		data, key := args[0], args[1]
		dataBytes := make([]byte, data.Length())
		keyBytes := make([]byte, key.Length())
		js.CopyBytesToGo(dataBytes, data)
		js.CopyBytesToGo(keyBytes, key)
		r := bytes.NewBuffer(dataBytes)
		w := bytes.NewBuffer(nil)
		if err := faux.Faux(r, w, keyBytes); err != nil {
			panic(err)
		}
		out := vee.Uint8ArrayConstructor.New(w.Len())
		js.CopyBytesToJS(out, w.Bytes())
		return out
	}))

	vee.KeepAlive()
}
