package ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"testing"
)

//EE:key or account management
func TestCreateAccount(t *testing.T) {
	privkey, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(privkey.PublicKey)

	var result struct {
		Privkey string `json:"privkey"`
		Address string `json:"address"`
	}

	result.Privkey = common.Bytes2Hex(crypto.FromECDSA(privkey))
	result.Address = address.Hex()
	print(address.Hex())
}

//EE: Fully client side account management without any backing Ethereum node
func TestGethAccuntManager(t *testing.T) {
	am := accounts.NewManager(TEMP_KEYSTORE_DIR, accounts.StandardScryptN, accounts.StandardScryptP)
	passphrase := "welcome1"
	newPassphrase := "welcome2"

	//create new account with passphrase
	newAcc, _ := am.NewAccount(passphrase)
	fmt.Println("New Account:", newAcc.Address.Hex())
	assert.Equal(t, 1, len(am.Accounts()))

	//exports account as a JSON key, encrypted with newPassphrase.
	jsonAcc, err := am.Export(newAcc, "password", newPassphrase)
	assert.NotNil(t, err)
	println(err.Error())
	jsonAcc, err = am.Export(newAcc, passphrase, newPassphrase)
	fmt.Println("Export decrypted keyJSON:", string(jsonAcc))

	//deletes the key matched by account if the passphrase is correct.
	_ = am.DeleteAccount(newAcc, passphrase)
	assert.Equal(t, 0, len(am.Accounts()))

	//stores the given encrypted JSON key into the key directory.
	impAcc, _ := am.Import(jsonAcc, newPassphrase, passphrase)
	fmt.Println("Import decrypted keyJSON:", impAcc.Address.Hex())
	assert.Equal(t, 1, len(am.Accounts()))
}
