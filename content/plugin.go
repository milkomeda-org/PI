// Copyright The ZHIYUN Co. All rights reserved.
// Created by Administrator at 2020/10/13.

package content

import "context"

// Processor Data processor
type Processor interface {
	Process(context context.Context, content *Content) interface{}
}
