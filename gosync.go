package main

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

	Sync []SyncItem
}

type Storage struct {
}

func main() {
}
