const hre = require("hardhat");
require("dotenv").config();

async function main() {
	const [deployer] = await hre.ethers.getSigners();
	const ownerAddress = deployer.address;

	const owlTokenParams = [ownerAddress];
	const [owlTokenContract, owlTokenAddresss] = await deploy("OwlToken", owlTokenParams);
	console.log(`OwlToken contract deployed to : ${owlTokenAddresss}\nParams = ${owlTokenParams}`);

	const backendAddress = process.env.BACKEND_WALLET;
	const boxGen0Params = [3800000, owlTokenAddresss, deployer.address, backendAddress];
	const [boxGen0Contract, boxGen0Address] = await deploy("MysteryBoxGen0", boxGen0Params);
	console.log(`BoxGen0 contract deployed to : ${boxGen0Address}\nParams = ${boxGen0Params}`);

	// const owlTokenAddr = "0xf7d82CEA33a4D7C2F45e6Fd86EDB5b110FE70139";
	// const backendAddress = "0x047BE992fBB36702b3684ccB91a5ec4B0e2eA042";
	// const boxGen0Addr = "0xE82671921e9f6da6F11f4Fb616860d6CD342E790";

	// const [
	// 	owlTokenContract,
	// 	boxGen0Contract,
	// ] = await Promise.all(
	// 	[
	// 		hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddr)),
	// 		hre.ethers.getContractFactory("MysteryBoxGen0").then((c) => c.attach(boxGen0Addr)),
	// 	]
	// );

	// init contract config
	const operations = [
		await owlTokenContract.connect(deployer).mint(
			deployer.address, BigInt(1000 * 10 ** 18)
		),
	];

	await Promise.all(operations).then((results) => Promise.all(results.map(result => result.wait())));

	console.log(`\nconst [owlTokenAddr, boxGen0Addr] = ["${owlTokenAddresss}", "${boxGen0Address}"];\n`);
}

async function deploy(name, args) {
	const contract = await hre.ethers.deployContract(name, args).then((c) => c.waitForDeployment());
	const addr = await contract.getAddress();
	return [contract, addr];
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});