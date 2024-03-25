require("dotenv").config();
const { ethers } = require("hardhat");

async function deploy(name, args) {
	const contract = await hre.ethers.deployContract(name, args).then((c) => c.waitForDeployment());
	const addr = await contract.getAddress();
	return [contract, addr];
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

module.exports = {
	deploy, generateSignForMint
};