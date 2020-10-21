// Copyright The ZHIYUN Co. All rights reserved.
// Created by Administrator at 2020/10/13.

package content

// Plugin plugin interface
type Plugin interface {
	Name() string
	Type() string
}

// RenderingPlugin content rendering
type RenderingPlugin interface {
	// this function should return rendered
	Render(src string) (string, error)
}
