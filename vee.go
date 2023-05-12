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
		data := args[0]
		key := args[1]
		dataLen := data.Length()
		dataBytes := make([]byte, dataLen)
		keyLen := key.Length()
		keyBytes := make([]byte, keyLen)
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
