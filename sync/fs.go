package sync

type FSFactory struct {
}

func (factory *FSFactory) Register() {
}

type FS interface {
	Create()
	Delete()
}

type FileMeta interface {
	IsDir()
	Md5Sum()
	UpdateTime()
}

type File interface {
	ReadDir()
	Read()
	Write()
	Close()
}
