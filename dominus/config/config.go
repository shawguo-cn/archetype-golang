package config


var Tendermint_RPC_URL = "http://shawguo.imwork.net:46657"

type Config struct {

	//mongodb
	MongodbUrl string
	MongodbTestDb string

	//gin options
	GinModel string
	GinPort string
}