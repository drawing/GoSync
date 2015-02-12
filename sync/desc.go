package sync

import (
	"time"
)

type DescMeta struct {
	Name   string
	Root   string
	Driver string
	Tree   *DescNode
}

type DescNode struct {
	Name       string
	UpdateTime time.Time
	Md5        []byte
	IsDir      bool
	Children   map[string]DescNode
}
