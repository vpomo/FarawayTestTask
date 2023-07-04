import { SignerWithAddress } from "@nomiclabs/hardhat-ethers/signers";
import { expect } from "chai";
import { deployments, ethers, network, run } from "hardhat";
import { defaultAbiCoder, keccak256, parseEther } from "ethers/lib/utils";

import { BigNumber } from "ethers";

const { getContract, getSigners, provider } = ethers;
const { AddressZero, HashZero } = ethers.constants;
const hre = require("hardhat");

let controller;
let collection;
let owner: SignerWithAddress;
let second: SignerWithAddress;

describe("Test Collections basic functionality", function () {

    before(async function () {
        const controllerFactory = await ethers.getContractFactory("contracts/Controller.sol:Controller");
        controller = await controllerFactory.deploy();
        await controller.deployed();

        [owner, second] = await ethers.getSigners();
    })

    it("Should create collection", async function () {
        const name = "Test NFT name";
        const symbol = "Test.NFT";
        const baseUri = "https://api.faraway.com"

        await controller.createCollection(name, symbol, baseUri);

        const collections = await controller.getCollections();

        const nftFactory = await ethers.getContractFactory("contracts/NFT.sol:NFT");
        collection = await nftFactory.attach(collections[0]);

        const collectionName = await collection.name();
        expect(collectionName).to.equal(name);
        const collectionSymbol = await collection.symbol();
        expect(collectionSymbol).to.equal(symbol);
    });

    it("Should mint token", async function () {

        const beforeBal = await collection.balanceOf(second.address);

        await controller.mint(collection.address, second.address);

        const afterBal = await collection.balanceOf(second.address);

        expect(beforeBal).to.equal(BigNumber.from("0"));
        expect(afterBal).to.equal(BigNumber.from("1"));
    });

});
