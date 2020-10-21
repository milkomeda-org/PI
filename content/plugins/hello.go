package main

import (
	"bytes"
	"context"
	"pi/platform/statement/content"
)

type Hello struct {
}

func (h Hello) Process(context context.Context, content *content.HTTPContent) interface{} {
	content.W.Write(bytes.NewBufferString("hello plugin").Bytes())
	return "hello"
}

func (h Hello) Name() string {
	return "hello"
}

func (h Hello) Type() string {
	return "test"
}

func Instance() interface{} {
	return &Hello{}
}
