package types

import (
	"mvm/dvm"
)

const (
	// DVMAddressLength defines default address length (Move address length)
	DVMAddressLength = 20

	// DVMGasPrice is a gas unit price for DVM execution.
	DVMGasPrice = 1
	// DVMGasLimit defines the max gas value for DVM execution.
	DVMGasLimit = ^uint64(0)/1000 - 1
)

// DVM Event to Event conversion params.
const (
	// EventTypeProcessingGas is the initial gas for processing event type.
	EventTypeProcessingGas = 10000
	// EventTypeNoGasLevels defines number of nesting levels that do not charge gas.
	EventTypeNoGasLevels = 2
)

var (
	// DVMStdLibAddress is the Move stdlib addresses.
	DVMStdLibAddress = make([]byte, DVMAddressLength)

	DVMStdLibAddressShortStr = "0x1"
)

type (
	CompiledItem struct {
		ByteCode []byte               `json:"byte_code,omitempty"`
		Name     string               `json:"name,omitempty"`
		Methods  []*dvm.Function      `json:"methods,omitempty"`
		Types    []*dvm.Struct        `json:"types,omitempty"`
		CodeType CompiledItemCodeType `json:"code_type"`
	}

	CompiledItemCodeType int32
)

type Metadata struct {
	Metadata *dvm.Metadata `json:"metadata,omitempty"`
}

const (
	CompiledItemModule CompiledItemCodeType = 0
	CompiledItemScript CompiledItemCodeType = 1
)

// NewVMPublishModuleRequests builds a new dvmTypes.VMPublishModule VM request.
func NewVMPublishModuleRequests(signerAddrRaw string, code []byte) *dvm.VMPublishModule {
	return &dvm.VMPublishModule{
		Sender:       []byte(signerAddrRaw),
		MaxGasAmount: DVMGasLimit,
		GasUnitPrice: DVMGasPrice,
		Code:         code,
	}
}

// NewVMExecuteScriptRequest builds a new dvmTypes.VMExecuteScript VM request.
func NewVMExecuteScriptRequest(signerAddrRaw string, code []byte, blockHeight uint64, args ...ScriptArg) *dvm.VMExecuteScript {
	vmArgs := make([]*dvm.VMArgs, 0, len(args))
	for _, arg := range args {
		vmArgs = append(vmArgs, &dvm.VMArgs{
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	return &dvm.VMExecuteScript{
		Senders:      [][]byte{[]byte(signerAddrRaw)},
		MaxGasAmount: DVMGasLimit,
		GasUnitPrice: DVMGasPrice,
		Block:        blockHeight,
		Timestamp:    0,
		Code:         code,
		TypeParams:   nil,
		Args:         vmArgs,
	}
}

func init() {
	DVMStdLibAddress[DVMAddressLength-1] = 1
}
