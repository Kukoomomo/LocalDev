import "@nomiclabs/hardhat-web3";
import "@nomiclabs/hardhat-ethers";
import "@nomiclabs/hardhat-waffle";

import { task, types } from "hardhat/config";
import { ethers } from "ethers";

const l1RpcProvider = new ethers.providers.JsonRpcProvider('http://127.0.0.1:9545')
const amount = ethers.utils.parseEther('100')
const deposit = ethers.utils.parseEther('1')
const prvKeys = [
  '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
  '0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d',
  // '0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a',
  // '0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6',
  // '0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a',
  // '0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba'
]
const nodeIDs = [
  '0x7a6b7d46da49c9c5045a6ecb387415d7827506f1aff9bead32240772a39544d4',
  '0xd4e8928003dd91c37eb6dfbef9d13b91a17b8e46a8208994153f8aade1597877',
  // '0x2a1776580e7155022477c406dbd091983712d222d216d5e11d3c99b4e47671a8',
  // '0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6',
  // '0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a',
  // '0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba'
]
var wallets = new Array();

for (var i = 0; i < prvKeys.length; i++) {
  wallets.push(new ethers.Wallet(prvKeys[i], l1RpcProvider))
}

task("deploy")
  .setAction(async (taskArgs, hre) => {
    const tokenFactory = await hre.ethers.getContractFactory('BitTokenERC20')
    const token = await tokenFactory.deploy('name', 'symbol')
    await token.deployed();
    console.log("export BIT_ERC20_TOKEN=%s", token.address.toLocaleLowerCase());

    for (var i = 0; i < wallets.length; i++) {
      await token.connect(wallets[i]).mint(amount)
      // console.log("balance of", wallets[i].address, "is:", await token.balanceOf(wallets[i].address))
    }

    const sequencerFactory = await hre.ethers.getContractFactory('Sequencer')
    const sequencer = await sequencerFactory.deploy()
    await sequencer.deployed()
    await sequencer.initialize(token.address)
    console.log("export ENV_SEQUENCER_CONTRACT_ADDRESS=%s", sequencer.address.toLocaleLowerCase());
    console.log("sequencer contract owner :", await sequencer.owner())
    console.log("sequencer bit address :", await sequencer.bitToken())

    // create sequencer
    for (var i = 1; i < wallets.length; i++) {
      await token.connect(wallets[i]).approve(sequencer.address, deposit)
      await sequencer.connect(wallets[i]).createSequencer(deposit, wallets[i].address, nodeIDs[i])
    }
    await sequencer.updateSequencerLimit(100)

    await sequencer.updateScheduler(wallets[0].address)
    console.log("scheduler is :",await sequencer.scheduler())
  });

task("updateScheduler")
  .addParam("sequencer")
  .addParam("scheduler")
  .setAction(async (taskArgs, hre) => {
    const sequencerFactory = await hre.ethers.getContractFactory('Sequencer')
    const sequencer = sequencerFactory.attach(taskArgs.sequencer)
    await sequencer.updateScheduler(taskArgs.scheduler)
    console.log(await sequencer.scheduler())
  });


task("init")
  // .addParam("token")
  .setAction(async (taskArgs, hre) => {
    const tokenFactory = await hre.ethers.getContractFactory('BitTokenERC20')
    const token = tokenFactory.attach("0xd778E3b2Ab28dd5D7a46e3bCF4D6d8e6eb716B57")
    // 0xd778E3b2Ab28dd5D7a46e3bCF4D6d8e6eb716B57
    for (var i = 0; i < wallets.length; i++) {

     var res  = await token.connect(wallets[i]).mint(amount)
     console.log("res",res,"address:",wallets[i].addres)
      // console.log("balance of", wallets[i].address, "is:", await token.balanceOf(wallets[i].address))
    }
  });