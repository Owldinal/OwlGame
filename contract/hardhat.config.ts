import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import * as dotenv from "dotenv";

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.24",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200
      }
    }
  },
  sourcify: {
    enabled: false
  },
  defaultNetwork: "localhost",

  networks: {
    hardhat: {
      accounts: { mnemonic: process.env.GANACHE_MNEMONIC }
    },
    localhost: {
      url: 'http://127.0.0.1:8545',
      accounts: { mnemonic: process.env.GANACHE_MNEMONIC }
    },
    ganache: {
      url: 'https://node.owldinal.xyz',
      accounts: { mnemonic: process.env.GANACHE_MNEMONIC, count: 300 }
    },
    merlinTestnet: {
      url: 'https://merlin-testnet.blockpi.network/v1/rpc/f1f602d890e8a1ccce4cc679aa4c49def500fbb2',
      timeout: 20000000,
      accounts: [process.env.MERLIN_TEST_PRIVATE_KEY || ''],
    },
    merlinMainnet: {
      url: 'https://rpc.merlinchain.io',
      accounts: [process.env.MERLIN_MAIN_PRIVATE_KEY || ''],
    }
  },
  etherscan: {
    apiKey: {
      merlinTestnet: "no-api-key-needed",
      merlinMainnet: "no-api-key-needed"
    },
    customChains: [
      {
        network: "merlinTestnet",
        chainId: 686868,
        urls: {
          apiURL: "https://testnet-scan-v1.merlinchain.io/api",
          browserURL: "https://testnet-scan-v1.merlinchain.io"
        }
      },
      {
        network: "merlinMainnet",
        chainId: 4200,
        urls: {
          apiURL: "https://scan.merlinchain.io/api",
          browserURL: "https://scan-v1.merlinchain.io"
        }
      }
    ]
  }
};

export default config;
