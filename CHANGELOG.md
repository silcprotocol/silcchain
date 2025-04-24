# CHANGELOG

## UNRELEASED

### DEPENDENCIES

- [\#31](https://github.com/silcprotocol/silcchain/pull/31) Migrated example_chain to silcd
- Migrated evmos/go-ethereum to cosmos/go-ethereum
- Migrated evmos/cosmos-sdk to cosmos/cosmos-sdk

### BUG FIXES

- Fixed example chain's cmd by adding NoOpEVMOptions to tmpApp in root.go
- Added RPC support for `--legacy` transactions (Non EIP-1559)

### IMPROVEMENTS

### FEATURES

### STATE BREAKING

- Refactored evmos/os into cosmos/evm
- Renamed x/evm to x/vm
- Renamed protobuf files from evmos to cosmos org
- [\#83](https://github.com/silcprotocol/silcchain/pull/83) Remove base fee v1 from x/feemarket

### API-Breaking

- Refactored evmos/os into cosmos/evm
- Renamed x/evm to x/vm
- Renamed protobuf files from evmos to cosmos org
