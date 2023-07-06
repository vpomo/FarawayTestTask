export { };
const hre = require("hardhat");
import { BigNumber } from "ethers";

const ethers = hre.ethers;
const network = hre.network.name;


async function main() {
    console.log()
    const deployer = (await hre.ethers.getSigners())[0];
    console.log("deployer", deployer.address);

    const Controller = await ethers.getContractFactory("contracts/Controller.sol:Controller");
    const controller = await Controller.deploy()
    await controller.deployed()
    await contractVerify(controller.address, 'Controller')
}

main()
    .then(() => process.exit(0))
    .catch(error => {
        console.error(error);
        process.exit(1);
    });

async function contractVerify(contractAddress, contractName) {
    console.log('Verifying ' + contractName + ' ...');
    try {
        await run("verify:verify", { address: contractAddress });
        console.log(contractName + " verified")
    } catch {
        console.log(contractName + " not verified")
    };
}