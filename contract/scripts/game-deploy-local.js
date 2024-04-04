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
	let tx;

	// owlTokenAddress = "0x2279B7A0a67DB372996a5FaB50D91eAA73d2eBe6";
	// owldinalNftAddress = "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318";
	// owlGameAddress = "0x49fd2BE640DB2910c2fAb69bB8531Ab6E76127ff";
	// genOneBoxAddress = "0x4631BCAbD6dF18D94796344963cB60d44a4136b6";

	await deployOrConnect();

	console.log(`
NFT_OWL_ADDR=${owldinalNftAddress}
OWL_TOKEN_ADDR=${owlTokenAddress}
NFT_MYSTERY_BOX_ADDR=${genOneBoxAddress}
OWL_GAME_ADDR=${owlGameAddress}
	`);

	// init contract config
	const operations = [
		await owlGameContract.connect(deployer).initialize(owlTokenAddress, owldinalNftAddress, genOneBoxAddress),
	];

	await Promise.all(operations).then((results) => Promise.all(results.map(result => result.wait())));

	console.log(`\nconst [owlTokenAddress, owldinalNftAddress, genOneBoxAddress, owlGameAddress] = ["${owlTokenAddress}", "${owldinalNftAddress}", "${genOneBoxAddress}", "${owlGameAddress}"];\n`);


	// Test stakeOwldinalNft
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 2);
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 3);
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 4);
	tx = await owlGameContract.connect(deployer).stakeOwldinalNft([2, 3, 4]);
	printTxDetail(tx, 'stakeOwldinalNft 2,3,4');
	tx = await owlGameContract.connect(deployer).unstakeOwldinalNft([2]);
	printTxDetail(tx, 'unstakeOwldinalNft 2');
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 2);
	tx = await owlGameContract.connect(deployer).stakeOwldinalNft([2]);
	printTxDetail(tx, 'stakeOwldinalNft 2');

	// Test mintMysteryBox
	await owlTokenContract.connect(deployer).approve(owlGameAddress, BigInt(100000 * 1 * 10 ** 18));
	tx = await owlGameContract.connect(deployer).mintMysteryBox(1);
	printTxDetail(tx, 'mintMysteryBox Count=1');

	await owlTokenContract.connect(deployer).approve(owlGameAddress, BigInt(100000 * 10 * 10 ** 18));
	tx = await owlGameContract.connect(deployer).mintMysteryBox(10);
	printTxDetail(tx, 'mintMysteryBox Count=10');

	await owlTokenContract.connect(deployer).approve(owlGameAddress, BigInt(100000 * 100 * 10 ** 18));
	tx = await owlGameContract.connect(deployer).mintMysteryBox(100);
	printTxDetail(tx, 'mintMysteryBox Count=100');

	// add prize
	const prizeAmount = BigInt(6_0000_0000 * 10 ** 18);
	await owlTokenContract.connect(deployer).mint(ownerAddress, prizeAmount);
	await owlTokenContract.connect(deployer).approve(owlGameAddress, prizeAmount);
	await owlGameContract.connect(deployer).addPrize(prizeAmount);

	const tokenList5 = [];
	const tokenList50 = [];

	const fruitIds = [];
	const elfIds = [];

	for (i = 1; i < 111; i++) {
		const owner = await genOneBoxContract.ownerOf(i);
		const boxType = await genOneBoxContract.getBoxType(i);
		if (owner == ownerAddress) {
			if (tokenList5.length < 5) {
				tokenList5.push(i);
			} else if (tokenList50.length < 50) {
				tokenList50.push(i);
			} else {
				break;
			}

			if (boxType == 1) {
				elfIds.push(i);
			} else if (boxType == 2) {
				fruitIds.push(i);
			}
		}
	}

	console.log(JSON.stringify(elfIds));
	console.log(JSON.stringify(fruitIds));

	// Stake Token
	await genOneBoxContract.connect(deployer).setApprovalForAll(owlGameAddress, true);
	tx = await owlGameContract.connect(deployer).stakeMysteryBox(tokenList5);
	printTxDetail(tx, 'stakeMysteryBox Count=5');
	tx = await owlGameContract.connect(deployer).stakeMysteryBox(tokenList50);
	printTxDetail(tx, 'stakeMysteryBox Count=50');

	// update fruit
	tx = await owlGameContract.connect(deployer).updateAllFruitRewards();
	printTxDetail(tx, 'updateAllFruitRewards');

	// upstake
	tx = await owlGameContract.connect(deployer).claimAndUnstakeMysteryBox(tokenList5);
	printTxDetail(tx, 'claimAndUnstakeMysteryBox Count=5');
	tx = await owlGameContract.connect(deployer).claimAndUnstakeMysteryBox(tokenList50);
	printTxDetail(tx, 'claimAndUnstakeMysteryBox Count=50');
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