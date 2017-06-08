package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethdb"
	"os"
	"testing"
)

const (
	//public
	PublicDBPath  = "/home/deploy/.ethereum/geth/chaindata"
	StateRootHash = "0x2096f2080d0a504b5ba739867b6000068474831d5d12f81360eb45d36f125676"
	LookupAccount = "0xa299a0e687e5b750eddd8c108d04e2753ae3c145"
	//testnet
	TestnetDBPath = "/home/deploy/.ethereum/testnet/geth/chaindata"

	TEMP_KEYSTORE_DIR = "/tmp/keystore"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain Setup")
	os.RemoveAll(TEMP_KEYSTORE_DIR)
	os.Exit(m.Run())
}

func TheEthdb(file string, t *testing.T) ethdb.Database {

	db, err := ethdb.NewLDBDatabase(TestnetDBPath, 0, 0)
	if err != nil {
		t.Error("failed to open leveldb:", err)
		t.FailNow()
		return nil
	}
	return db
}
