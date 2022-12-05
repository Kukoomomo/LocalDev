# LocalDev
```bash
export SEQUENCER_CONTRACT_ADDRESS=0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
export SEQUENCER_L1_RPC=http://localhost:9545
export USING_BVM=true

geth --datadir data init genesis.json
geth --datadir data1 init genesis.json
# geth --datadir data2 init genesis.json

geth --datadir data --password password.txt --unlock 0 --verbosity 4 --nodiscover
geth --datadir data1 --rpcport 8085 --port 30306 --verbosity 4 --nodiscover

# new
geth attach data/geth.ipc 
miner.start(1)
personal.unlockAccount(eth.accounts[0],"",1800000)
eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[0], value: web3.toWei(1, "ether")})

personal.unlockAccount(eth.accounts[1],"",1800000)

eth.sendTransaction({from: eth.accounts[1], to: eth.accounts[0], value: web3.toWei(1, "ether")})

geth attach data1/geth.ipc
admin.nodeInfo.enode
net.peerCount
admin.peers

admin.addPeer("enode://d94cf4d758b9035f1b7323e7317852c027ce3902db4bd7ed2bf3a97bc95f1efae14b78601b7d0537eaf09e1bcdcc4050d49bd0becc0864888900ebdb59656193@127.0.0.1:30303?discport=0")
```