package interactor

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewNegotiationInteractorImpl,
)
