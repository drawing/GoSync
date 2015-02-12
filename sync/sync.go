package sync

import (
	"bytes"
	"errors"
	"io"
)

type Sync struct {
}

func (s *Sync) Sync(root string, from, to *DescNode, ffs, tfs FS) error {
	// file
	fmeta, err := ffs.
	// directory
	files, err := ffile.ReadDir()
	if err != nil {
		return err
	}

	// upload diff file
	for _, v := range files {
		// no change
		vdesc, present := from.DescNode[v.Name()]
		if present && v.UpdateTime() == vdesc.UpdateTime {
			continue
		}

		node := DescNode{name: v.Name(), UpdateTime: v.UpdateTime(), IsDir: v.IsDir()}

		// is directory
		if v.IsDir() {
			s.Sync(root+"/"+v.Name(), node, to, ffs, tfs)
			from.DescNode[v.Name()] = node
			continue
		}

		// need upload file
		_, err := io.Copy(tfile, ffile)
		if err != nil {
			return err
		}
		from.DescNode[v.Name()] = node
	}
}

func (s *Sync) LightSync(from *DescMeta, to *DescMeta) error {
	fromfs := factory.GetFS(from.Driver)
	if fromfs == nil {
		return errors.New("from fs is nil")
	}
	tofs := factory.GetFS(to.Driver)
	if tofs == nil {
		return errors.New("to fs is nil")
	}

	return s.Sync(from, to, fromfs, tofs)
}

func (s *Sync) CompleteSync(from *DescMeta, to *DescMeta) {
}
