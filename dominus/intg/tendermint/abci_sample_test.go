package tendermint

import (
	"testing"
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/bitly/go-simplejson"
	"encoding/hex"
	"github.com/shawguo-cn/archetype-golang/dominus/config"
)


//EE: data is hex encoded in abci_query result.
//./tendermint node --proxy_app=dummy
func TestDummy(t *testing.T) {

	resp, _ := http.Get(config.Tendermint_RPC_URL + "/broadcast_tx_sync?tx=\"abc123=shawguo\"");
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	resp, _ = http.Get(config.Tendermint_RPC_URL + "/abci_query?query=\"abc123\"");
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	js, _ := simplejson.NewJson(body)
	queryResult, _ := js.Get("result").GetIndex(1).Get("result").Get("Data").String()
	//data,_:= js.Get("result").GetIndex(1).Get("Data").String()

	decode, _ := hex.DecodeString(queryResult)
	js, _ = simplejson.NewJson(decode)
	value, _ := js.Get("value").String()

	fmt.Println(value)
}
