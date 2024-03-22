const hre = require("hardhat");
require("dotenv").config();


async function main() {
	const MINT_PRICE = hre.ethers.parseUnits("0.0045", "ether");
	const [owlTokenAddr, boxGen0Addr] = ["0x5FbDB2315678afecb367f032d93F642f64180aa3", "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"];

	const deployers = await hre.ethers.getSigners();
	console.log(`Current Signer = ${deployers[0].address}`);

	const [
		owlTokenContract,
		boxGen0Contract,
	] = await Promise.all(
		[
			hre.ethers.getContractFactory("OwlToken").then((c) => c.attach(owlTokenAddr)),
			hre.ethers.getContractFactory("MysteryBoxGen0").then((c) => c.attach(boxGen0Addr)),
		]
	);

	console.log(`symbol = ${await owlTokenContract.symbol()}`);
	for (i = 0; i <= 5; i++) {
		console.log(`Cuurent BTCBalance For ${i} = ${await hre.ethers.provider.getBalance(deployers[i].address)}`);
		console.log(`Cuurent Owl Balance For ${i} = ${await owlTokenContract.balanceOf(deployers[i].address)}`);
	}

	console.log(`TokenID = ${await boxGen0Contract.tokenIdCounter()}`);

	let { hash, signature } = await generateSignForMint(deployers[0], boxGen0Contract);
	console.log(`hash = ${hash}, sign = ${signature}`);
	console.log(`mint result = ${await boxGen0Contract.connect(deployers[0]).mintByVoya(hash, signature, { value: MINT_PRICE })}`);

	// for (i = 0; i < 100; i++) {
	// 	await blindBoxContract.connect(deployer).buyBox();
	// }
	// console.log(`Box Count = ${await blindBoxContract.balanceOf(deployer.address)}`);

	// for (i = 1; i < 51; i++) {
	// 	await blindBoxContract.connect(deployer).openBox(i);
	// }
	// console.log(`Box Count = ${await blindBoxContract.balanceOf(deployer.address)}`);
	// console.log(`NFT Owl Count = ${await nftOwlContract.balanceOf(deployer.address)}`);
	// console.log(`NFT Fairy Count = ${await nftFairyContract.balanceOf(deployer.address)}`);
	// console.log(`NFT Fruit Count = ${await nftFruitContract.balanceOf(deployer.address)}`);


}

async function generateSignForMint(signer, boxGen0Contract) {
	const privateKey = process.env.BACKEND_PRIVATE_KEY;
	const wallet = new ethers.Wallet(privateKey);
	const tokenId = await boxGen0Contract.tokenIdCounter();
	const contractAddr = await boxGen0Contract.getAddress();
	const signerAddr = signer.address;
	const message = ethers.AbiCoder.defaultAbiCoder().encode(['address', 'uint256', 'address'],
		[signerAddr, tokenId, contractAddr]);

	const hash = ethers.keccak256(message);
	const signature = await wallet.signMessage(ethers.getBytes(hash));

	return { hash, signature }
}

main().catch((error) => {
	console.error(error);
	process.exitCode = 1;
});