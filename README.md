# LocalDev
```bash
export SEQUENCER_CONTRACT_ADDRESS=0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
export SEQUENCER_L1_RPC=http://localhost:9545
export USING_BVM=true

geth --datadir data init genesis.json
geth --datadir data1 init genesis.json

geth --datadir data --password '' --unlock 0 --verbosity 4
 
# new
geth attach data/geth.ipc 
miner.start(1)
personal.unlockAccount(eth.accounts[0],"",1800000)
eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[0], value: web3.toWei(1, "ether")})

geth --datadir data1 --rpcport 8085 --port 30306 --verbosity 4

geth attach data1/geth.ipc
admin.nodeInfo.enode
net.peerCount
admin.peers

admin.addPeer("enode://d51ca4d6749de54a08d054796ecb6b60a124ec4aeda0a29fbbaa5b110df310c6ff8b480f1efefae996d22b47ca4f1b57112c8e869d391f2b105e98d1622bf762@127.0.0.1:30303")


```