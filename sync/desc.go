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
	Md5        string
	IsDir      bool
	Children   []DescNode
}
