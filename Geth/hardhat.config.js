/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.17",
  networks: {
    hardhat: {
      gas: "auto",
      allowUnlimitedContractSize: true, // TODO: Remove this when we reduce the size of the verifier contracts
      blockGasLimit: 80000000, // TODO: Remove this when we reduce the size of the verifier contracts
      mining: {
        auto: true,
        interval: 5000,
        mempool: {
          order: "fifo",
        },
      },
    },
  },
};
