package sync

import (
	"errors"
	"io"
)

type Sync struct {
}

func (s *Sync) updateOutline(desc *Outline) error {
	fs := DefaultFactory.GetFS(desc.Driver)
	if fs == nil {
		return errors.New("desc fs is nil")
	}

	meta, err := fs.Stat(desc.Root)
	if err != nil {
		return err
	}

	var root *Node

	if desc.Tree != nil {
		if desc.Tree.UpdateTime == meta.UpdateTime() {
			return nil
		}
		root = desc.Tree
	} else {
		root = &Node{Name: meta.Name(), UpdateTime: meta.UpdateTime(), IsDir: meta.IsDir()}
	}

	if !meta.IsDir() {
		desc.Tree = root
		return nil
	}

	if root.Children == nil {
		root.Children = make(map[string]*Node)
	}

	var files []FileMeta

	dir, err := fs.Open(desc.Root)
	if err != nil {
		return nil
	}
	files, err = dir.ReadDir()
	if err != nil {
		return err
	}

	dir.Close()

	for _, v := range files {
		// no change
		vdesc, present := root.Children[v.Name()]
		if present && v.UpdateTime() == vdesc.UpdateTime {
			continue
		}

		node := &Node{Name: v.Name(), UpdateTime: v.UpdateTime(), IsDir: v.IsDir()}

		// is directory
		if v.IsDir() {
			tmpMeta := &Outline{Name: desc.Name, Root: desc.Root + "/" + v.Name(), Driver: desc.Driver}
			s.updateOutline(tmpMeta)
			node.Children = tmpMeta.Tree.Children
		}

		root.Children[v.Name()] = node
	}

	desc.Tree = root
	return nil
}

func (s *Sync) doSyncFromTo(from, to *Outline) error {
	fromfs := DefaultFactory.GetFS(from.Driver)
	if fromfs == nil {
		return errors.New("from fs is nil")
	}
	tofs := DefaultFactory.GetFS(to.Driver)
	if tofs == nil {
		return errors.New("to fs is nil")
	}

	if from.Tree == nil {
		return nil
	}

	node := &Node{
		Name:       from.Tree.Name,
		UpdateTime: from.Tree.UpdateTime,
		IsDir:      from.Tree.IsDir}

	if to.Tree != nil {
		if from.Tree.UpdateTime == to.Tree.UpdateTime {
			return nil
		}
		node.Children = to.Tree.Children
	}

	if !from.Tree.IsDir {
		// upload
		tfile, err := tofs.Create(to.Root + "/" + node.Name)
		if err != nil {
			return err
		}

		ffile, err := fromfs.Create(from.Root + "/" + node.Name)
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

	}

	// upload diff file
	for _, v := range from.Tree.Children {
		// no change
		desc, present := to.Tree.Children[v.Name]
		if present && v.UpdateTime == desc.UpdateTime {
			continue
		}

		tmpFromMeta := &Outline{Name: v.Name, Root: from.Root + "/" + v.Name, Driver: from.Driver, Tree: v}
		tmpToMeta := &Outline{Name: desc.Name, Root: to.Root + "/" + v.Name, Driver: to.Driver, Tree: desc}
		s.doSyncFromTo(tmpFromMeta, tmpToMeta)

		node.Children[v.Name] = tmpToMeta.Tree
	}

	to.Tree = node
	return nil
}

func LightSync(from *Outline, to *Outline) error {
	var s Sync

	var err error = nil

	from.LoadTree()
	to.LoadTree()

	err = s.updateOutline(from)
	if err != nil {
		return err
	}

	err = s.updateOutline(to)
	if err != nil {
		return err
	}

	err = s.doSyncFromTo(from, to)
	if err != nil {
		return err
	}

	from.SaveTree()
	to.SaveTree()

	return nil
}
