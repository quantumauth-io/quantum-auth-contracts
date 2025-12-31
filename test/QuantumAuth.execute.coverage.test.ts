import { expect } from "chai";
import hre from "hardhat";
import { HDNodeWallet } from "ethers";

// Hardhat 3: provider is on the NetworkConnection, not ethers.provider / hre.network.provider
const { ethers, provider } = await hre.network.connect();

const MODE_NORMAL = 0;
const MODE_RECOVERY = 1;

function emptyBytes32(): string {
    return "0x" + "00".repeat(32);
}

function emptyUserOp(sender: string, signature: string) {
    return {
        sender,
        nonce: 0n,
        initCode: "0x",
        callData: "0x",
        accountGasLimits: emptyBytes32(),
        preVerificationGas: 0n,
        gasFees: emptyBytes32(),
        paymasterAndData: "0x",
        signature,
    };
}

// keyId encodes expected signer address in low 20 bytes (bytes32(uint160(addr)))
function keyIdFromAddress(addr: string): string {
    return ethers.zeroPadValue(addr, 32);
}

// Hardhat default mnemonic (works unless you changed config)
const DEFAULT_MNEMONIC = "test test test test test test test test test test test junk";

function walletForIndex(index: number) {
    return HDNodeWallet.fromPhrase(DEFAULT_MNEMONIC, undefined, `m/44'/60'/0'/0/${index}`);
}

function signDigestWithWallet(wallet: HDNodeWallet, digest32: string): string {
    const sig = wallet.signingKey.sign(digest32);
    return ethers.Signature.from(sig).serialized;
}

async function setBalance(address: string, hexWei: string) {
    await provider.request({
        method: "hardhat_setBalance",
        params: [address, hexWei],
    });
}

async function impersonate(address: string) {
    await provider.request({
        method: "hardhat_impersonateAccount",
        params: [address],
    });
    return ethers.getSigner(address);
}

async function stopImpersonate(address: string) {
    await provider.request({
        method: "hardhat_stopImpersonatingAccount",
        params: [address],
    });
}

async function deployAccount(tpmIndex: number) {
    const [eoa1, eoa2] = await ethers.getSigners();

    const Verifier = await ethers.getContractFactory("TPMVerifierSecp256k1");
    const verifier = await Verifier.deploy();

    const Harness = await ethers.getContractFactory("QuantumAuthAccountHarness");
    const dummyEntryPoint = "0x0000000000000000000000000000000000000001";

    const tpmKeyId = keyIdFromAddress(tpmIndex === 0 ? eoa1.address : eoa2.address);

    const account = await Harness.deploy(
        dummyEntryPoint,
        eoa1.address,
        eoa2.address,
        await verifier.getAddress(),
        tpmKeyId
    );

    return { account, verifier, dummyEntryPoint, eoa1, eoa2, tpmKeyId };
}

describe("QuantumAuthAccount execute coverage (Hardhat 3)", function () {
    it("execute: reverts if caller is not EntryPoint", async function () {
        const { account, eoa1 } = await deployAccount(0);

        await expect(account.execute(eoa1.address, 0n, "0x")).to.be.revert(ethers);
    });

    it("execute: works when called by EntryPoint", async function () {
        const { account, dummyEntryPoint } = await deployAccount(0);

        const Receiver = await ethers.getContractFactory("Receiver");
        const recv = await Receiver.deploy();

        // impersonate dummy entry point
        const ep = await impersonate(dummyEntryPoint);
        await setBalance(dummyEntryPoint, "0x1000000000000000000");

        const data = recv.interface.encodeFunctionData("ping", [123n]);
        await account.connect(ep).execute(await recv.getAddress(), 0n, data);

        await stopImpersonate(dummyEntryPoint);

        await stopImpersonate(dummyEntryPoint);
    });

    it("execute: bubbles revert data from target (assembly revert path)", async function () {
        const { account, dummyEntryPoint } = await deployAccount(0);

        const Reverter = await ethers.getContractFactory("Reverter");
        const rev = await Reverter.deploy();

        await setBalance(dummyEntryPoint, "0x1000000000000000000");
        const ep = await impersonate(dummyEntryPoint);

        const data = rev.interface.encodeFunctionData("boom");
        await expect(account.connect(ep).execute(await rev.getAddress(), 0n, data)).to.be.revert(ethers);

        await stopImpersonate(dummyEntryPoint);
    });

    it("withdraw: always reverts NotAuthorized", async function () {
        const { account, eoa1 } = await deployAccount(0);

        await expect(account.withdraw(eoa1.address, 1n)).to.be.revert(ethers);
    });

    it("validateSignature: reverts on invalid mode", async function () {
        const { account } = await deployAccount(0);

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:badmode"));
        const abi = ethers.AbiCoder.defaultAbiCoder();

        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [9, "0x", "0x", "0x"]);
        const userOp = emptyUserOp(await account.getAddress(), bundle);

        await expect(account.exposedValidateSignature(userOp, userOpHash)).to.be.revert(ethers);
    });

    it("MODE_NORMAL: fails when TPM signature is wrong (covers TPM false branch)", async function () {
        const { account, eoa1 } = await deployAccount(0);

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:normal:bad-tpm"));

        // Valid EOA1 signature (prefixed)
        const eoa1Sig = await eoa1.signMessage(ethers.getBytes(userOpHash));

        // Wrong TPM sig: signed by wallet index 1 instead of 0
        const badTpmSig = signDigestWithWallet(walletForIndex(1), userOpHash);

        const abi = ethers.AbiCoder.defaultAbiCoder();
        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [
            MODE_NORMAL,
            eoa1Sig,
            "0x",
            badTpmSig,
        ]);

        const userOp = emptyUserOp(await account.getAddress(), bundle);
        const res = await account.exposedValidateSignature(userOp, userOpHash);

        expect(res).to.equal(1n);
    });

    it("MODE_RECOVERY: fails when sig2 is empty (covers empty signature branch)", async function () {
        const { account, eoa1 } = await deployAccount(0);

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:recovery:empty-sig2"));
        const sig1 = await eoa1.signMessage(ethers.getBytes(userOpHash));

        const abi = ethers.AbiCoder.defaultAbiCoder();
        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [
            MODE_RECOVERY,
            sig1,
            "0x",
            "0x",
        ]);

        const userOp = emptyUserOp(await account.getAddress(), bundle);
        const res = await account.exposedValidateSignature(userOp, userOpHash);

        expect(res).to.equal(1n);
    });
});
