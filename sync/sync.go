package sync

type Sync struct {
}

func (s *Sync) LightSync(from *DescMeta, to *DescMeta) {
	// check from's tree and disk
}

func (s *Sync) CompleteSync(from *DescMeta, to *DescMeta) {
}
