import "@typechain/hardhat";
import "hardhat-gas-reporter";
import "hardhat-abi-exporter";
import "solidity-coverage";
import "@nomiclabs/hardhat-ethers"
import "@openzeppelin/hardhat-upgrades"
import "./tasks/Sequencer";

module.exports = {
  defaultNetwork: 'hardhat',
  defender: {
      apiKey: "[apiKey]",
      apiSecret: "[apiSecret]",
  },
  abiExporter: {
      path: './abi',
      runOnCompile: true,
      clear: true,
  },
  networks: {
      hardhat: {
          allowUnlimitedContractSize: true,
      },
      l1: {
        url: "http://localhost:9545",
        chainId: 31337,
        gas: 'auto',
        gasPrice: 'auto',
      },
      l2: {
          url: "http://localhost:8545",
          chainId: 17,
          gas: 'auto',
          gasPrice: 'auto',
          accounts :['ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'],
      }
  },
  solidity: {
      version: '0.8.12',
      settings: {
          optimizer: {
              enabled: true,
              runs: 1000,
          },
      }
  },
  gasReporter: {
      enabled: true,
      showMethodSig: true,
      maxMethodDiff: 10,
  },
  contractSizer: {
      alphaSort: true,
      runOnCompile: true,
      disambiguatePaths: false,
  },
  paths: {
      sources: "./contracts",
      tests: "./test",
      cache: "./cache",
      artifacts: "./artifacts"
  }
}