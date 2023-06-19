package main

import (
	"log"

	tendermint "github.com/hyperledger-labs/yui-relayer/chains/tendermint/module"
	"github.com/hyperledger-labs/yui-relayer/cmd"
	mock "github.com/hyperledger-labs/yui-relayer/provers/mock/module"
	ethereum "github.com/neoiss/yui-ibc-solidity/pkg/relay/module"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		ethereum.Module{},
		mock.Module{},
	); err != nil {
		log.Fatal(err)
	}
}
