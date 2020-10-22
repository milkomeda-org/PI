// Copyright The ZHIYUN Co. All rights reserved.
// Created by Administrator at 2020/10/13.

package content

// Plugin plugin interface
type Plugin interface {
	Name() string
	Type() string
}

// Render content rendering
type Render interface {
	// this function should return rendered
	Rend(src string) (string, error)
}
