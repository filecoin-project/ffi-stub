package ffi

import (
	"github.com/filecoin-project/filecoin-ffi/cgo"
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

	Debug         bool
	ActorRedirect cid.Cid
}

type ApplyRet struct {
	Return             []byte
	ExitCode           uint64
	GasUsed            int64
	MinerPenalty       abi.TokenAmount
	MinerTip           abi.TokenAmount
	BaseFeeBurn        abi.TokenAmount
	OverEstimationBurn abi.TokenAmount
	Refund             abi.TokenAmount
	GasRefund          int64
	GasBurned          int64
	ExecTraceBytes     []byte
	FailureInfo        string
}

func CreateFVM(*FVMOpts) (*FVM, error) {
	return &FVM{}, nil
}

func (*FVM) ApplyMessage(msgBytes []byte, chainLen uint) (*ApplyRet, error) {
	return nil, nil
}

func (*FVM) ApplyImplicitMessage(msgBytes []byte) (*ApplyRet, error) {
	return nil, nil
}

func (*FVM) Flush() (cid.Cid, error) {
	return cid.Undef, nil
}
