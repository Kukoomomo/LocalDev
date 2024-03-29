const { task } = require("hardhat/config");
const fs = require('fs');
const ethers = require('ethers');
const { exit } = require("process");

const emtpyHash = "0x0000000000000000000000000000000000000000000000000000000000000000";

async function main(hash, step, file) {
    const provider = new ethers.providers.JsonRpcProvider("http://localhost:8545");
    var res = await provider.send("debug_generateProofForTest", [hash, 0, 0, parseInt(step)]);
    fs.writeFileSync(file, JSON.stringify(res));
    console.log("wrote proof to " + file);
}

task("generateOsp", "generate osp")
    .addParam("hash", "the transaction hash to prove")
    .addParam("step", "the step to prove")
    .addParam("file", "where to save the proof")
    .setAction(async (taskArgs) => {
        await main(taskArgs.hash, taskArgs.step, taskArgs.file);
    });

module.exports = {};
