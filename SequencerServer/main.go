package main

import (
	"encoding/hex"
	"fmt"

	"github.com/Kukoomomo/LocalDev/SequencerServer/bindings"
	"github.com/Kukoomomo/LocalDev/SequencerServer/server"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func main() {
	sequencer, err := server.New()
	if err != nil {
		return
	}
	//sequencer.Session.Contract.WatchSequencerCreate()
	startNum := uint64(0)
	ctx := sequencer.Session.CallOpts.Context

	createCh := make(chan *bindings.SequencerSequencerCreate, 4096)
	createSub, err := sequencer.Session.Contract.WatchSequencerCreate(&bind.WatchOpts{Context: ctx}, createCh)
	if err != nil {
		fmt.Print("ERR:can not watch create")
		return
	}
	defer createSub.Unsubscribe()
	createIter, err := sequencer.Session.Contract.FilterSequencerCreate(&bind.FilterOpts{Start: startNum, Context: ctx})
	if err != nil {
		fmt.Print("ERR:can not filter create event")
		return
	}
	for createIter.Next() {
		createCh <- createIter.Event
	}

	deleteCh := make(chan *bindings.SequencerSequencerDelete, 4096)
	deleteSub, err := sequencer.Session.Contract.WatchSequencerDelete(&bind.WatchOpts{Context: ctx}, deleteCh)
	if err != nil {
		fmt.Print("ERR:can not watch delete")
		return
	}
	defer deleteSub.Unsubscribe()
	deleteIter, err := sequencer.Session.Contract.FilterSequencerDelete(&bind.FilterOpts{Start: startNum, Context: ctx})
	if err != nil {
		fmt.Print("ERR:can not filter create event")
		return
	}
	for deleteIter.Next() {
		deleteCh <- deleteIter.Event
	}
	for {
		select {
		case ev := <-createCh:
			fmt.Printf("Create Sequencer in height: %v,%v,%v,%v \n", ev.Raw.BlockNumber, ev.Arg0, ev.Arg1, hex.EncodeToString(ev.Arg2))
		case ev := <-deleteCh:
			fmt.Printf("Delete Sequencer in height: %v,%v,%v \n", ev.Raw.BlockNumber, ev.Arg0, hex.EncodeToString(ev.Arg1))
		}
	}
}
