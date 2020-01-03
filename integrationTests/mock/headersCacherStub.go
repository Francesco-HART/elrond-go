package mock

import (
	"errors"
	"github.com/ElrondNetwork/elrond-go/data"
)

type HeadersCacherStub struct {
	AddCalled                           func(headerHash []byte, header data.HeaderHandler)
	RemoveHeaderByHashCalled            func(headerHash []byte)
	RemoveHeaderByNonceAndShardIdCalled func(hdrNonce uint64, shardId uint32)
	GetHeaderByNonceAndShardIdCalled    func(hdrNonce uint64, shardId uint32) ([]data.HeaderHandler, [][]byte, error)
	GetHeaderByHashCalled               func(hash []byte) (data.HeaderHandler, error)
	ClearCalled                         func()
	RegisterHandlerCalled               func(handler func(shardHeaderHash []byte))
	KeysCalled                          func(shardId uint32) []uint64
	LenCalled                           func() int
	MaxSizeCalled                       func() int
}

func (hcs *HeadersCacherStub) AddHeader(headerHash []byte, header data.HeaderHandler) {
	if hcs.AddCalled != nil {
		hcs.AddCalled(headerHash, header)
	}
}

func (hcs *HeadersCacherStub) RemoveHeaderByHash(headerHash []byte) {
	if hcs.RemoveHeaderByHashCalled != nil {
		hcs.RemoveHeaderByHashCalled(headerHash)
	}
}

func (hcs *HeadersCacherStub) RemoveHeaderByNonceAndShardId(hdrNonce uint64, shardId uint32) {
	if hcs.RemoveHeaderByNonceAndShardIdCalled != nil {
		hcs.RemoveHeaderByNonceAndShardIdCalled(hdrNonce, shardId)
	}
}

func (hcs *HeadersCacherStub) GetHeaderByNonceAndShardId(hdrNonce uint64, shardId uint32) ([]data.HeaderHandler, [][]byte, error) {
	if hcs.GetHeaderByNonceAndShardIdCalled != nil {
		return hcs.GetHeaderByNonceAndShardIdCalled(hdrNonce, shardId)
	}
	return nil, nil, errors.New("err")
}

func (hcs *HeadersCacherStub) GetHeaderByHash(hash []byte) (data.HeaderHandler, error) {
	if hcs.GetHeaderByHashCalled != nil {
		return hcs.GetHeaderByHashCalled(hash)
	}
	return nil, nil
}

func (hcs *HeadersCacherStub) Clear() {
	if hcs.ClearCalled != nil {
		hcs.ClearCalled()
	}
}

func (hcs *HeadersCacherStub) RegisterHandler(handler func(shardHeaderHash []byte)) {
	if hcs.RegisterHandlerCalled != nil {
		hcs.RegisterHandlerCalled(handler)
	}
}

func (hcs *HeadersCacherStub) Keys(shardId uint32) []uint64 {
	if hcs.KeysCalled != nil {
		return hcs.KeysCalled(shardId)
	}
	return nil
}

func (hcs *HeadersCacherStub) Len() int {
	return 0
}

func (hcs *HeadersCacherStub) MaxSize() int {
	return 100
}

func (hcs *HeadersCacherStub) IsInterfaceNil() bool {
	return hcs == nil
}
