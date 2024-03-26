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
    merlinTestnet: {
      url: 'https://testnet-rpc.merlinchain.io',
      timeout: 20000000,
      accounts: [process.env.MERLIN_TEST_PRIVATE_KEY || ''],
    },
    merlinMainnet: {
      url: 'https://rpc.merlinchain.io',
      timeout: 200000000,
      gasPrice: 50000000,
      accounts: [process.env.MERLIN_TEST_PRIVATE_KEY || ''],
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
          apiURL: "https://testnet-scan.merlinchain.io/api",
          browserURL: "https://testnet-scan.merlinchain.io"
        }
      },
      {
        network: "merlinMainnet",
        chainId: 4200,
        urls: {
          apiURL: "https://scan.merlinchain.io/api",
          browserURL: "https://scan.merlinchain.io"
        }
      }
    ]
  }
};

export default config;
