// Copyright The MILKOMEDA Org. All rights reserved.
// Created by vinson on 2021/5/21.

package pi

// Web the web feature
type Web struct {
	handleFunc *Interposer
}

func (w *Web) InitHandleFunc(interposer *Interposer) {
	w.handleFunc = interposer
}

func (w *Web) HandleFunc() *Interposer {
	return w.handleFunc
}
