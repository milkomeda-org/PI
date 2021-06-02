// Copyright The MILKOMEDA Org. All rights reserved.
// Created by vinson on 2021/5/21.

package pi

import (
	"net/http"
)

type WebHandleFunc http.Handler

// PI the launcher
type PI struct {
	W Weber
}

// New a instant of the pi launcher
func New() *PI {
	return &PI{}
}

// NewWeb a instant of the pi launcher with web
func NewWeb(weber Weber) *PI {
	return &PI{W: weber}
}
