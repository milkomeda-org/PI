// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/10/20.

package handlers

import (
	"context"
	"pi/platform/statement/content"
	"plugin"
)

var processor []content.Processor

func Handle(ctx context.Context, content *content.Content) {
	for _, c := range processor {
		c.Process(ctx, content)
	}
}

func LoadPlugin() {
	plu, err := plugin.Open("../plugins/hello.so")
	if err != nil {
		panic(err)
	}
	P, err := plu.Lookup("P")
	if err != nil {
		panic(err)
	}
	if m, ok := P.(content.Processor); ok {
		processor = append(processor, m)
	}

}
