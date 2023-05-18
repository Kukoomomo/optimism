package challenger

import (
	_ "net/http/pprof"

	eth "github.com/ethereum/go-ethereum"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"
)

// BuildOutputLogFilter creates a filter query for the L2OutputOracle contract.
//
// The `OutputProposed` event is encoded as:
// 0: bytes32 indexed outputRoot,
// 1: uint256 indexed l2OutputIndex,
// 2: uint256 indexed l2BlockNumber,
// 3: uint256 l1Timestamp
func BuildOutputLogFilter(l2ooABI *abi.ABI) (eth.FilterQuery, error) {
	// Listen for `OutputProposed` events from the L2 Output Oracle contract
	event := l2ooABI.Events["OutputProposed"]
	query := eth.FilterQuery{
		Topics: [][]common.Hash{
			{event.ID},
		},
	}

	return query, nil
}
