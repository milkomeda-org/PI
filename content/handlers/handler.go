// Copyright The ZHIYUN Co. All rights reserved.
// Created by vinson on 2020/10/20.

package handlers

import (
	"bytes"
	"context"
	"fmt"
	con "pi/platform/statement/content"
	"plugin"
)

var processor []interface{}

func Handle(ctx context.Context, content *con.HTTPContent) {
	for _, c := range processor {
		if v, ok := c.(con.RenderingPlugin); ok {
			r, _ := v.Render("123")
			_, _ = content.W.Write(bytes.NewBufferString(r).Bytes())
		}
	}
}

func LoadPlugin() {
	plu, err := plugin.Open("/home/vinson/go/src/PI/content/plugins/smartmd/render_smartmd.so")
	if err != nil {
		panic(err)
	}
	instance, err := plu.Lookup("Instance")
	if err != nil {
		panic(err)
	}
	if f, ok := instance.(func() interface{}); ok {
		if p, ok := f().(con.Plugin); ok {
			processor = append(processor, p)
		} else {
			fmt.Println("plugin instance return type invalid")
		}
	} else {
		fmt.Println("plugin instance execute error")
	}

}
