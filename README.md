# LocalDev
export SEQUENCER_CONTRACT_ADDRESS=0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0
export SEQUENCER_L1_RPC=http://localhost:9545

```bash
geth --datadir data init genesis.json
geth --datadir data

# new
geth attach data/geth.ipc
miner.start(1)
personal.unlockAccount(eth.accounts[0])
eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[0], value: web3.toWei(1, "ether")})
```