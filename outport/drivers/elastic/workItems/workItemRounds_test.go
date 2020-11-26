package workItems_test

import (
	"errors"
	"testing"

	"github.com/ElrondNetwork/elrond-go/outport/drivers/elastic/workItems"
	"github.com/ElrondNetwork/elrond-go/outport/mock"
	"github.com/ElrondNetwork/elrond-go/outport/types"
	"github.com/stretchr/testify/require"
)

func TestItemRounds_Save(t *testing.T) {
	called := false
	itemRounds := workItems.NewItemRounds(
		&mock.ElasticProcessorStub{
			SaveRoundsInfoCalled: func(infos []types.RoundInfo) error {
				called = true
				return nil
			},
		},
		[]types.RoundInfo{
			{},
		},
	)
	require.False(t, itemRounds.IsInterfaceNil())

	err := itemRounds.Save()
	require.NoError(t, err)
	require.True(t, called)
}

func TestItemRounds_SaveRoundsShouldErr(t *testing.T) {
	localErr := errors.New("local err")
	itemRounds := workItems.NewItemRounds(
		&mock.ElasticProcessorStub{
			SaveRoundsInfoCalled: func(infos []types.RoundInfo) error {
				return localErr
			},
		},
		[]types.RoundInfo{
			{},
		},
	)
	require.False(t, itemRounds.IsInterfaceNil())

	err := itemRounds.Save()
	require.Equal(t, localErr, err)
}