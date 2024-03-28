const hre = require("hardhat");
require("dotenv").config();
const { deploy } = require('./utils');

async function main() {
	const [deployer] = await hre.ethers.getSigners();
	const ownerAddress = deployer.address;
	const backendAddress = process.env.BACKEND_WALLET;

	let owlTokenAddress, boxGen0Address;
	let owlTokenContract, boxGen0Contract;

	owlTokenAddress = "0x5cf806FE094c457D993c89B2A3d599d50d4cbBF4";
	// boxGen0Address = "0xE82671921e9f6da6F11f4Fb616860d6CD342E790";

	if (owlTokenAddress) {
		owlTokenContract = await hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddress));
		console.log(`OwlToken contract connected to : ${owlTokenAddress}`);
	} else {
		const owlTokenParams = [ownerAddress];
		[owlTokenContract, owlTokenAddress] = await deploy("OwlToken", owlTokenParams);
		console.log(`OwlToken contract deployed to : ${owlTokenAddress}\nParams = ${owlTokenParams.join(" ")}`);
	}

	if (boxGen0Address) {
		boxGen0Contract = await hre.ethers.getContractFactory("MysteryBoxGen0").then((c) => c.attach(boxGen0Address));
		console.log(`BoxGen0 contract connected to : ${owlTokenAddress}`);
	} else {
		const boxGen0Params = [10000000, owlTokenAddress, deployer.address, backendAddress];
		[boxGen0Contract, boxGen0Address] = await deploy("MysteryBoxGen0", boxGen0Params);
		console.log(`BoxGen0 contract deployed to : ${boxGen0Address}\nParams = ${boxGen0Params.join(" ")}`);
		console.log(`npx hardhat verify --network merlinTestnet ${boxGen0Address} ${boxGen0Params.join(" ")}`);
	}

	// init contract config
	const operations = [
		await owlTokenContract.connect(deployer).mint(
			deployer.address, BigInt(1000 * 10 ** 18)
		),
	];

	// await Promise.all(operations).then((results) => Promise.all(results.map(result => result.wait())));

	console.log("\n----\n");
	console.log("Copy to gen0-test: \n");
	console.log(`owlTokenAddress = "${owlTokenAddress}";`);
	console.log(`boxGen0Address = "${boxGen0Address}";`);
	console.log("\n----\n");
	console.log("Copy to backend .env: \n");
	console.log(`NFT_OWL_ADDR=${boxGen0Address}`);
	console.log("\n----\n");
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});