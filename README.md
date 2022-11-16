# LocalDev
```bash
export SEQUENCER_CONTRACT_ADDRESS=0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
export SEQUENCER_L1_RPC=http://localhost:9545
export USING_BVM=true

geth --datadir data init genesis.json
geth --datadir data1 init genesis.json

geth --datadir data

# new
geth attach data/geth.ipc
miner.start(1)
personal.unlockAccount(eth.accounts[0])
eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[0], value: web3.toWei(1, "ether")})

geth --datadir data1 init genesis.json
geth --datadir data1 --rpc --rpcport 8085 --port 30306

geth attach data1/geth.ipc
admin.nodeInfo.enode
net.peerCount
admin.peers

admin.addPeer("enode://b4d916b421d5e959bc27e91de1609d6d19d14ac371362777112ac9b85eb45b7afa8054f7105d4edc668c5951b0cb93f89d5920f1ad401bfe18eb6dc6399223d6@127.0.0.1:30303")
```