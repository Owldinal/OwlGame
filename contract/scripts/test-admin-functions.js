const hre = require("hardhat");
const { deploy, sendEth, sleep, printTxDetail } = require('./utils');

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

	owldinalNftAddress = "0x3FD1e177e98017902fc1a0F51De10Dc28645D034";
	owlTokenAddress = "0xe13481571787F0fdb2C7AD8D295c18856cA76494";
	genOneBoxAddress = "0xB73b4C72Dfc10B46cdF4CCDe63Def3336566B3BD";
	owlGameAddress = "0xfF454DbF7aBe0aF44C78c16DDccD6f3F80166F12";

	await deployOrConnect();
	console.log(`\nTest\n====`);

	// start test

	// await testTransferMysteryBox();
	// await testWithdrawPrizePool();
}


async function testTransferMysteryBox() {
	let tx;

	console.log(`\ntransfer mystery box by normal user, should exception`);
	const userWithoutWhitelist = (await hre.ethers.getSigners())[45];
	console.log(`current user is ${userWithoutWhitelist.address}`);
	try {
		tx = await genOneBoxContract.connect(userWithoutWhitelist).transferFrom(
			userWithoutWhitelist.address,
			deployer.address,
			1401
		);
		await printTxDetail(tx, "Should exception when transfer mystery box");
		return;
	} catch (e) {
		console.log(`Got exception when transfer mystery box for normal user. ${e}`);
	}


	console.log(`\ntransfer mystery box by whitelist`);
	const userInWhitList = (await hre.ethers.getSigners())[23];
	console.log(`current user is ${userInWhitList.address}`);
	tx = await genOneBoxContract.connect(userInWhitList).transferFrom(
		userInWhitList.address,
		deployer.address,
		54
	);
	await printTxDetail(tx, "Success transfer");
	tx = await genOneBoxContract.connect(deployer).transferFrom(
		deployer.address,
		userInWhitList.address,
		54
	);
	await printTxDetail(tx, "Success transfer back");


	console.log(`\nTest global enable transfer mystery box`);
	tx = await genOneBoxContract.connect(deployer).setTransferEnable(
		true
	);
	await printTxDetail(tx, "Enable global transfer");
	tx = await genOneBoxContract.connect(userWithoutWhitelist).transferFrom(
		userWithoutWhitelist.address,
		deployer.address,
		1401
	);
	await printTxDetail(tx, "Success transfer");
	tx = await genOneBoxContract.connect(deployer).transferFrom(
		deployer.address,
		userWithoutWhitelist.address,
		1401
	);
	await printTxDetail(tx, "Success transfer back");
	tx = await genOneBoxContract.connect(deployer).setTransferEnable(
		false
	);
	await printTxDetail(tx, "Disable global transfer");


	console.log(`\nTest add & delete white list for transfer mystery box`);
	tx = await genOneBoxContract.connect(deployer).addTransferWhiteList(
		[userWithoutWhitelist.address]
	);
	await printTxDetail(tx, "Add white list");
	tx = await genOneBoxContract.connect(userWithoutWhitelist).transferFrom(
		userWithoutWhitelist.address,
		deployer.address,
		1401
	);
	await printTxDetail(tx, "Success transfer");
	tx = await genOneBoxContract.connect(deployer).transferFrom(
		deployer.address,
		userWithoutWhitelist.address,
		1401
	);
	await printTxDetail(tx, "Success transfer back");
	tx = await genOneBoxContract.connect(deployer).deleteTransferWhiteList(
		[userWithoutWhitelist.address]
	);
	await printTxDetail(tx, "Delete white list");

}

async function testWithdrawPrizePool() {
	const balance = await owlTokenContract.balanceOf(owlGameAddress);
	console.log(`Game balance is : ${balance}`);
	console.log(`Deployer balance is : ${await owlTokenContract.balanceOf(deployer.address)}`);

	tx = await owlGameContract.connect(deployer).withdraw(deployer.address);
	await printTxDetail(tx, "Success withdraw");

	console.log(`Game balance is : ${await owlTokenContract.balanceOf(owlGameAddress)}`);
	console.log(`Deployer balance is : ${await owlTokenContract.balanceOf(deployer.address)}`);

	tx = await owlTokenContract.connect(deployer).approve(owlGameAddress, balance);
	await printTxDetail(tx, "Success approve ");
	tx = await owlGameContract.connect(deployer).addPrize(balance);
	await printTxDetail(tx, "Success add prize");

	console.log(`Game balance is : ${await owlTokenContract.balanceOf(owlGameAddress)}`);
	console.log(`Deployer balance is : ${await owlTokenContract.balanceOf(deployer.address)}`);
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

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});