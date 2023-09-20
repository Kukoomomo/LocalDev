import { Signer, Contract } from 'ethers'
import hre from 'hardhat'

const { expect } = require("chai")
const { ethers } = require("hardhat")

describe("Assemply", () => {
    let accounts: Signer[]
    let zkevmConfig: Contract

    before(async () => {
        accounts = await ethers.getSigners()
        await deployZKEVMConfig()

    });

    const deployZKEVMConfig = async () => {
        const Factory = await hre.ethers.getContractFactory('ZKEVMConfig')
        zkevmConfig = await Factory.deploy()
        await zkevmConfig.deployed()
    }
    // 
    it('Test applyL1ToL2Alias', async () => {
        const addr = '0x6900000000000000000000000000000000000002'
        const res = await zkevmConfig.applyL1ToL2Alias(addr)
        console.log(res)
    })

    it('Test extractBlockNumber', async () => {
        const data = '0x00000000000000000000000000000000000000000000000000000000000000030000000064d5c10c00000000000000000000000000000000000000000000000000000000000000000000000001c86c88000000000000000000000000000000000000000000000000000000000000000000000000000000040000000064d5c11200000000000000000000000000000000000000000000000000000000000000000000000001c7fa6e000000000000000000000000000000000000000000000000000000000000000000000000000000050000000064d5c11800000000000000000000000000000000000000000000000000000000000000000000000001c78871000000000000000000000000000000000000000000000000000000000000000000000000000000060000000064d5c11e00000000000000000000000000000000000000000000000000000000000000000000000001c71690000000000000000000000000000000000000000000000000000000000000000000000000000000070000000064d5c12400000000000000000000000000000000000000000000000000000000000000000000000001c6a4cc000000000000000000000000000000000000000000000000000000000000000000000000000000080000000064d5c12a00000000000000000000000000000000000000000000000000000000000000000000000001c63324000000000000000000000000000000000000000000000000000000000000000000000000000000090000000064d5c13000000000000000000000000000000000000000000000000000000000000000000000000001c5c1990000000000000000000000000000000000000000000000000000000000000000000000000000000a0000000064d5c13600000000000000000000000000000000000000000000000000000000000000000000000001c5502a0000000000000000000000000000000000000000000000000000000000000000000000000000000b0000000064d5c13c00000000000000000000000000000000000000000000000000000000000000000000000001c4ded70000000000000000000000000000000000000000000000000000000000000000000000000000000c0000000064d5c14200000000000000000000000000000000000000000000000000000000000000000000000001c46da10000000000000000000000000000000000000000000000000000000000000000000000000000000d0000000064d5c14800000000000000000000000000000000000000000000000000000000000000000000000001c3fc870000000000000000000000000000000000000000000000000000000000000000000000000000000e0000000064d5c14e00000000000000000000000000000000000000000000000000000000000000000000000001c38b890000000000000000000000000000000000000000000000000000000000000000000000000000000f0000000064d5c15400000000000000000000000000000000000000000000000000000000000000000000000001c31aa8000000000000000000000000000000000000000000000000000000000000000000000000000000100000000064d5c15a00000000000000000000000000000000000000000000000000000000000000000000000001c2a9e3000000000000000000000000000000000000000000000000000000000000000000000000000000110000000064d5c16000000000000000000000000000000000000000000000000000000000000000000000000001c2393a000000000000000000000000000000000000000000000000000000000000000000000000000000120000000064d5c16600000000000000000000000000000000000000000000000000000000000000000000000001c1c8ad000000000000000000000000000000000000000000000000000000000000000000000000000000130000000064d5c16c00000000000000000000000000000000000000000000000000000000000000000000000001c1583c000000000000000000000000000000000000000000000000000000000000000000000000000000140000000064d5c17300000000000000000000000000000000000000000000000000000000000000000000000001c0e7e7000000000000000000000000000000000000000000000000000000000000000000000000000000150000000064d5c17900000000000000000000000000000000000000000000000000000000000000000000000001c077af000000000000000000000000000000000000000000000000000000000000000000000000000000160000000064d5c17f00000000000000000000000000000000000000000000000000000000000000000000000001c00793000000000000000000000000000000000000000000000000000000000000000000000000000000170000000064d5c18500000000000000000000000000000000000000000000000000000000000000000000000001bf9793000000000000000000000000000000000000000000000000000000000000000000000000000000180000000064d5c18b00000000000000000000000000000000000000000000000000000000000000000000000001bf27af000000000000000000000000000000000000000000000000000000000000000000000000000000190000000064d5c19100000000000000000000000000000000000000000000000000000000000000000000000001beb7e70000000000000000000000000000000000000000000000000000000000000000000000000000001a0000000064d5c19700000000000000000000000000000000000000000000000000000000000000000000000001be483b0000000000000000000000000000000000000000000000000000000000000000000000000000001b0000000064d5c19d00000000000000000000000000000000000000000000000000000000000000000000000001bdd8aa0000000000000000000000000000000000000000000000000000000000000000000000000000001c0000000064d5c1a300000000000000000000000000000000000000000000000000000000000000000000000001bd69350000000000000000000000000000000000000000000000000000000000000000000000000000001d0000000064d5c1a900000000000000000000000000000000000000000000000000000000000000000000000001bcf9dc0000000000000000000000000000000000000000000000000000000000000000000000000000001e0000000064d5c1af00000000000000000000000000000000000000000000000000000000000000000000000001bc8a9f0000000000000000000000000000000000000000000000000000000000000000000000000000001f0000000064d5c1b500000000000000000000000000000000000000000000000000000000000000000000000001bc1b7e000000000000000000000000000000000000000000000000000000000000000000000000000000200000000064d5c1bb00000000000000000000000000000000000000000000000000000000000000000000000001bbac79000000000000000000000000000000000000000000000000000000000000000000000000000000210000000064d5c1c100000000000000000000000000000000000000000000000000000000000000000000000001bb3d8f000000000000000000000000000000000000000000000000000000000000000000000000000000220000000064d5c1c700000000000000000000000000000000000000000000000000000000000000000000000001bacec1000000000000000000000000000000000000000000000000000000000000000000000000000000230000000064d5c1cd00000000000000000000000000000000000000000000000000000000000000000000000001ba600f000000000000000000000000000000000000000000000000000000000000000000000000000000240000000064d5c1d300000000000000000000000000000000000000000000000000000000000000000000000001b9f178000000000000000000000000000000000000000000000000000000000000000000000000000000250000000064d5c1d900000000000000000000000000000000000000000000000000000000000000000000000001b982fd000000000000000000000000000000000000000000000000000000000000000000000000000000260000000064d5c1df00000000000000000000000000000000000000000000000000000000000000000000000001b9149e000000000000000000000000000000000000000000000000000000000000000000000000000000270000000064d5c1e500000000000000000000000000000000000000000000000000000000000000000000000001b8a65a000000000000000000000000000000000000000000000000000000000000000000000000000000280000000064d5c1eb00000000000000000000000000000000000000000000000000000000000000000000000001b83832000000000000000000000000000000000000000000000000000000000000000000000000000000290000000064d5c1f100000000000000000000000000000000000000000000000000000000000000000000000001b7ca250000000000000000000000000000000000000000000000000000000000000000000000000000002a0000000064d5c1f700000000000000000000000000000000000000000000000000000000000000000000000001b75c340000000000000000000000000000000000000000000000000000000000000000000000000000002b0000000064d5c1fd00000000000000000000000000000000000000000000000000000000000000000000000001b6ee5e0000000000000000000000000000000000000000000000000000000000000000000000000000002c0000000064d5c20300000000000000000000000000000000000000000000000000000000000000000000000001b680a40000000000000000000000000000000000000000000000000000000000000000000000000000002d0000000064d5c20900000000000000000000000000000000000000000000000000000000000000000000000001b613050000000000000000000000000000000000000000000000000000000000000000000000000000002e0000000064d5c20f00000000000000000000000000000000000000000000000000000000000000000000000001b5a5820000000000000000000000000000000000000000000000000000000000000000000000000000002f0000000064d5c21500000000000000000000000000000000000000000000000000000000000000000000000001b5381a000000000000000000000000000000000000000000000000000000000000000000000000000000300000000064d5c21b00000000000000000000000000000000000000000000000000000000000000000000000001b4cacd000000000000000000000000000000000000000000000000000000000000000000000000000000310000000064d5c22100000000000000000000000000000000000000000000000000000000000000000000000001b45d9c000000000000000000000000000000000000000000000000000000000000000000000000000000320000000064d5c22700000000000000000000000000000000000000000000000000000000000000000000000001b3f086000000000000000000000000000000000000000000000000000000000000000000000000000000330000000064d5c22d00000000000000000000000000000000000000000000000000000000000000000000000001b3838b000000000000000000000000000000000000000000000000000000000000000000000000000000340000000064d5c23300000000000000000000000000000000000000000000000000000000000000000000000001b316ac000000000000000000000000000000000000000000000000000000000000000000000000000000350000000064d5c23900000000000000000000000000000000000000000000000000000000000000000000000001b2a9e8000000000000000000000000000000000000000000000000000000000000000000000000000000360000000064d5c23f00000000000000000000000000000000000000000000000000000000000000000000000001b23d3f000000000000000000000000000000000000000000000000000000000000000000000000000000370000000064d5c24500000000000000000000000000000000000000000000000000000000000000000000000001b1d0b1000000000000000000000000000000000000000000000000000000000000000000000000000000380000000064d5c24b00000000000000000000000000000000000000000000000000000000000000000000000001b1643e000000000000000000000000000000000000000000000000000000000000000000000000000000390000000064d5c25100000000000000000000000000000000000000000000000000000000000000000000000001b0f7e60000000000000000000000000000000000000000000000000000000000000000000000000000003a0000000064d5c25700000000000000000000000000000000000000000000000000000000000000000000000001b08baa0000000000000000000000000000000000000000000000000000000000000000000000000000003b0000000064d5c25d00000000000000000000000000000000000000000000000000000000000000000000000001b01f890000000000000000000000000000000000000000000000000000000000000000000000000000003c0000000064d5c26300000000000000000000000000000000000000000000000000000000000000000000000001afb3830000000000000000000000000000000000000000000000000000000000000000000000000000003d0000000064d5c26900000000000000000000000000000000000000000000000000000000000000000000000001af47980000000000000000000000000000000000000000000000000000000000000000000000000000003e0000000064d5c26f00000000000000000000000000000000000000000000000000000000000000000000000001aedbc80000000000000000000000000000000000000000000000000000000000000000000000000000003f0000000064d5c27600000000000000000000000000000000000000000000000000000000000000000000000001ae7013000000000000000000000000000000000000000000000000000000000000000000000000000000400000000064d5c27c00000000000000000000000000000000000000000000000000000000000000000000000001ae0478000000000000000000000000000000000000000000000000000000000000000000000000000000410000000064d5c28200000000000000000000000000000000000000000000000000000000000000000000000001ad98f8000000000000000000000000000000000000000000000000000000000000000000000000000000420000000064d5c28800000000000000000000000000000000000000000000000000000000000000000000000001ad2d93000000000000000000000000000000000000000000000000000000000000000000000000000000430000000064d5c28e00000000000000000000000000000000000000000000000000000000000000000000000001acc249000000000000000000000000000000000000000000000000000000000000000000000000000000440000000064d5c29400000000000000000000000000000000000000000000000000000000000000000000000001ac571a000000000000000000000000000000000000000000000000000000000000000000000000000000450000000064d5c29a00000000000000000000000000000000000000000000000000000000000000000000000001abec06000000000000000000000000000000000000000000000000000000000000000000000000000000460000000064d5c2a000000000000000000000000000000000000000000000000000000000000000000000000001ab810c000000000000000000000000000000000000000000000000000000000000000000000000000000470000000064d5c2a600000000000000000000000000000000000000000000000000000000000000000000000001ab162d000000000000000000000000000000000000000000000000000000000000000000000000000000480000000064d5c2ac00000000000000000000000000000000000000000000000000000000000000000000000001aaab69000000000000000000000000000000000000000000000000000000000000000000000000000000490000000064d5c2b200000000000000000000000000000000000000000000000000000000000000000000000001aa40c00000000000000000000000000000000000000000000000000000000000000000000000000000004a0000000064d5c2b800000000000000000000000000000000000000000000000000000000000000000000000001a9d6310000000000000000000000000000000000000000000000000000000000000000000000000000004b0000000064d5c2be00000000000000000000000000000000000000000000000000000000000000000000000001a96bbd0000000000000000000000000000000000000000000000000000000000000000000000000000004c0000000064d5c2c400000000000000000000000000000000000000000000000000000000000000000000000001a901640000000000000000000000000000000000000000000000000000000000000000000000000000004d0000000064d5c2ca00000000000000000000000000000000000000000000000000000000000000000000000001a897250000000000000000000000000000000000000000000000000000000000000000000000000000004e0000000064d5c2d000000000000000000000000000000000000000000000000000000000000000000000000001a82d010000000000000000000000000000000000000000000000000000000000000000000000000000004f0000000064d5c2d600000000000000000000000000000000000000000000000000000000000000000000000001a7c2f7000000000000000000000000000000000000000000000000000000000000000000000000000000500000000064d5c2dc00000000000000000000000000000000000000000000000000000000000000000000000001a75908000000000000000000000000000000000000000000000000000000000000000000000000000000510000000064d5c2e200000000000000000000000000000000000000000000000000000000000000000000000001a6ef33000000000000000000000000000000000000000000000000000000000000000000000000000000520000000064d5c2e800000000000000000000000000000000000000000000000000000000000000000000000001a68579000000000000000000000000000000000000000000000000000000000000000000000000000000530000000064d5c2ee00000000000000000000000000000000000000000000000000000000000000000000000001a61bd9000000000000000000000000000000000000000000000000000000000000000000000000000000540000000064d5c2f400000000000000000000000000000000000000000000000000000000000000000000000001a5b254000000000000000000000000000000000000000000000000000000000000000000000000000000550000000064d5c2fa00000000000000000000000000000000000000000000000000000000000000000000000001a548e9000000000000000000000000000000000000000000000000000000000000000000000000000000560000000064d5c30000000000000000000000000000000000000000000000000000000000000000000000000001a4df98000000000000000000000000000000000000000000000000000000000000000000000000000000570000000064d5c30600000000000000000000000000000000000000000000000000000000000000000000000001a47662000000000000000000000000000000000000000000000000000000000000000000000000000000580000000064d5c30c00000000000000000000000000000000000000000000000000000000000000000000000001a40d46000000000000000000000000000000000000000000000000000000000000000000000000000000590000000064d5c31200000000000000000000000000000000000000000000000000000000000000000000000001a3a4440000000000000000000000000000000000000000000000000000000000000000000000000000005a0000000064d5c31800000000000000000000000000000000000000000000000000000000000000000000000001a33b5c0000000000000000000000000000000000000000000000000000000000000000000000000000005b0000000064d5c31e00000000000000000000000000000000000000000000000000000000000000000000000001a2d28f0000000000000000000000000000000000000000000000000000000000000000000000000000005c0000000064d5c32400000000000000000000000000000000000000000000000000000000000000000000000001a269dc0000000000000000000000000000000000000000000000000000000000000000000000000000005d0000000064d5c32a00000000000000000000000000000000000000000000000000000000000000000000000001a201430000000000000000000000000000000000000000000000000000000000000000000000000000005e0000000064d5c33000000000000000000000000000000000000000000000000000000000000000000000000001a198c40000000000000000000000000000000000000000000000000000000000000000000000000000005f0000000064d5c33600000000000000000000000000000000000000000000000000000000000000000000000001a1305f000000000000000000000000000000000000000000000000000000000000000000000000000000600000000064d5c33c00000000000000000000000000000000000000000000000000000000000000000000000001a0c814000000000000000000000000000000000000000000000000000000000000000000000000000000610000000064d5c34200000000000000000000000000000000000000000000000000000000000000000000000001a05fe3000000000000000000000000000000000000000000000000000000000000000000000000000000620000000064d5c348000000000000000000000000000000000000000000000000000000000000000000000000019ff7cd000000000000000000000000000000000000000000000000000000000000000000000000000000630000000064d5c34e000000000000000000000000000000000000000000000000000000000000000000000000019f8fd1000000000000000000000000000000000000000000000000000000000000000000000000000000640000000064d5c355000000000000000000000000000000000000000000000000000000000000000000000000019f27ef000000000000000000000000000000000000000000000000000000000000000000000000000000650000000064d5c35b000000000000000000000000000000000000000000000000000000000000000000000000019ec0270000000000000000'
        const res = await zkevmConfig.decode(data)
        console.log(res)
    })
});