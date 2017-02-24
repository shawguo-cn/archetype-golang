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
func newTrie(root common.Hash) *trie.Trie {
	db, _ := ethdb.NewLDBDatabase("/home/deploy/.ethereum/geth/chaindata", 0, 0)
	trie, err := trie.New(root, db); if err != nil {
		println(err.Error())
		os.Exit(10)
	}
	return trie
}

//EE: Ethereum account state
func TestStateRootIterator(t *testing.T) {
	trie := newTrie(common.HexToHash(stateRootHash))
	i := 0;
	for it := trie.Iterator(); it.Next(); {
		i = i + 1;
		fmt.Printf("Adress:%v, %vth\n", common.BytesToAddress(it.Key).Hex(), i)
		//decode account
		var value = new(state.Account)
		rlp.DecodeBytes(it.Value, value)
		fmt.Printf("balance:%v nonce:%v root:%v\n", value.Balance, value.Nonce, value.Root.Hex())
		//i = i + 1; if i == 100 {
		//	return
		//}
	}
}

func TestStateLookup(t *testing.T) {
	trie := newTrie(common.HexToHash("0x56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421"))
	address := common.HexToAddress(lookupAccount)
	val := trie.Get(address[:])
	if len(val) == 0 {
		//TODO nil?
		os.Exit(10)
	}

	var value = new(state.Account)
	rlp.DecodeBytes(val, value)
	fmt.Printf("balance:%v nonce:%v root:%v\n", value.Balance, value.Nonce, value.Root.Hex())
}

func TestBlankTrie(t *testing.T) {
	blankTrie := newTrie(common.Hash{})
	val, _ := rlp.EncodeToBytes("hello")
	blankTrie.Update([]byte("\x01\x01\x02"), val)
	println(common.Bytes2Hex(blankTrie.Root()))

	//*blankTrie.roo

}