package ffi

import (
	"github.com/filecoin-project/ffi-stub/cgo"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/ipfs/go-cid"
)

type FVM struct{}

type FVMOpts struct {
	FVMVersion uint64
	Externs    cgo.Externs

	Epoch          abi.ChainEpoch
	BaseFee        abi.TokenAmount
	BaseCircSupply abi.TokenAmount
	NetworkVersion network.Version
	StateBase      cid.Cid
	Manifest       cid.Cid
	Tracing        bool
}
