package repository

import (
	"fmt"
	"sync"

	"github.com/runetale/runetale-handshake-server/domain/entity"
	repository "github.com/runetale/runetale-handshake-server/domain/interface"
	"github.com/runetale/runetale-handshake-server/utility"
)

type PeerRepositoryImpl struct {
	cache *entity.PeerCache
	mu    *sync.Mutex

	log *utility.Logger
}

func NewPeerRepositoryImpl(cache *entity.PeerCache, log *utility.Logger) repository.PeerRepository {
	return &PeerRepositoryImpl{
		cache: cache,
		mu:    &sync.Mutex{},

		log: log,
	}
}

func (r *PeerRepositoryImpl) FindByClientNodeKey(cmk string) (*entity.PeerStream, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if load, ok := r.cache.Cmap.Load(cmk); ok {
		return load.(*entity.PeerStream), ok
	}

	return nil, false
}

func (r *PeerRepositoryImpl) IsCreated(cmk string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.cache.Cmap.Load(cmk); ok {
		return ok
	}

	return false
}

func (r *PeerRepositoryImpl) Create(peer *entity.PeerStream) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.cache.Cmap.Store(peer.ClientNodeKey, peer)
	return nil
}

func (r *PeerRepositoryImpl) Delete(peer *entity.PeerStream) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, loaded := r.cache.Cmap.LoadAndDelete(peer.ClientNodeKey)
	if loaded {
		r.log.Logger.Info(fmt.Sprintf("peer deregisterd %s", peer.ClientNodeKey))
		return nil
	} else {
		return entity.ErrNotExistentPeer
	}
}
