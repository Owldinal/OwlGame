const hre = require("hardhat");
const { deploy } = require('./utils');

let deployer, ownerAddress, backendAddress;

// deploy contracts
let owlTokenContract, owlTokenAddress;
let owldinalNftContract, owldinalNftAddress;
let owlGameContract, owlGameAddress;
let genOneBoxContract, genOneBoxAddress;

async function main() {
	[deployer] = await hre.ethers.getSigners();
	ownerAddress = deployer.address;
	backendAddress = process.env.BACKEND_WALLET;

	owlTokenAddress = "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6";
	owldinalNftAddress = "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318";
	owlGameAddress = "0x49fd2BE640DB2910c2fAb69bB8531Ab6E76127ff";
	genOneBoxAddress = "0x4631BCAbD6dF18D94796344963cB60d44a4136b6";

	await deployOrConnect();

	// init contract config
	const operations = [
		await owlGameContract.connect(deployer).initialize(owlTokenAddress, owldinalNftAddress, genOneBoxAddress),
	];

	await Promise.all(operations).then((results) => Promise.all(results.map(result => result.wait())));

	console.log(`\nconst [owlTokenAddress, owldinalNftAddress, genOneBoxAddress, owlGameAddress] = ["${owlTokenAddress}", "${owldinalNftAddress}", "${genOneBoxAddress}", "${owlGameAddress}"];\n`);

	await owlGameContract.connect(deployer).unstakeOwldinalNft([2, 3, 4]);

	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 2);
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 3);
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 4);
	const tokenIdList = await owlGameContract.connect(deployer).stakeOwldinalNft([2, 3, 4]);
	console.log(`stake owldinal nft: ${JSON.stringify(tokenIdList)}`);
	
	
}

async function deployOrConnect() {
	if (owlTokenAddress) {
		owlTokenContract = await hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddress));
		console.log(`OwlToken contract connected to : ${owlTokenAddress}`);
	} else {
		const params = [ownerAddress];
		[owlTokenContract, owlTokenAddress] = await deploy("OwlToken", params);
		console.log(`OwlToken contract deployed to : ${owlTokenAddress}\nParams = ${params.join(" ")}`);

		await owlTokenContract.connect(deployer).mint(deployer.address, BigInt(10000000 * 10 ** 18));
	}

	if (owldinalNftAddress) {
		owldinalNftContract = await hre.ethers.getContractFactory("Owldinal").then((c) => c.attach(owldinalNftAddress));
		console.log(`Owldinal contract connected to : ${owldinalNftAddress}`);
	} else {
		const params = [10000000, owlTokenAddress, deployer.address, backendAddress];
		[owldinalNftContract, owldinalNftAddress] = await deploy("Owldinal", params);
		console.log(`Owldinal contract deployed to : ${owldinalNftAddress}\nParams = ${params.join(" ")}`);

		await owldinalNftContract.connect(deployer).mintByAdmin();
		await owldinalNftContract.connect(deployer).mintByAdmin();
		await owldinalNftContract.connect(deployer).mintByAdmin();
	}

	if (owlGameAddress) {
		owlGameContract = await hre.ethers.getContractFactory("OwlGame").then((c) => c.attach(owlGameAddress));
		console.log(`OwlGame contract connected to : ${owlGameAddress}`);
	} else {
		const params = [];
		[owlGameContract, owlGameAddress] = await deploy("OwlGame", params);
		console.log(`OwlGame contract deployed to : ${owlGameAddress}\nParams = ${params.join(" ")}`);
	}

	if (genOneBoxAddress) {
		genOneBoxContract = await hre.ethers.getContractFactory("OwldinalGenOneBox").then((c) => c.attach(genOneBoxAddress));
		console.log(`OwldinalGenOneBox contract connected to : ${genOneBoxAddress}`);
	} else {
		const params = [owlTokenAddress, owlGameAddress];
		[genOneBoxContract, genOneBoxAddress] = await deploy("OwldinalGenOneBox", params);
		console.log(`OwldinalGenOneBox contract deployed to : ${genOneBoxAddress}\nParams = ${params.join(" ")}`);
	}

}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});