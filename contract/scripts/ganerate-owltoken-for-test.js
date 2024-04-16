const hre = require("hardhat");
const { deploy } = require('./utils');
const { ethers } = require("hardhat");

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

	// Server 
	owldinalNftAddress = "0xF6E4Af62bD0f298311079503563ea4B1cfF6Dead";
	owlTokenAddress = "0x15463B075F37c269830801feCd8CdAC76Aa1A310";
	genOneBoxAddress = "0x4b54fC01A0714f79Ec0Bd1856f62Dc710601ef7f";
	owlGameAddress = "0xe53A375D3B997FB22f8bF85910280A011042aDf2";

	await deployOrConnect();

	console.log(`
owlTokenAddress= "${owlTokenAddress}";
owldinalNftAddress= "${owldinalNftAddress}";
owlGameAddress= "${owlGameAddress}";
genOneBoxAddress= "${genOneBoxAddress}";
		`);

	console.log(`
EVENT_START_BLOCK=${blockNumber}
NFT_OWL_ADDR=${owldinalNftAddress}
OWL_TOKEN_ADDR=${owlTokenAddress}
NFT_MYSTERY_BOX_ADDR=${genOneBoxAddress}
OWL_GAME_ADDR=${owlGameAddress}
			`);

	const userPrivKeyList = []
	const userAddressList = []
	for (j = 1; j < 80; j++) {
		var user = (await hre.ethers.getSigners())[j];
		console.log(`User: ${user.address}`);

		userPrivKeyList.push(user.privateKey);
		userAddressList.push(user.address);

		try {
			await owlTokenContract.connect(deployer).mint(user.address, BigInt(100000 * 100000) * decimal);
		} catch (e) {
			console.log(`Failed mint owlToken for User: ${user.address}`);
		}

		for (i = 0; i < 3; i++) {
			try {

				const tx = await owldinalNftContract.connect(deployer).mintByAdmin();
				const receipt = await tx.wait();

				var mintBoxLogs = receipt.logs.filter(log => log.topics[0] === "0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230");
				// console.log(`Logs: ${JSON.stringify(mintBoxLogs)}`);

				if (!mintBoxLogs || mintBoxLogs.length === 0) {
					console.error("MintBox event not found");
					return;
				}
				const decodedEventData = ethers.AbiCoder.defaultAbiCoder().decode(
					["uint256", "uint256", "string"],
					mintBoxLogs[0].data
				);

				const tokenId = decodedEventData[0];
				console.log(`Minted tokenId: ${tokenId}`);

				await owldinalNftContract.connect(deployer).transferFrom(deployer.address, user.address, tokenId);

			} catch (e) {
				console.log(`Failed mint nft for User: ${user.address}`);
			}
		}
	}
}

async function printTxDetail(tx, msg) {
	const receipt = await tx.wait();
	const gasUsed = receipt.gasUsed;
	console.log(`${msg} [Gas=${gasUsed}, DataLen=${tx.data.length}]`);
}

async function deployOrConnect() {
	if (owlTokenAddress) {
		owlTokenContract = await hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddress));
		console.log(`OwlToken contract connected to : ${owlTokenAddress}`);
	} else {
		const params = [ownerAddress];
		[owlTokenContract, owlTokenAddress] = await deploy("OwlToken", params);
		console.log(`OwlToken contract deployed to : ${owlTokenAddress}\nParams = ${params.join(" ")}`);

		await owlTokenContract.connect(deployer).mint(deployer.address, BigInt(100000000) * decimal);
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
		const params = [ownerAddress];
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