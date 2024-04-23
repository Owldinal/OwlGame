const hre = require("hardhat");
const { deploy, printTxDetail } = require('./utils');

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

	owldinalNftAddress = "0x6b18e87beb44a72eb48da76a881f9104cb97a180";
	owlTokenAddress = "0x62e99191071Fc1C5947CF1e21Aa95708dcc51AdB";
	// genOneBoxAddress = "0xdc34339327122f9f862e0553Dd35156cF78f6a03";
	// owlGameAddress = "0x0F85ef658068Da66908d6435aA93D1f4Db8b9937";

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
	// await owlGameContract.connect(deployer).setMoonBoost(true);
	// await genOneBoxContract.connect(deployer).addTransferWhiteList([
	// 	"0xb78EA3993200e1e241A4c0670a89cFfDFB5CD560",
	// 	"0xf3EB4f8d15cd76bA3130806Cc7ddE1EE4b8f6e42",
	// ]);

	// const prizeAmount = BigInt(6_0000_0000n) * decimal;
	// await owlTokenContract.connect(deployer).mint(ownerAddress, prizeAmount);
	// console.log(`mint prizeAmount success`);
	// await owlTokenContract.connect(deployer).approve(owlGameAddress, prizeAmount);
	// console.log(`approve prizeAmount success`);
	// await owlGameContract.connect(deployer).addPrize(prizeAmount);
	// console.log(`prizeAmount = ${prizeAmount}`);


	await owlTokenContract.connect(deployer).approve(owlGameAddress, BigInt(100000 * 1) * decimal);
	console.log("done approve");
	tx = await owlGameContract.connect(deployer).requestMint(1);
	await printTxDetail(tx, 'requestMint Count=1');

	return;
}

async function deployOrConnect() {
	if (owlTokenAddress) {
		owlTokenContract = await hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddress));
		console.log(`OwlToken contract connected to : ${owlTokenAddress}`);
	} else {
		const params = [ownerAddress];
		[owlTokenContract, owlTokenAddress] = await deploy("OwlToken", params);
		console.log(`OwlToken contract deployed to : ${owlTokenAddress}\nParams = ${params.join(" ")}\n`);
	}


	if (owldinalNftAddress) {
		owldinalNftContract = await hre.ethers.getContractFactory("Owldinal").then((c) => c.attach(owldinalNftAddress));
		console.log(`Owldinal contract connected to : ${owldinalNftAddress}`);
	} else {
		const params = [10000000, owlTokenAddress, deployer.address, backendAddress];
		[owldinalNftContract, owldinalNftAddress] = await deploy("Owldinal", params);
		console.log(`Owldinal contract deployed to : ${owldinalNftAddress}\nParams = ${params.join(" ")}\n`);
	}

	if (owlGameAddress) {
		owlGameContract = await hre.ethers.getContractFactory("OwlGame").then((c) => c.attach(owlGameAddress));
		console.log(`OwlGame contract connected to : ${owlGameAddress}`);
	} else {
		const params = [backendAddress];
		[owlGameContract, owlGameAddress] = await deploy("OwlGame", params);
		console.log(`OwlGame contract deployed to : ${owlGameAddress}\nParams = ${params.join(" ")}\n`);
	}

	if (genOneBoxAddress) {
		genOneBoxContract = await hre.ethers.getContractFactory("OwldinalGenOneBox").then((c) => c.attach(genOneBoxAddress));
		console.log(`OwldinalGenOneBox contract connected to : ${genOneBoxAddress}`);
	} else {
		const params = [owlTokenAddress, owlGameAddress];
		[genOneBoxContract, genOneBoxAddress] = await deploy("OwldinalGenOneBox", params);
		console.log(`OwldinalGenOneBox contract deployed to : ${genOneBoxAddress}\nParams = ${params.join(" ")}\n`);
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