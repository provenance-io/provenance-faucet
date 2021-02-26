package faucet

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/std"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/p2p"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

const (
	stdTxCodecType   = "cosmos-sdk/StdTx"
	msgSendCodecType = "cosmos-sdk/MsgSend"
)

type Faucet struct {
	appCli          string
	chainID         string
	keyringPassword string
	keyName         string
	faucetAddress   string
	keyMnemonic     string
	keyNodeAddr     string
	denom           string
	creditAmount    uint64
	maxCredit       uint64
	cdc             *codec.ProtoCodec
}

func NewFaucet(opts ...Option) (*Faucet, error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt(options)
	}

	chainID, err := getChainID(options.KeyNodeAddr, options.AppCli)
	if err != nil {
		return nil, err
	}

	reg := types.NewInterfaceRegistry()
	std.RegisterInterfaces(reg)
	bank.RegisterInterfaces(reg)

	e := Faucet{
		appCli:          options.AppCli,
		keyringPassword: options.KeyringPassword,
		keyName:         options.KeyName,
		keyMnemonic:     options.KeyMnemonic,
		keyNodeAddr:     options.KeyNodeAddr,
		denom:           options.Denom,
		creditAmount:    options.CreditAmount,
		maxCredit:       options.MaxCredit,
		chainID:         chainID,
		cdc:             codec.NewProtoCodec(reg),
	}
	return &e, e.loadKey()
}

func getChainID(nodeAddr string, executable string, ) (string, error) {
	output, err := cmdexec(nodeAddr, executable, []string{"status"})
	if err != nil {
		return "", err
	}

	cdc := codec.NewLegacyAmino()
	codec.RegisterEvidences(cdc)
	cryptocodec.RegisterCrypto(cdc)

	var status resultStatus
	if err := cdc.UnmarshalJSON([]byte(output), &status); err != nil {
		return "", err
	}

	return status.NodeInfo.Network, nil
}

// ResultStatus is node's info, same as Tendermint, except that we use our own
// PubKey.
type resultStatus struct {
	NodeInfo      p2p.DefaultNodeInfo
	SyncInfo      ctypes.SyncInfo
	ValidatorInfo validatorInfo
}

// ValidatorInfo is info about the node's validator, same as Tendermint,
// except that we use our own PubKey.
type validatorInfo struct {
	Address     bytes.HexBytes
	PubKey      cryptotypes.PubKey
	VotingPower int64
}
