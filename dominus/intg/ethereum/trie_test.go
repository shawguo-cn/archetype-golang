package ethereum

import (
	"testing"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/core/state"
	"os"
	"fmt"
)

const (
	//public
	levelDb = "/home/deploy/.ethereum/geth/chaindata"
	stateRootHash = "0x2096f2080d0a504b5ba739867b6000068474831d5d12f81360eb45d36f125676"
	lookupAccount = "0xa299a0e687e5b750eddd8c108d04e2753ae3c145"

	//testnet


)

// Used for testing
func newTrieFromTestnet() *trie.Trie {
	db, _ := ethdb.NewLDBDatabase("/home/deploy/.ethereum/geth/chaindata", 0, 0)
	trie, err := trie.New(common.HexToHash("0x2096f2080d0a504b5ba739867b6000068474831d5d12f81360eb45d36f125676"), db); if err != nil {
		println(err.Error())
		os.Exit(10)
	}
	return trie
}

func TestTrieIterator(t *testing.T) {
	trie := newTrieFromTestnet()
	i := 0;
	for it := trie.Iterator(); it.Next(); {
		fmt.Println(common.BytesToAddress(it.Key).Hex())
		//decode account
		var value = new(state.Account)
		rlp.DecodeBytes(it.Value, value)
		fmt.Printf("balance:%v nonce:%v root:%v\n", value.Balance, value.Nonce, value.Root.Hex())
		i = i + 1; if i == 100 {
			return
		}
	}
}

func TestTrieLookup(t *testing.T) {
	trie := newTrieFromTestnet()
	address := common.HexToAddress(lookupAccount)
	val := trie.Get(address[:])
	if len(val) == 0 {
		os.Exit(10)
	}

	var value = new(state.Account)
	rlp.DecodeBytes(val, value)
	fmt.Printf("balance:%v nonce:%v root:%v\n", value.Balance, value.Nonce, value.Root.Hex())
}