// Copyright The MILKOMEDA Org. All rights reserved.
// Created by vinson on 2021/5/21.

package pi

import (
	"net/http"
	"time"
)

type WebHandleFunc http.Handler

type InterposerStat struct {
	Name      string
	Author    string
	CreatedAt time.Time
}

// Interposer the plugin's interface
type Interposer interface {
	Stat() *InterposerStat
	http.Handler
}

// PI the launcher
type PI struct {
	W Web
}

// New a instant of the pi launcher
func New() *PI {
	return &PI{}
}

// Router init the web router
func (pi *PI) Router(interposer Interposer) {
	pi.W.InitHandleFunc(&interposer)
}
