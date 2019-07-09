// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*headerMarshaling)(nil)

func (h Header) MarshalJSON() ([]byte, error) {
	type Header struct {
		ParentHash               common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash                common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase                 common.Address `json:"miner"            gencodec:"required"`
		Root                     common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash                   common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash              common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom                    Bloom          `json:"logsBloom"        gencodec:"required"`
		Difficulty               *hexutil.Big   `json:"difficulty"       gencodec:"required"`
		Number                   *hexutil.Big   `json:"number"           gencodec:"required"`
		GasLimit                 hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed                  hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		Time                     *hexutil.Big   `json:"timestamp"        gencodec:"required"`
		Extra                    hexutil.Bytes  `json:"extraData"        gencodec:"required"`
		MixDigest                common.Hash    `json:"mixHash"          gencodec:"required"`
		Nonce                    BlockNonce     `json:"nonce"            gencodec:"required"`
		PosWeight                uint32         `json:"posWeight"        gencodec:"required"`
		PosOldMatureSupply       *hexutil.Big   `json:"posOldMatureSupply"    gencodec:"required"`
		PosLastMatureCycleSupply *hexutil.Big   `json:"posLastMatureCycleSupply"    gencodec:"required"`
		PosLastCycleSupply       *hexutil.Big   `json:"posLastCycleSupply"    gencodec:"required"`
		PowOldMatureSupply       *hexutil.Big   `json:"powOldMatureSupply"    gencodec:"required"`
		PowLastMatureCycleSupply *hexutil.Big   `json:"powLastMatureCycleSupply"    gencodec:"required"`
		PowLastCycleSupply       *hexutil.Big   `json:"powLastCycleSupply"    gencodec:"required"`
		PosProduction            *hexutil.Big   `json:"posProduction"    gencodec:"required"`
		PowProduction            *hexutil.Big   `json:"powProduction"    gencodec:"required"`
		Hash                     common.Hash    `json:"hash"`
	}
	var enc Header
	enc.ParentHash = h.ParentHash
	enc.UncleHash = h.UncleHash
	enc.Coinbase = h.Coinbase
	enc.Root = h.Root
	enc.TxHash = h.TxHash
	enc.ReceiptHash = h.ReceiptHash
	enc.Bloom = h.Bloom
	enc.Difficulty = (*hexutil.Big)(h.Difficulty)
	enc.Number = (*hexutil.Big)(h.Number)
	enc.GasLimit = hexutil.Uint64(h.GasLimit)
	enc.GasUsed = hexutil.Uint64(h.GasUsed)
	enc.Time = (*hexutil.Big)(h.Time)
	enc.Extra = h.Extra
	enc.MixDigest = h.MixDigest
	enc.Nonce = h.Nonce
	enc.PosWeight = h.PosWeight
	enc.PosOldMatureSupply = (*hexutil.Big)(h.PosOldMatureSupply)
	enc.PosLastMatureCycleSupply = (*hexutil.Big)(h.PosLastMatureCycleSupply)
	enc.PosLastCycleSupply = (*hexutil.Big)(h.PosLastCycleSupply)
	enc.PowOldMatureSupply = (*hexutil.Big)(h.PowOldMatureSupply)
	enc.PowLastMatureCycleSupply = (*hexutil.Big)(h.PowLastMatureCycleSupply)
	enc.PowLastCycleSupply = (*hexutil.Big)(h.PowLastCycleSupply)
	enc.PosProduction = (*hexutil.Big)(h.PosProduction)
	enc.PowProduction = (*hexutil.Big)(h.PowProduction)
	enc.Hash = h.Hash()
	return json.Marshal(&enc)
}

