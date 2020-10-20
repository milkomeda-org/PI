package main

import (
	"context"
	"pi/platform/statement/content"
)

type Hello struct {
}

func (h Hello) Process(context context.Context, content *content.Content) interface{} {
	return "hello"
}

var P Hello
