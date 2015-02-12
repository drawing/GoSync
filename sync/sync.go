package sync

import (
	"bytes"
	"errors"
)

type Sync struct {
}

func (s *Sync) Sync(from, to *DescNode, fromfs, tofs FS) error {
	// file
	if !from.IsDir {
		if to != nil && bytes.Compare(from.Md5, to.Md5) == 0 {
			// to is newest file
			return nil
		}
		// to is not the newest file, upload
		return nil
	}

	// dir
	for _, v := range from.DescNode {
		if v.UpdateTime != xx.UpdateTime {
			// need upload
		}
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
