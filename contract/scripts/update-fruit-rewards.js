const hre = require("hardhat");
const { deploy } = require('./utils');

let deployer, ownerAddress;

// deploy contracts
let owlGameContract, owlGameAddress;

const decimal = BigInt(10) ** BigInt(18)

async function main() {
	const chainId = await hre.ethers.provider.getNetwork();
	console.log(`Current Chain Id: ${JSON.stringify(chainId)}`);

	[deployer] = await hre.ethers.getSigners();
	ownerAddress = deployer.address;
	let tx;

	owlGameAddress = "0x8338B3295f87DEBa418c2F0bb7497414b0D73AC2";
	owlGameContract = await hre.ethers.getContractFactory("OwlGame").then((c) => c.attach(owlGameAddress));
	console.log(`OwlGame contract connected to : ${owlGameAddress}`);

	console.log(`[${new Date().toLocaleString()}] Start updateAllFruitRewards`);
	tx = await owlGameContract.connect(deployer).updateAllFruitRewards({ gasLimit: 100_000_000_000 });
	await printTxDetail(tx, 'updateAllFruitRewards');
}

async function printTxDetail(tx, msg) {
	const receipt = await tx.wait();
	const gasUsed = receipt.gasUsed;
	console.log(`[${new Date().toLocaleString()}] ${msg} [Gas=${gasUsed}, DataLen=${tx.data.length}]`);
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