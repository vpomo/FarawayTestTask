import * as dotenv from "dotenv";
import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "solidity-coverage"
import "@nomiclabs/hardhat-etherscan";
import "@openzeppelin/hardhat-upgrades";

dotenv.config();

const TEST_MNEMONIC = "test test test test test test test test test test test junk";
const TEST_ACCOUNT = { mnemonic: TEST_MNEMONIC, }


task("accounts", "Prints the list of accounts", async (taskArgs, hre) => {
    const accounts = await hre.ethers.getSigners();

    for (const account of accounts) {
        console.log(account.address);
    }
});

const accounts =
    process.env.DEPLOYER_PRIVATE_KEY !== undefined
        ? [process.env.DEPLOYER_PRIVATE_KEY]
        : [];

const config: HardhatUserConfig = {
  defaultNetwork: "hardhat",
  solidity: {
    compilers: [
      {
        version: "0.8.20",
        settings: {
          viaIR: true,
          optimizer: {
            enabled: true,
            runs: 200,
          },
        }
      }
    ],
    overrides: {},
  },
  contractSizer: {
    alphaSort: true,
    runOnCompile: true,
    disambiguatePaths: false,
  },
  networks: {
    hardhat: {
      forking: {
        url: "https://polygon-mainnet.g.alchemy.com/v2/" + (process.env.ALCHEMY_API_KEY || ''),
        blockNumber: 41447111,
      },
    },
    mainnet: {
      url: process.env.MAINNET_URI || '',
      accounts: process.env.MAINNET_PRIVATE_KEY ? [process.env.MAINNET_PRIVATE_KEY] : TEST_ACCOUNT,
    },
    goerli: {
      url: process.env.GOERLI_URI || '',
      accounts: process.env.GOERLI_PRIVATE_KEY ? [process.env.GOERLI_PRIVATE_KEY] : TEST_ACCOUNT,
    },
    mumbai: {
      url: process.env.MUMBAI_URI || '',
      accounts: process.env.DEPLOYER_PRIVATE_KEY ? [process.env.DEPLOYER_PRIVATE_KEY] : TEST_ACCOUNT,
    },
    fork: {
      url: process.env.FORK_URI || '',
      accounts: process.env.FORK_PRIVATE_KEY ? [process.env.FORK_PRIVATE_KEY] : TEST_ACCOUNT,
    },
  },
  mocha: {
    timeout: 0
  },
  etherscan: {
    apiKey: process.env.POLYGONSCAN_API_KEY || ''
  },
  typechain: {
    outDir: "typechain",
    target: "ethers-v5"
  },
  gasReporter: {
    enabled: true
  },
  docgen: {
    outputDir: 'docs',
    pages: 'files',
    exclude: [
      './interfaces',
      './oz',
      './test',
      './utils'
    ]
  }
};

export default config;
