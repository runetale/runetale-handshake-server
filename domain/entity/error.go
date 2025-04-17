package entity

import "errors"

var (
	ErrInvalidValue     = errors.New("ERR_INVALID_VALUE")
	ErrNotFound         = errors.New("ERR_NOT_FOUND")
	ErrInvalidHeader    = errors.New("ERR_INVALID_HEADER")
	ErrInvalidPublicKey = errors.New("ERR_INVALID_PUBLIC_KEY")
	ErrNotExistentPeer  = errors.New("ERR_NOT_EXISTENT_PEER")
	ErrNotFoundPeer     = errors.New("ERR_NOT_FOUND_PEER")
	ErrPeerNotRegister  = errors.New("ERR_PEER_NOT_REGISTER")
)
