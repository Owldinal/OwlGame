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

	owlTokenAddress = "0xF86EA7e636f43dd6dCD35B7E04742f5508F9CF73";
	owldinalNftAddress = "0x96fB752fc565c740F8b09a1760FE5a3D89dD18E5";
	owlGameAddress = "0xC0E5d058eeF687B0c3cEf2967D0B55AD81eb9C21";
	genOneBoxAddress = "0xC9761572c264ADE5253d4F56a574fa0F4905ca5d";

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

	await owldinalNftContract.connect(deployer).mintByAdmin();
	await owldinalNftContract.connect(deployer).mintByAdmin();
	await owldinalNftContract.connect(deployer).mintByAdmin();

	// init contract config
	const operations = [
		await owlGameContract.connect(deployer).initialize(owlTokenAddress, owldinalNftAddress, genOneBoxAddress),
	];

	await Promise.all(operations).then((results) => Promise.all(results.map(result => result.wait())));

	const prizeAmount = BigInt(6_0000_0000n) * decimal;
	console.log(`prizeAmount = ${prizeAmount}`);
	await owlTokenContract.connect(deployer).mint(ownerAddress, prizeAmount);
	await owlTokenContract.connect(deployer).approve(owlGameAddress, prizeAmount);
	await owlGameContract.connect(deployer).addPrize(prizeAmount);

	console.log(`\nconst [owlTokenAddress, owldinalNftAddress, genOneBoxAddress, owlGameAddress] = ["${owlTokenAddress}", "${owldinalNftAddress}", "${genOneBoxAddress}", "${owlGameAddress}"];\n`);

	return;

	// playerB
	await owldinalNftContract.connect(playerB).approve(owlGameAddress, 5);
	await owldinalNftContract.connect(playerB).approve(owlGameAddress, 6);
	await owldinalNftContract.connect(playerB).approve(owlGameAddress, 7);
	tx = await owlGameContract.connect(playerB).stakeOwldinalNft([5, 6, 7]);

	console.log("\nTest Invite");
	const inviteCodePlayerB = await owlGameContract.getInviteCode(playerB);
	console.log(`playerB: invite code = ${inviteCodePlayerB}, string = ${decodeInviteCode(inviteCodePlayerB)}`);
	tx = await owlGameContract.connect(deployer).handleInviteCode(inviteCodePlayerB);
	const inviteCodeDeployer = await owlGameContract.getInviteCode(deployer);
	console.log(`deployer: invite code = ${inviteCodeDeployer}, string = ${decodeInviteCode(inviteCodeDeployer)}`);

	// ----

	console.log("\nTest stakeOwldinalNft");
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 2);
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 3);
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 4);
	tx = await owlGameContract.connect(deployer).stakeOwldinalNft([2, 3, 4]);
	await printTxDetail(tx, 'stakeOwldinalNft 2,3,4');
	tx = await owlGameContract.connect(deployer).unstakeOwldinalNft([2]);
	await printTxDetail(tx, 'unstakeOwldinalNft 2');
	await owldinalNftContract.connect(deployer).approve(owlGameAddress, 2);
	tx = await owlGameContract.connect(deployer).stakeOwldinalNft([2]);
	await printTxDetail(tx, 'stakeOwldinalNft 2');

	console.log("\nTest mintMysteryBox");
	await owlTokenContract.connect(deployer).approve(owlGameAddress, BigInt(100000 * 1) * decimal);
	tx = await owlGameContract.connect(deployer).mintMysteryBox(1);
	await printTxDetail(tx, 'mintMysteryBox Count=1');

	await owlTokenContract.connect(deployer).approve(owlGameAddress, BigInt(100000 * 10) * decimal);
	tx = await owlGameContract.connect(deployer).mintMysteryBox(10, { gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'mintMysteryBox Count=10');

	await owlTokenContract.connect(deployer).approve(owlGameAddress, BigInt(100000 * 100) * decimal);
	tx = await owlGameContract.connect(deployer).mintMysteryBox(100, { gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'mintMysteryBox Count=100');

	console.log("\nTest addPrize");

	await owlTokenContract.connect(deployer).mint(ownerAddress, prizeAmount, { gasLimit: 100_000_000_000 });
	await owlTokenContract.connect(deployer).approve(owlGameAddress, prizeAmount, { gasLimit: 100_000_000_000 });
	await owlGameContract.connect(deployer).addPrize(prizeAmount, { gasLimit: 100_000_000_000 });

	const tokenList5 = [];
	const tokenList50 = [];

	const fruitIds = [];
	const elfIds = [];
	const burnedIds = [];

	for (i = 1; i < 111; i++) {
		const boxType = await genOneBoxContract.getBoxType(i);
		if (boxType == 3) {
			burnedIds.push(i);
			continue;
		}

		const owner = await genOneBoxContract.ownerOf(i);
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

	console.log(`Elf: ${JSON.stringify(elfIds)}`);
	console.log(`Fruit: ${JSON.stringify(fruitIds)}`);
	console.log(`Burned: ${JSON.stringify(burnedIds)}`);

	console.log("\nTest stakeToken");
	await genOneBoxContract.connect(deployer).setApprovalForAll(owlGameAddress, true);
	tx = await owlGameContract.connect(deployer).stakeMysteryBox(tokenList5, { gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'stakeMysteryBox Count=5');
	tx = await owlGameContract.connect(deployer).stakeMysteryBox(tokenList50, { gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'stakeMysteryBox Count=50');

	console.log("\nTest updateAllFruitRewards");
	tx = await owlGameContract.connect(deployer).updateAllFruitRewards({ gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'updateAllFruitRewards');

	console.log("\nTest unstake");
	tx = await owlGameContract.connect(deployer).claimAndUnstakeMysteryBox(tokenList5, { gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'claimAndUnstakeMysteryBox Count=5');

	return;

	tx = await owlGameContract.connect(deployer).claimAndUnstakeMysteryBox(tokenList50, { gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'claimAndUnstakeMysteryBox Count=50');
	tx = await owlGameContract.connect(deployer).unstakeOwldinalNft([2, 3, 4], { gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'unstakeOwldinalNft 2,3,4');
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