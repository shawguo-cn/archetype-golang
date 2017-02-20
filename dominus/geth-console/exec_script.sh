#!/usr/bin/env bash

export SCRIPT="loadScript(\"/opt/Development/go-workspace/go/src/github.com/shawguo-cn/archetype-golang/dominus/intg/ethereum/geth-console/getTransactionsByAccount.js\")"

geth --exec $SCRIPT attach http://localhost:8545