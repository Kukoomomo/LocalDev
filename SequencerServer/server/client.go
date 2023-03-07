package server

import (
	"context"

	"github.com/Kukoomomo/LocalDev/SequencerServer/bindings"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const WsUrl = "ws://127.0.0.1:9545"
const PrvKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const SequencerAddress = "0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"

type Sequencer struct {
	Session *bindings.SequencerSession

	Cancel context.CancelFunc
}

func New() (*Sequencer, error) {
	ctx, cancel := context.WithCancel(context.Background())
	l1, err := ethclient.DialContext(ctx, WsUrl)
	if err != nil {
		cancel()
		return nil, err
	}
	chainID, err := l1.NetworkID(context.Background())
	if err != nil {
		cancel()
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(PrvKey)
	if err != nil {
		cancel()
		return nil, err
	}
	//fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		cancel()
		return nil, err
	}
	callOpts := bind.CallOpts{
		From:    transactOpts.From,
		Pending: true,
		Context: ctx,
	}
	sequencer, err := bindings.NewSequencer(common.HexToAddress(SequencerAddress), l1)
	if err != nil {
		cancel()
		return nil, err
	}
	sequencerSession := &bindings.SequencerSession{
		Contract:     sequencer,
		CallOpts:     callOpts,
		TransactOpts: *transactOpts,
	}
	return &Sequencer{
		Session: sequencerSession,
		Cancel:  cancel,
	}, nil
}
