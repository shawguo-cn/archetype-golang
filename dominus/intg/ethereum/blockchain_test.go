package ethereum

import (
	"fmt"
	"github.com/ethereum/ethash"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/pow"
	"testing"
)

func thePow() pow.PoW {
	pow, _ := ethash.NewForTesting()
	return pow
}

func theTestnetChain(db ethdb.Database, t *testing.T) *core.BlockChain {
	var eventMux event.TypeMux
	blockchain, err := core.NewBlockChain(db, params.TestnetChainConfig, thePow(), &eventMux)
	if err != nil {
		t.Error("failed creating blockchain:", err)
		t.FailNow()
		return nil
	}
	return blockchain
}

func TestPrintBlock(t *testing.T) {
	blockNum := uint64(220967)
	//Transaction to Address 0x04a3676d9bee8f51e5b9af5ba86c9a9000200eac
	bchain := theTestnetChain(TheEthdb(TestnetDBPath, t), t)

	block := bchain.GetBlockByNumber(blockNum)
	fmt.Printf("%+v\n", block.Transactions()) //TODO PANIC=runtime error: invalid memory address or nil pointer dereference

	for _, element := range block.Transactions() {
		fmt.Printf("transaction hash=%v, to=%v\n", element.Hash().Hex(), element.To().Hex()) //TODO
	}

	genesis := bchain.Genesis()
	fmt.Printf("%+v", genesis)
}

func TestPreviousBlock(t *testing.T) {

	blockNum := uint64(220967)
	//Transaction to Address 0x04a3676d9bee8f51e5b9af5ba86c9a9000200eac
	bchain := theTestnetChain(TheEthdb(TestnetDBPath, t), t)

	previousHash := bchain.GetBlockByNumber(blockNum).Header().ParentHash
	println("previous hash: " + previousHash.Hex())
	block := bchain.GetBlockByHash(previousHash)
	fmt.Printf("%+v", block)
}
