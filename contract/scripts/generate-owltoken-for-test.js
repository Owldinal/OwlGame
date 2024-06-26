const hre = require("hardhat");
const { deploy, sleep } = require('./utils');
const { ethers } = require("hardhat");

let deployer, ownerAddress;

// deploy contracts
let owlTokenContract, owlTokenAddress;
let owldinalNftContract, owldinalNftAddress;
let owlGameContract, owlGameAddress;
let genOneBoxContract, genOneBoxAddress;

let addressList = [];

const decimal = BigInt(10) ** BigInt(18)

async function initAddress() {
	owldinalNftAddress = "0x3FD1e177e98017902fc1a0F51De10Dc28645D034";
	owlTokenAddress = "0xe13481571787F0fdb2C7AD8D295c18856cA76494";
	genOneBoxAddress = "0xB73b4C72Dfc10B46cdF4CCDe63Def3336566B3BD";
	owlGameAddress = "0xfF454DbF7aBe0aF44C78c16DDccD6f3F80166F12";

	addressList = [
		// "0xf3bBE5dBB9db8Ca9B04BBE0e9c814c1609572a74",
	];

	for (j = 1; j < 100; j++) {
		var user = (await hre.ethers.getSigners())[j];
		addressList.push(user.address);
	}
}

async function main() {
	[deployer] = await hre.ethers.getSigners();
	ownerAddress = deployer.address;
	backendAddress = process.env.BACKEND_WALLET;
	console.log(`Deployer Address = ${deployer.address}`);

	const blockNumber = await hre.ethers.provider.getBlockNumber();
	console.log(`Current Block number: ${blockNumber}`);
	await initAddress();

	await deployOrConnect();

	const userList = [];

	let tx;

	for (j = 0; j < addressList.length; j++) {
		var addr = addressList[j];
		console.log(`User: ${addr}`);

		var userObj = {
			address: addr,
			tokenIds: []
		}

		try {
			tx = await owlTokenContract.connect(deployer).mint(addr, BigInt(100000 * 100000) * decimal);
			await tx.wait();
		} catch (e) {
			console.log(`Failed mint owlToken for User: ${addr}`);
			console.log(e);
		}
		console.log(`done mint, balance = ${await owlTokenContract.balanceOf(addr)}`);
		await sleep(500);

		for (i = 0; i < 3; i++) {
			try {

				tx = await owldinalNftContract.connect(deployer).mintByAdmin();
				const receipt = await tx.wait();
				console.log("done mint nft");
				await sleep(500);

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
				tx = await owldinalNftContract.connect(deployer).transferFrom(deployer.address, addr, tokenId);
				await tx.wait();
				console.log(`done transfer nft, tokenId = ${tokenId}`);

				userObj.tokenIds.push(Number(tokenId));
				await sleep(500);
			} catch (e) {
				console.log(`Failed mint nft for User: ${addr}`);
			}
		}

		userList.push(userObj);
	}

	console.log("All Done: \n");
	console.log(JSON.stringify(userList));
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

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});

