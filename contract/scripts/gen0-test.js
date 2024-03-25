const hre = require("hardhat");
require("dotenv").config();
const { generateSignForMint } = require('./utils');

async function main() {
	const MINT_PRICE = hre.ethers.parseUnits("0.0045", "ether");
	const [deployer] = await hre.ethers.getSigners();
	const ownerAddress = deployer.address;
	const backendAddress = process.env.BACKEND_WALLET;
	console.log(`ownerAddress = ${ownerAddress}`);
	console.log(`backendAddress = ${backendAddress}`);

	let owlTokenAddress, boxGen0Address;

	owlTokenAddress = "0x8A791620dd6260079BF849Dc5567aDC3F2FdC318";
	boxGen0Address = "0x610178dA211FEF7D417bC0e6FeD39F05609AD788";

	const [
		owlTokenContract,
		boxGen0Contract,
	] = await Promise.all(
		[
			hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddress)),
			hre.ethers.getContractFactory("MysteryBoxGen0").then((c) => c.attach(boxGen0Address)),
		]
	);

	console.log(`symbol = ${await owlTokenContract.symbol()}`);
	console.log(`TokenID = ${await boxGen0Contract.tokenIdCounter()}`);
	console.log(`Balance = ${await boxGen0Contract.connect(deployer).balanceOf(deployer.address)}`);

	let { hash, signature } = {
		"wallet": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
		"hash": "0x8d93beefb9215275253cbbeb8ade84c5540865d0a0b47cfba6929221b5341279",
		"signature": "0x4578f53dff89ec040e6b9d6dda13ebaa1499611afa5dff11efdc21e992049f5d7bdee33367d85879b880161bda77710b5af6e4e390d88126a824808cc0879faa1b"
	}
	// let { hash, signature } = await generateSignForMint(deployer, boxGen0Contract);

	console.log(`hash = ${hash}, sign = ${signature}`);
	console.log(`mint result = ${await boxGen0Contract.connect(deployer).mint(signature, { value: MINT_PRICE })}`);
	console.log(`Balance = ${await boxGen0Contract.connect(deployer).balanceOf(deployer.address)}`);
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});