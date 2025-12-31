
import { expect } from "chai";
import hre from "hardhat";

const { ethers } = await hre.network.connect();

describe("QAERC20 coverage", () => {
    it("constructor sets name/symbol/decimals", async () => {
        const [owner] = await ethers.getSigners();
        const Token = await ethers.getContractFactory("QAERC20");

        const token = await Token.deploy("QA Token", "QAT", 6, owner.address);

        expect(await token.name()).to.equal("QA Token");
        expect(await token.symbol()).to.equal("QAT");
        expect(await token.decimals()).to.equal(6);
    });

    it("mint: onlyOwner, updates balance", async () => {
        const [owner, other] = await ethers.getSigners();
        const Token = await ethers.getContractFactory("QAERC20");
        const token = await Token.deploy("QA Token", "QAT", 18, owner.address);

        // OZ v5 Ownable reverts with custom error: OwnableUnauthorizedAccount(address)
        await expect(token.connect(other).mint(other.address, 1n)).to.be.revert(ethers);

        await token.connect(owner).mint(other.address, 123n);
        expect(await token.balanceOf(other.address)).to.equal(123n);
    });

    it("burn burns caller balance", async () => {
        const [owner] = await ethers.getSigners();
        const Token = await ethers.getContractFactory("QAERC20");
        const token = await Token.deploy("QA Token", "QAT", 18, owner.address);

        await token.mint(owner.address, 100n);
        await token.burn(40n);

        expect(await token.balanceOf(owner.address)).to.equal(60n);
    });

    it("burnFrom requires allowance and reduces it", async () => {
        const [owner, spender] = await ethers.getSigners();
        const Token = await ethers.getContractFactory("QAERC20");
        const token = await Token.deploy("QA Token", "QAT", 18, owner.address);

        await token.mint(owner.address, 100n);

        // No allowance => revert (OZ uses custom error ERC20InsufficientAllowance)
        await expect(token.connect(spender).burnFrom(owner.address, 1n)).to.be.revert(ethers);

        await token.approve(spender.address, 25n);
        await token.connect(spender).burnFrom(owner.address, 20n);

        expect(await token.balanceOf(owner.address)).to.equal(80n);
        expect(await token.allowance(owner.address, spender.address)).to.equal(5n);
    });
});
