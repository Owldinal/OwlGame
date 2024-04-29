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

	owldinalNftAddress = "0x06BAaC534601493C0678Ca81a78Fe1BF03d52Bef";
	owlTokenAddress = "0xf8bB8324Cd226f6229dbB8792C66119832791A59";
	genOneBoxAddress = "0xfB3ABE390CE28D7749951f2dA7B0A7487D234396";
	owlGameAddress = "0x73a11097dCf0817909039d2661a15cbc8F6624eF";

	await deployOrConnect();

	console.log(`
owldinalNftAddress = "${owldinalNftAddress}";
owlTokenAddress = "${owlTokenAddress}";
genOneBoxAddress = "${genOneBoxAddress}";
owlGameAddress = "${owlGameAddress}";
		`);

	console.log(`
EVENT_START_BLOCK=${blockNumber}
NFT_OWL_ADDR=${owldinalNftAddress}
OWL_TOKEN_ADDR=${owlTokenAddress}
NFT_MYSTERY_BOX_ADDR=${genOneBoxAddress}
OWL_GAME_ADDR=${owlGameAddress}
			`);

	await owlGameContract.connect(deployer).initialize(owlTokenAddress, owldinalNftAddress, genOneBoxAddress);

	// await owlGameContract.connect(deployer).initialize(owlTokenAddress, owldinalNftAddress, genOneBoxAddress);
	// await owlGameContract.connect(deployer).setMoonBoost(false);
	// await genOneBoxContract.connect(deployer).setTransferEnable(true);

	// const prizeAmount = BigInt(6_0000_0000n) * decimal;
	// await owlTokenContract.connect(deployer).mint(ownerAddress, prizeAmount);
	// console.log(`mint prizeAmount success`);
	// await owlTokenContract.connect(deployer).approve(owlGameAddress, prizeAmount);
	// console.log(`approve prizeAmount success`);
	// await owlGameContract.connect(deployer).addPrize(prizeAmount);
	// console.log(`prizeAmount = ${prizeAmount}`);

	return;
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