func (h *Header) UnmarshalJSON(input []byte) error {
	type Header struct {
		ParentHash               *common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash                *common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase                 *common.Address `json:"miner"            gencodec:"required"`
		Root                     *common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash                   *common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash              *common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom                    *Bloom          `json:"logsBloom"        gencodec:"required"`
		Difficulty               *hexutil.Big    `json:"difficulty"       gencodec:"required"`
		Number                   *hexutil.Big    `json:"number"           gencodec:"required"`
		GasLimit                 *hexutil.Uint64 `json:"gasLimit"         gencodec:"required"`
		GasUsed                  *hexutil.Uint64 `json:"gasUsed"          gencodec:"required"`
		Time                     *hexutil.Big    `json:"timestamp"        gencodec:"required"`
		Extra                    *hexutil.Bytes  `json:"extraData"        gencodec:"required"`
		MixDigest                *common.Hash    `json:"mixHash"          gencodec:"required"`
		Nonce                    *BlockNonce     `json:"nonce"            gencodec:"required"`
		PosWeight                uint32          `json:"posWeight"        gencodec:"required"`
		PosOldMatureSupply       *hexutil.Big    `json:"posOldMatureSupply"    gencodec:"required"`
		PosLastMatureCycleSupply *hexutil.Big    `json:"posLastMatureCycleSupply"    gencodec:"required"`
		PosLastCycleSupply       *hexutil.Big    `json:"posLastCycleSupply"    gencodec:"required"`
		PowOldMatureSupply       *hexutil.Big    `json:"powOldMatureSupply"    gencodec:"required"`
		PowLastMatureCycleSupply *hexutil.Big    `json:"powLastMatureCycleSupply"    gencodec:"required"`
		PowLastCycleSupply       *hexutil.Big    `json:"powLastCycleSupply"    gencodec:"required"`
		PosProduction            *hexutil.Big    `json:"posProduction"    gencodec:"required"`
		PowProduction            *hexutil.Big    `json:"powProduction"    gencodec:"required"`
	}
	var dec Header
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.ParentHash == nil {
		return errors.New("missing required field 'parentHash' for Header")
	}
	h.ParentHash = *dec.ParentHash
	if dec.UncleHash == nil {
		return errors.New("missing required field 'sha3Uncles' for Header")
	}
	h.UncleHash = *dec.UncleHash
	if dec.Coinbase == nil {
		return errors.New("missing required field 'miner' for Header")
	}
	h.Coinbase = *dec.Coinbase
	if dec.Root == nil {
		return errors.New("missing required field 'stateRoot' for Header")
	}
	h.Root = *dec.Root
	if dec.TxHash == nil {
		return errors.New("missing required field 'transactionsRoot' for Header")
	}
	h.TxHash = *dec.TxHash
	if dec.ReceiptHash == nil {
		return errors.New("missing required field 'receiptsRoot' for Header")
	}
	h.ReceiptHash = *dec.ReceiptHash
	if dec.Bloom == nil {
		return errors.New("missing required field 'logsBloom' for Header")
	}
	h.Bloom = *dec.Bloom
	if dec.Difficulty == nil {
		return errors.New("missing required field 'difficulty' for Header")
	}
	h.Difficulty = (*big.Int)(dec.Difficulty)
	if dec.Number == nil {
		return errors.New("missing required field 'number' for Header")
	}
	h.Number = (*big.Int)(dec.Number)
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for Header")
	}
	h.GasLimit = uint64(*dec.GasLimit)
	if dec.GasUsed == nil {
		return errors.New("missing required field 'gasUsed' for Header")
	}
	h.GasUsed = uint64(*dec.GasUsed)
	if dec.Time == nil {
		return errors.New("missing required field 'timestamp' for Header")
	}
	h.Time = (*big.Int)(dec.Time)
	if dec.Extra == nil {
		return errors.New("missing required field 'extraData' for Header")
	}
	h.Extra = *dec.Extra
	if dec.MixDigest == nil {
		return errors.New("missing required field 'mixHash' for Header")
	}
	h.MixDigest = *dec.MixDigest
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for Header")
	}
	h.Nonce = *dec.Nonce
	h.PosWeight = dec.PosWeight
	if dec.PosOldMatureSupply == nil {
		return errors.New("missing required field 'posOldMatureSupply' for Header")
	}
	h.PosOldMatureSupply = (*big.Int)(dec.PosOldMatureSupply)
	if dec.PosLastMatureCycleSupply == nil {
		return errors.New("missing required field 'posLastMatureCycleSupply' for Header")
	}
	h.PosLastMatureCycleSupply = (*big.Int)(dec.PosLastMatureCycleSupply)
	if dec.PosLastCycleSupply == nil {
		return errors.New("missing required field 'posLastCycleSupply' for Header")
	}
	h.PosLastCycleSupply = (*big.Int)(dec.PosLastCycleSupply)
	if dec.PowOldMatureSupply == nil {
		return errors.New("missing required field 'powOldMatureSupply' for Header")
	}
	h.PowOldMatureSupply = (*big.Int)(dec.PowOldMatureSupply)
	if dec.PowLastMatureCycleSupply == nil {
		return errors.New("missing required field 'powLastMatureCycleSupply' for Header")
	}
	h.PowLastMatureCycleSupply = (*big.Int)(dec.PowLastMatureCycleSupply)
	if dec.PowLastCycleSupply == nil {
		return errors.New("missing required field 'powLastCycleSupply' for Header")
	}
	h.PowLastCycleSupply = (*big.Int)(dec.PowLastCycleSupply)
	if dec.PosProduction == nil {
		return errors.New("missing required field 'posProduction' for Header")
	}
	h.PosProduction = (*big.Int)(dec.PosProduction)
	if dec.PowProduction == nil {
		return errors.New("missing required field 'powProduction' for Header")
	}
	h.PowProduction = (*big.Int)(dec.PowProduction)
	return nil
}
