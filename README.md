# LocalDev
```bash
export SEQUENCER_CONTRACT_ADDRESS=0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
export SEQUENCER_L1_RPC=http://localhost:9545
export USING_BVM=true

geth --datadir data init genesis.json
geth --datadir data1 init genesis.json
geth --datadir data2 init genesis.json
geth --datadir data3 init genesis.json

./bin/geth --datadir data --password password.txt --unlock 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 --nodiscover --rollup.role scheduler --verbosity 4 --scheduler.batchepoch 5

geth --datadir data1 --rpcport 8085 --port 30306 --password password.txt --unlock 0x70997970c51812dc3a010c7d01b50e0d17dc79c8 --nodiscover --verbosity 4  --rollup.role sequencer

geth --datadir data2 --rpcport 8086 --port 30307 --password password.txt --unlock 0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc --nodiscover --verbosity 4 --rollup.role sequencer

geth --datadir data3 --rpcport 8087 --port 30308  --nodiscover --verbosity 4 --sequencer.mode=true

# new
geth attach data/geth.ipc 
miner.start(1)
personal.unlockAccount(eth.accounts[0],"",1800000)
eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[0], value: web3.toWei(1, "ether")})

admin.nodeInfo.id
geth attach data1/geth.ipc
admin.nodeInfo.enode
net.peerCount
admin.peers

admin.addPeer("enode://8704732f1059123c982281a7e160741fbf7019df55a2190606d3eb1e8706bbea00897c2bb02790c76268a0b2ff62272c5b9bc5ba3d8fd961f5324dd523b56c38@127.0.0.1:30303?discport=0")

eth.sendTransaction({from: eth.accounts[0], to: "0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266", value: web3.toWei(1, "ether")})

eth.sendTransaction({from: eth.accounts[0], to: "0x70997970c51812dc3a010c7d01b50e0d17dc79c8", value: web3.toWei(1, "ether")})

eth.sendTransaction({from: eth.accounts[0], to: "0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc", value: web3.toWei(1, "ether")})
```

```bash
# delete sequencer 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
yarn hardhat withdrawAll --network l1
# create sequencer 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC 
yarn hardhat createSequencer --network l1
# add deposit for sequencer 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
yarn hardhat deposit --network l1        
```