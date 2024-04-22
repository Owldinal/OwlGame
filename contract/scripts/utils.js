require("dotenv").config();
const { ethers } = require("hardhat");

async function deploy(name, args) {
	const contract = await hre.ethers.deployContract(name, args).then((c) => c.waitForDeployment());
	const addr = await contract.getAddress();
	return [contract, addr];
}


function sleep(ms) {
	return new Promise(resolve => setTimeout(resolve, ms));
}

async function generateSignForMint(signer, boxGen0Contract) {
	const privateKey = process.env.BACKEND_PRIVATE_KEY;
	const wallet = new ethers.Wallet(privateKey);
	const contractAddr = await boxGen0Contract.getAddress();
	const signerAddr = signer.address;
	const message = ethers.AbiCoder.defaultAbiCoder().encode(['address', 'address'],
		[signerAddr, contractAddr]);

	const hash = ethers.keccak256(message);
	const signature = await wallet.signMessage(ethers.getBytes(hash));

	return { hash, signature }
}

async function sendEth(signer, receiverAddress, amountInEther) {
	const amountInWei = hre.ethers.parseEther(amountInEther);

	const tx = {
		to: receiverAddress,
		value: amountInWei
	};

	const transactionResponse = await signer.sendTransaction(tx);

	await transactionResponse.wait();

	console.log(`Transaction successful with hash: ${transactionResponse.hash}`);
}

async function printTxDetail(tx, msg) {
	const receipt = await tx.wait();
	const gasUsed = receipt.gasUsed;
	console.log(`${msg} [Gas=${gasUsed}, DataLen=${tx.data.length}]`);
}

module.exports = {
	deploy, generateSignForMint, sendEth, sleep, printTxDetail
};