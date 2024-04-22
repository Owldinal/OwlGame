const hre = require("hardhat");
const { deploy } = require('./utils');

let deployer, ownerAddress;

// deploy contracts
let owlTokenContract, owlTokenAddress;
let owldinalNftContract, owldinalNftAddress;
let owlGameContract, owlGameAddress;
let genOneBoxContract, genOneBoxAddress;

const decimal = BigInt(10) ** BigInt(18)

async function main() {
	[deployer] = await hre.ethers.getSigners();
	ownerAddress = deployer.address;
	backendAddress = process.env.BACKEND_WALLET;
	let tx;
	console.log(`Address = ${deployer.address}`);

	const blockNumber = await hre.ethers.provider.getBlockNumber();
	console.log(`Current Block number: ${blockNumber}`);

	owldinalNftAddress = "0x3cb8094fb209d21Aab2E474e91476EE9EEc332C9";
	owlTokenAddress = "0xf254dcD9e270B8E0B92A68e197BE7C6355e8d87b";
	genOneBoxAddress = "0x850DA4D6cC3d1EdADA39C48d42B4b664d51F3c4d";
	owlGameAddress = "0xf0d330DB633b58446c895ff957a355FEC45b9511";

	await deployOrConnect();

	console.log(`
owldinalNftAddress = "${owldinalNftAddress}";
owlTokenAddress = "${owlTokenAddress}";
genOneBoxAddress = "${genOneBoxAddress}";
owlGameAddress = "${owlGameAddress}";
		`);

	console.log(`
EVENT_START_BLOCK=${blockNumber}
NFT_OWL_ADDR=${owldinalNftAddress}
OWL_TOKEN_ADDR=${owlTokenAddress}
NFT_MYSTERY_BOX_ADDR=${genOneBoxAddress}
OWL_GAME_ADDR=${owlGameAddress}
			`);

	await owlGameContract.connect(deployer).initialize(owlTokenAddress, owldinalNftAddress, genOneBoxAddress);
	await owlGameContract.connect(deployer).setMoonBoost(true);
	await genOneBoxContract.connect(deployer).addTransferWhiteList([
		"0xb78EA3993200e1e241A4c0670a89cFfDFB5CD560",
		"0xf3EB4f8d15cd76bA3130806Cc7ddE1EE4b8f6e42",
	]);

	const prizeAmount = BigInt(6_0000_0000n) * decimal;
	await owlTokenContract.connect(deployer).mint(ownerAddress, prizeAmount);
	console.log(`mint prizeAmount success`);
	await owlTokenContract.connect(deployer).approve(owlGameAddress, prizeAmount);
	console.log(`approve prizeAmount success`);
	await owlGameContract.connect(deployer).addPrize(prizeAmount);
	console.log(`prizeAmount = ${prizeAmount}`);

	return;
}

async function deployOrConnect() {
	if (owlTokenAddress) {
		owlTokenContract = await hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddress));
		console.log(`OwlToken contract connected to : ${owlTokenAddress}`);
	} else {
		const params = [ownerAddress];
		[owlTokenContract, owlTokenAddress] = await deploy("OwlToken", params);
		console.log(`OwlToken contract deployed to : ${owlTokenAddress}\nParams = ${params.join(" ")}`);
	}


	if (owldinalNftAddress) {
		owldinalNftContract = await hre.ethers.getContractFactory("Owldinal").then((c) => c.attach(owldinalNftAddress));
		console.log(`Owldinal contract connected to : ${owldinalNftAddress}`);
	} else {
		const params = [10000000, owlTokenAddress, deployer.address, backendAddress];
		[owldinalNftContract, owldinalNftAddress] = await deploy("Owldinal", params);
		console.log(`Owldinal contract deployed to : ${owldinalNftAddress}\nParams = ${params.join(" ")}`);
	}

	if (owlGameAddress) {
		owlGameContract = await hre.ethers.getContractFactory("OwlGame").then((c) => c.attach(owlGameAddress));
		console.log(`OwlGame contract connected to : ${owlGameAddress}`);
	} else {
		const params = [backendAddress];
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

function decodeInviteCode(encoded) {
	let inviteCode = '';
	for (let i = 0; i < 5; i++) {
		const charValue = Number((encoded >> (BigInt(i) * 5n)) & 0x1Fn);
		inviteCode += String.fromCharCode(charValue + 0x41);
	}
	return inviteCode;
}

function encodeInviteCode(inviteCode) {
	if (inviteCode.length !== 5) {
		throw new Error("Invalid invite code length");
	}
	let encoded = 0;
	for (let i = 0; i < 5; i++) {
		let char = inviteCode[i];
		let charValue = char.charCodeAt(0) - 0x41;
		encoded |= (charValue << (i * 5));
	}
	return encoded;
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});