# IBC-Solidity

![Test](https://github.com/neoiss/yui-ibc-solidity/workflows/Test/badge.svg)
[![GoDoc](https://godoc.org/github.com/neoiss/yui-ibc-solidity?status.svg)](https://pkg.go.dev/github.com/neoiss/yui-ibc-solidity?tab=doc)

[IBC](https://github.com/cosmos/ibc) implementations in Solidity.

This is available not only for Ethereum and Hyperledger Besu, but also for Polygon PoS and other blockchains that supports EVM.

NOTE: This is yet pre-beta non-production-quality software.

## Features

- Implementation of [ICS](https://github.com/cosmos/ibc/tree/master/spec/core)
- Implementation of [ICS-20](https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer)
- [ERC-20 Token Transfer](./contracts/app/ICS20TransferBank.sol)

## Documents

Please see [here](./docs/README.md).

In addition, a tutorial is [here](https://labs.hyperledger.org/yui-docs/yui-ibc-solidity/).

## Supported Light Client

You can deploy a Light Client that implements [the IClient interface](./contracts/core/IClient.sol) to [integrate with IBC-Solidity](./docs/architecture.md#ibcclient).

Here are some such examples:
- [IBFT 2.0 Light Client](./contracts/core/IBFT2Client.sol)
- [Tendermint Light Client](https://github.com/datachainlab/tendermint-sol/tree/use-ibc-sol-hmy)
- [Mock Client](./contracts/core/MockClient.sol)

## Related projects

- A demo of trustless bridge between Harmony and Cosmos(Tendermint): https://github.com/datachainlab/harmony-cosmos-bridge-demo
- A demo of trustless bridge between Celo and Cosmos: https://github.com/ChorusOne/celo-cosmos-bridge

## Development and Testing

Launch two Besu chains(ethereum-compatible) with the contracts deployed with the following command:

```sh
# If NO_GEN_CODE is empty, setup-script will generate a proto3 marshaler in solidity
$ NO_GEN_CODE=1 ./scripts/setup.sh testtwochainz
```

After launch the chains, execute the following command:

```
$ make e2e-test
```

## E2E-test with IBC-Relayer

An example of E2E with IBC-Relayer([yui-relayer](https://github.com/hyperledger-labs/yui-relayer)) can be found here:
- https://github.com/datachainlab/yui-relayer-build/tree/v0.2/tests/cases/tm2eth
- https://github.com/datachainlab/yui-relayer-build/blob/v0.2/.github/workflows/v0.2-tm2eth.yml

## For Developers

To develop this project, you need the code generator [solidity-protobuf](https://github.com/datachainlab/solidity-protobuf) to generate encoders and decoders in solidity from proto files.

Currently, you need to use [this version](https://github.com/datachainlab/solidity-protobuf/tree/fce34ce0240429221105986617f64d8d4261d87d).

## Maintainers

- [Jun Kimura](https://github.com/bluele)
