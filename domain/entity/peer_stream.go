package entity

import (
	"github.com/runetale/client-go/runetale/runetale/v1/negotiation"
)

type PeerStream struct {
	// only one unique key that every node has
	ClientNodeKey string
	// stream connections with grpc clients that all nodes have
	Stream negotiation.NegotiationService_ConnectServer
}

func NewPeerStream(nk string, s negotiation.NegotiationService_ConnectServer) *PeerStream {
	return &PeerStream{
		ClientNodeKey: nk,
		Stream:        s,
	}
}
