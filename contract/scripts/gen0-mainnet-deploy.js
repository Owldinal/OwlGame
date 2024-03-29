const hre = require("hardhat");
require("dotenv").config();
const { deploy } = require('./utils');

async function main() {
	const [deployer] = await hre.ethers.getSigners();
	const ownerAddress = deployer.address;
	const backendAddress = process.env.BACKEND_WALLET;
	const voyaAddress = "0x480E158395cC5b41e5584347c495584cA2cAf78d";
	let boxGen0Address, boxGen0Contract;

	const tokenOneAddress = "0x83498c8EAa9DF19B9Bf0Ed26AbD820cf22F79842";
	const boxGen0Params = [8150000, voyaAddress, tokenOneAddress, backendAddress];
	[boxGen0Contract, boxGen0Address] = await deploy("Owldinal", boxGen0Params);
	console.log(`BoxGen0 contract deployed to : ${boxGen0Address}\nParams = ${boxGen0Params.join(" ")}`);
	console.log(`npx hardhat verify --network merlinMainnet ${boxGen0Address} ${boxGen0Params.join(" ")}`);

	console.log("\n----\n");
	console.log("Copy to backend .env: \n");
	console.log(`NFT_OWL_ADDR=${boxGen0Address}`);
	console.log("\n----\n");
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});