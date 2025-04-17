package entity

import "sync"

// peers local cache
type PeerCache struct {
	Cmap sync.Map // Cmap is cache map
}

func NewPeerCache() *PeerCache {
	return &PeerCache{
		Cmap: sync.Map{},
	}
}
