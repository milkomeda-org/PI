// Copyright The MILKOMEDA Org. All rights reserved.
// Created by vinson on 2021/6/2.

package pi

import "net/http"

// Interposer the plugin's interface
type Interposer interface {
	Stat() *InterposerStat
	http.Handler
}

type Weber interface {
	Router(interposer Interposer)
	HandleFunc() *Interposer
}
