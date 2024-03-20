import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import * as dotenv from "dotenv";

dotenv.config();

const config: HardhatUserConfig = {
  solidity: "0.8.24",
  networks: {
    localhost: {
      url: 'http://127.0.0.1:7545',
      accounts: [process.env.LOCALHOST_PRIVATE_KEY || '']
    },
    hardhatLocal: {
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
      timeout: 2000000,
      accounts: [process.env.MERLIN_TEST_PRIVATE_KEY || '',
        "0x3e9cebd0c8d24c22e77f0e276161b414356489300fdfc9dd53c94a06dab09a40"],
    },
    sepolia: {
      url: 'https://sepolia.drpc.org',
      timeout: 2000000,
      accounts: [process.env.SEPOLIA_PRIVATE_KEY || '',]
    }
  }
};

export default config;
