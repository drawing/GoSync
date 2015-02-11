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
	UpdateTime time.Time
	Md5        []byte
	IsDir      bool
	Children   []DescNode
}
