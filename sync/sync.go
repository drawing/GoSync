package sync

import (
	"bytes"
	"errors"
	"io"
)

type Sync struct {
	fromRoot string
	toRoot   string
	fromFs   FS
	toFs     FS
}

func (s *Sync) doSyncDesc(desc *DescMeta) error {
	fs := DefaultFactory.GetFS(desc.Driver)
	if fs == nil {
		return errors.New("desc fs is nil")
	}

	// directory
	meta, err := fs.Stat(desc.Root)
	if err != nil {
		return err
	}

	var files []FileMeta

	if meta.IsDir() {
		dir, err := s.fromFs.Open(desc.Root)
		if err != nil {
			return nil
		}
		files, err = dir.ReadDir()
		if err != nil {
			return err
		}

		dir.Close()
	} else {
		files = append(files, file)
	}

	for _, v := range files {
		// no change
		vdesc, present := desc.DescNode[v.Name()]
		if present && v.UpdateTime() == vdesc.UpdateTime {
			continue
		}

		node := DescNode{Name: v.Name(), UpdateTime: v.UpdateTime(), IsDir: v.IsDir()}

		// is directory
		if v.IsDir() {
			tmpMeta := DescMeta{Name: desc.Name, Root: desc.Root + "/" + v.Name(), Driver: desc.Driver}
			s.doSyncDesc(tmpMeta)
			node.Children = tmpMeta.Tree
		}

		desc.DescNode[v.Name()] = node
	}
}

func (s *Sync) SyncFile(path string, from, to *DescNode, ffs, tfs FS) error {
	// file
	node := DescNode{name: fmeta.Name(), UpdateTime: v.fmeta(), IsDir: v.IsDir()}
	_, err := io.Copy(tfile, ffile)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sync) SyncDir(path string, from, to *DescNode) error {
	// directory
	ffile, err := s.fromFs.Open(s.fromRoot + path)
	if err != nil {
		return nil
	}

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
			s.SyncDir(path+"/"+v.Name(), node, to)
			from.DescNode[v.Name()] = node
			continue
		}

		// need upload file
		tfile, err := s.toFs.Create(s.toRoot + path + "/" + v.Name())
		if err != nil {
			return err
		}

		_, err = io.Copy(tfile, ffile)
		if err != nil {
			return err
		}

		err = tfile.Close()
		if err != nil {
			return err
		}

		from.DescNode[v.Name()] = node
	}
}

func (s *Sync) LightSync(from *DescMeta, to *DescMeta) error {
	fromfs := DefaultFactory.GetFS(from.Driver)
	if fromfs == nil {
		return errors.New("from fs is nil")
	}
	tofs := DefaultFactory.GetFS(to.Driver)
	if tofs == nil {
		return errors.New("to fs is nil")
	}

	s.fromRoot = from.Root
	s.toRoot = to.Root

	s.fromFs = fromfs
	s.toFs = tofs

	return s.Sync("", from.Tree, to.Tree)
}

func (s *Sync) CompleteSync(from *DescMeta, to *DescMeta) {
}
