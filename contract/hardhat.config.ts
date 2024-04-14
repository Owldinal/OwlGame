import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import * as dotenv from "dotenv";

dotenv.config();

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.24",
    settings: {
      optimizer: {
        enabled: false,
        runs: 200
      }
    }
  },
  sourcify: {
    enabled: false
  },

  networks: {
    localhost: {
      url: 'http://127.0.0.1:8545',
      accounts: [process.env.HARDHAT_PRIVATE_KEY || '',
      process.env.HARDHAT_PRIVATE_KEY_1!,
      process.env.HARDHAT_PRIVATE_KEY_2!,
      process.env.HARDHAT_PRIVATE_KEY_3!,
      process.env.HARDHAT_PRIVATE_KEY_4!,
      process.env.HARDHAT_PRIVATE_KEY_5!,]
    },
    ganache: {
      url: 'http://127.0.0.1:7545',
      // url: 'https://api.owldinal.xyz/node',
      accounts: { mnemonic: process.env.GANACHE_MNEMONIC }
    },
    merlinTestnet: {
      url: 'https://merlin-testnet.blockpi.network/v1/rpc/f1f602d890e8a1ccce4cc679aa4c49def500fbb2',
      timeout: 20000000,
      accounts: [process.env.MERLIN_TEST_PRIVATE_KEY || ''],
    },
    merlinMainnet: {
      url: 'https://rpc.merlinchain.io',
      timeout: 200000000,
      gasPrice: 50000000,
      accounts: [process.env.MERLIN_MAIN_PRIVATE_KEY || ''],
    },
    sepolia: {
      url: `https://sepolia.infura.io/v3/${process.env.SEPOLIA_INFURA_KEY}`,
      timeout: 200000000,
      accounts: [process.env.MERLIN_TEST_PRIVATE_KEY || '',]
    }
  },
  etherscan: {
    apiKey: {
      merlinTestnet: "no-api-key-needed",
      merlinMainnet: "no-api-key-needed",
      sepolia: process.env.SEPOLIA_ETHERSCAN_API_KEY!,
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
          apiURL: "https://scan-v1.merlinchain.io/api",
          browserURL: "https://scan-v1.merlinchain.io"
        }
      }
    ]
  }
};

export default config;
