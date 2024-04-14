const hre = require("hardhat");
const { deploy } = require('./utils');
const { ethers } = require("hardhat");

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

	// Server 
	// owldinalNftAddress = "0xF6E4Af62bD0f298311079503563ea4B1cfF6Dead";
	// owlTokenAddress = "0x15463B075F37c269830801feCd8CdAC76Aa1A310";
	// genOneBoxAddress = "0x4b54fC01A0714f79Ec0Bd1856f62Dc710601ef7f";
	// owlGameAddress = "0xe53A375D3B997FB22f8bF85910280A011042aDf2";
	// TestNet
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


	const addressList = [
		"0x5CbAd14e910b95Ba630482f1e65C49F075171f7f",
		"0x57931f17ee7fc505ce6b225df7b1213b0953feb9",
		"0x6ce580b0e9081bd60ba8d11f37f4eed1796556b8",
		"0x0e0245324eec7fbff8c91741d3d20aedeef699e9",
		"0x13644428c78BaB08889d3b571eD607E6f90c164d",
		"0xd182278fdBEdFb61fD7B8a8D478059A0a40c31c",
		"0x6ce580b0e9081bd60ba8d11f37f4eed1796556b8",
		"0x5f48a1c54b2a827d990e76B963Fa37CcF05c8e35",
		"0x4d5755f067611a58ab0fbb4becce1816127bb0fe",
		"0x6ef683722fcfdfd77f69c67f9648ae856b2e3ed9",
		"0x187f6526ff054E1d6bA02A44509e435BEbCD95E1",
		"0x3Cbcbf60f62fe8544B23B03CA7C7aF2C9b6f2B7d",
		"0xd53282922b7cf6687005e256cc980210d521fc41",
		"0xbb3Db0959822B70a42417ae0dC6a79553a055344",
		"0xc80ddc83a7848dba336d1f9a300bcbedd9df7730",
		"0x236dCf7295eAC62e0373f1af8D469ADDA305b878",
		"0x187f6526ff054E1d6bA02A44509e435BEbCD95E1",
		"0x236dCf7295eAC62e0373f1af8D469ADDA305b878",
		"0xd04d7108367f78a8c62cb4e6645ce3dfc6065bc2",
		"0x40006d4a2f3f4eabe8bbb0aebc8c158835db1084",
		"0x744d5e9e24168645baa3656cb7d14efec30092b7",
		"0xd182278fdBEdFb61fD7B8a8D478059A0a40c31c6",
		"0xdDB4Bf4BeeAa68672A83db1789fAb53c24D84247",
		"0x7b6ee97ba71f8d95b6186ae71585a0cae981879f",
		"0x0517Da61De392E738aaBAFe14cE33FbE4b14BA1D",
		"0x3F1Df87349d9040a22E4414eB4dDb6173E300278",
		"0xa9a7e20e47896d77a7c6936af5c42eaaf4377cfd",
		"0xE2493CFD86101a368ba8d550596Aa7d7C69f321c",
		"0xed9045b3f19bde6135ea2ebecab79c4dce8383a2",
		"0x7835c546a273222819451e869b2e6a25375720d8",
		"0xb910189d34b6bef7d1015cf677ea5101b1d6c389",
		"0xb1b3cc18cb4dc201246f9d70c21bd6c2d59db729",
		"0x0d60c0a7aa6d050759f784fb056eab213416670c",
		"0xac3d39c4e5b15aa5c27be113c7a9c6937f067cf8",
		"0xc1aBa1b8aCde028dff8bbbca4cEF77fF6285e663",
		"0x0f630ff892576fe76c0eeceb18c919b337b0fa30",
		"0x65c2f0f3da2671c2f344bcd390e162a9a480044e",
		"0x1d8f320da36e02ee71fa5529c049563b9f49186b",
		"0x68FFC6A11134f3ddA32d5B008bf7B996939fB694",
		"0x918e2319134f718a59fcaba235dcbe0c02a2d215",
		"0x6D9afaAC2F6d7dcba72f33dcbC6960B5332d67b3",
		"0x2fd66c76684734682f29875BA17a0C4AB3a501A4",
		"0xda9e0fe894cdabf3c17e2cdc4f9db86c9fbbe70a",
		"0x08b1E7388e10ED568b214De194438e42b9CAA945",
		"0x456d5441f1b9c8b1ddf0568ca3218f2139c3efd6",
		"0xECffb8De150A879D579C58f004727E92cea14fF9",
		"0x5f88fe96ebc7f9857a04e9ea55bbaa9b889d0f06",
		"0x18202bE2FAEe609F555Ce52Bd2FEE048F7e1150E",
		"0x8f8509cd9f7b0cf2f7f530bc50be3161553c7f22",
		"0xc49f6E54728F48Aaae984A3c69C3C9D161B41031",
		"0xf3fd1060abd4422df4bb34fc58dcc336da541332",
		"0x3bbdd6e34d4476af80d294191512a1ace20de0f1",
		"0x439086b57b820de8bd2d83dcefd9263abeeac4e7",
		"0xdF4379735Ea99D9fD5b6a4CDEA517a5E6e20E44E",
		"0x18071D75576c967b238A3e316384e9a1CAE4145a",
		"0xb52A76A2A93C3DAE3f3bE6820D17C5E2E37BCf27"
	];

	for (j = 0; j < addressList.length; j++) {
		var user = addressList[j];
		console.log(`For User ${addressList[j]}`);

		try {
			await owlTokenContract.connect(deployer).mint(user, BigInt(5 * 100000) * decimal);
		} catch (e) {
			console.log(`${e}`);
		}

		for (i = 0; i < 3; i++) {
			try {
				const tx = await owldinalNftContract.connect(deployer).mintByAdmin();
				const receipt = await tx.wait();

				var mintBoxLogs = receipt.logs.filter(log => log.topics[0] === "0xf5d3f864a50c2df29b92152f2936fc5520ee555438f668048785c1868cd34230");
				// console.log(`Logs: ${JSON.stringify(mintBoxLogs)}`);

				if (!mintBoxLogs || mintBoxLogs.length === 0) {
					console.error("MintBox event not found");
					return;
				}
				const decodedEventData = ethers.AbiCoder.defaultAbiCoder().decode(
					["uint256", "uint256", "string"],
					mintBoxLogs[0].data
				);

				const tokenId = decodedEventData[0];
				console.log(`Minted tokenId: ${tokenId}`);

				await owldinalNftContract.connect(deployer).transferFrom(deployer.address, user, tokenId);
			} catch (e) {
				console.log(`${e}`);
			}
		}
	}
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