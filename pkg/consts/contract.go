package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	IBCHostAddress = "0x06e5dEB55CAffb1339fb447d860E305249EaD495"
	IBCHandlerAddress = "0xF938fE7482Fe4d1b3f84E28F1D6407836AA27d99"
	IBCIdentifierAddress = "0x04437cDC1d0126dEB0e762e037dcEF3d5510A1f3"
	IBFT2ClientAddress = "0x4DB8e6C8BdE4c9AFCEDb590C5446c965c073BED8"
	MockClientAddress = "0xeB50cA91c99ceD8EDa25D362AdE560df9661Bc31"
	SimpleTokenAddress = "0x72f25b3D42e279917b4bd9284c22b99cBe521076"
	ICS20TransferBankAddress = "0xa37a1a9Cc31e44adFb68Da558fc1F00f77983794"
	ICS20BankAddress = "0xF251fB1Ca5445777Fed7bb760eB3c49636BA8DC3"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}

func (contractConfig) GetIBCHandlerAddress() common.Address {
	return common.HexToAddress(IBCHandlerAddress)
}

func (contractConfig) GetIBCIdentifierAddress() common.Address {
	return common.HexToAddress(IBCIdentifierAddress)
}

func (contractConfig) GetIBFT2ClientAddress() common.Address {
	return common.HexToAddress(IBFT2ClientAddress)
}

func (contractConfig) GetMockClientAddress() common.Address {
	return common.HexToAddress(MockClientAddress)
}

func (contractConfig) GetSimpleTokenAddress() common.Address {
	return common.HexToAddress(SimpleTokenAddress)
}

func (contractConfig) GetICS20TransferBankAddress() common.Address {
	return common.HexToAddress(ICS20TransferBankAddress)
}

func (contractConfig) GetICS20BankAddress() common.Address {
	return common.HexToAddress(ICS20BankAddress)
}
