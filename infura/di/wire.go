//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/runetale/runetale-handshake-server/domain/entity"
	infura "github.com/runetale/runetale-handshake-server/infura/interface"
	repository "github.com/runetale/runetale-handshake-server/infura/repository"
	interactor "github.com/runetale/runetale-handshake-server/usecase/interactor"
	"github.com/runetale/runetale-handshake-server/utility"
)

var wireSet = wire.NewSet(
	interactor.WireSet,
	repository.WireSet,
)

// Interactor
func InitialNegotiationInteractor(cache *entity.PeerCache, logger *utility.Logger) (usecase infura.NegotiationInteractor) {
	wire.Build(wireSet)
	return
}
