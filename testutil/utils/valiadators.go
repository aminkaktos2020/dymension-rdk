package utils

import (
	"testing"

	"github.com/stretchr/testify/require"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// NewValidator is a testing helper method to create validators in tests
func NewValidator(t testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey) stakingtypes.Validator {
	v, err := stakingtypes.NewValidator(operator, pubKey, stakingtypes.Description{})
	require.NoError(t, err)
	return v
}
