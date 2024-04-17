const hre = require("hardhat");
const { sendEth } = require('./utils');

async function main() {
	let [deployer] = await hre.ethers.getSigners();
	const balanceInEth = hre.ethers.formatEther(await hre.ethers.provider.getBalance(deployer.address));
	console.log(`From Address = ${deployer.address}, balance = ${balanceInEth}`);

	const toAddress = process.env.BACKEND_WALLET;
	const toBalanceInEth = hre.ethers.formatEther(await hre.ethers.provider.getBalance(toAddress));
	console.log(`To Address = ${toAddress}, balance = ${toBalanceInEth}`);

	await sendEth(deployer, toAddress, "1");
}


main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});