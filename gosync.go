package main

import (
	"./sync"
)

type Sync struct {
	From string
	To   string
}

type Path struct {
	PathName   string
	PathType   string
	PathLoc    string
	PathExtend string
}

type Config struct {
	LightInterval  uint64
	DiskInterval   uint64
	EntireInterval uint64

	Sources []string

	SyncItem []Sync
}

type Storage struct {
}

func main() {
	var from *sync.DescMeta
	var to *sync.DescMeta
	var sc sync.Sync
	sc.LightSync(from, to)
}
