// Copyright The ZHIYUN Co. All rights reserved.
// Created by Administrator at 2020/10/13.

package content

// Content Runtime data
type Content interface {
	ContentType() string
}

// HTTPContent http data content
type HTTPContent struct {

}

func (h *HTTPContent) ContentType() string {
	return "http"
}

// TCPContent tcp data content
type TCPContent struct {

}

func (t *TCPContent) ContentType() string {
	return "tcp"
}