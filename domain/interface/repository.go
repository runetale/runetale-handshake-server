package repository

import "github.com/runetale/runetale-handshake-server/domain/entity"

type PeerRepository interface {
	FindByClientNodeKey(cmk string) (*entity.PeerStream, bool)
	IsCreated(cmk string) bool
	Create(peer *entity.PeerStream) error
	Delete(peer *entity.PeerStream) error
}
