package sync

import (
	"time"
)

var DefaultFactory FSFactory

type FSFactory struct {
	mpfs map[string]FS
}

func (factory *FSFactory) Register(name string, fs FS) {
	if factory.mpfs == nil {
		factory.mpfs = make(map[string]FS)
	}

	factory[name] = fs
}

func (factory *FSFactory) GetFS(name string) FS {
	if factory.mpfs == nil {
		return nil
	}

	fs, _ = factory.mpfs[name]
	return fs
}

type FS interface {
	Create(name string) (file File, err error)
	Open(name string) (file File, err error)
	Remove(name string) error
	Mkdir(name string) error
	Stat(name string) (FileMeta, error)
}

type FileMeta interface {
	IsDir() bool
	Md5Sum() []byte
	UpdateTime() time.Time
	Name() string
}

type File interface {
	ReadDir() ([]FileMeta, error)
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
	Close() error
	Stat() (FileMeta, error)
}
