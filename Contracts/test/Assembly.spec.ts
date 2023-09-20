import { Signer, Contract, BigNumber, utils, constants } from 'ethers'
import hre from 'hardhat'

const { expect } = require("chai")
const { ethers } = require("hardhat")

describe("Assemply", () => {
    let accounts: Signer[]
    let contract: Contract
    let assembly: Contract

    before(async () => {
        accounts = await ethers.getSigners()
        await deployERC20()
        await deployAssembly()
    });

    const deployERC20 = async () => {
        const Factory = await hre.ethers.getContractFactory('TokenERC20')
        contract = await Factory.deploy("TEST", "TEST")
        await contract.deployed()
    }

    const deployAssembly = async () => {
        const Factory = await hre.ethers.getContractFactory('AssemblyLib')
        assembly = await Factory.deploy()
        await assembly.deployed()
    }

    it('Test at', async () => {
        const o_code = await assembly.at(contract.address)
        console.log(o_code)
    })

    it('Test add', async () => {
        const res = await assembly.add()
        console.log(res)
    })

    it('Test extractBlockNumber', async () => {
        const bkWitness = '0x0000000000000000000000000000000000000000000000000000000000000065'
        const res = await assembly.extractBlockNumber(bkWitness)
        console.log(res)
    })
});
