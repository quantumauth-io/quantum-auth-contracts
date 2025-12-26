import { expect } from "chai";
import hre from "hardhat";
import { HDNodeWallet } from "ethers";

const { ethers } = await hre.network.connect();

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

describe("QuantumAuthAccount - 2-of-3 policy", function () {
    it("MODE_NORMAL: passes with EOA1 + TPM", async function () {
        const [eoa1, eoa2] = await ethers.getSigners();

        const Verifier = await ethers.getContractFactory("TPMVerifierSecp256k1");
        const verifier = await Verifier.deploy();

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:normal:eoa1"));

        // TPM sig: raw digest signature by same key as eoa1 (index 0)
        const tpmKeyId = keyIdFromAddress(eoa1.address);
        const tpmSig = signDigestWithWallet(walletForIndex(0), userOpHash);

        // EOA sig: prefixed (as your account expects)
        const eoa1Sig = await eoa1.signMessage(ethers.getBytes(userOpHash));

        const Harness = await ethers.getContractFactory("QuantumAuthAccountHarness");
        const dummyEntryPoint = "0x0000000000000000000000000000000000000001";

        const account = await Harness.deploy(
            dummyEntryPoint,
            eoa1.address,
            eoa2.address,
            await verifier.getAddress(),
            tpmKeyId
        );

        const abi = ethers.AbiCoder.defaultAbiCoder();
        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [MODE_NORMAL, eoa1Sig, "0x", tpmSig]);

        const userOp = emptyUserOp(await account.getAddress(), bundle);
        const res = await account.exposedValidateSignature(userOp, userOpHash);
        expect(res).to.equal(0n);
    });

    it("MODE_NORMAL: passes with EOA2 + TPM", async function () {
        const [eoa1, eoa2] = await ethers.getSigners();

        const Verifier = await ethers.getContractFactory("TPMVerifierSecp256k1");
        const verifier = await Verifier.deploy();

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:normal:eoa2"));

        // TPM sig: raw digest signature by same key as eoa2 (index 1)
        const tpmKeyId = keyIdFromAddress(eoa2.address);
        const tpmSig = signDigestWithWallet(walletForIndex(1), userOpHash);

        const eoa2Sig = await eoa2.signMessage(ethers.getBytes(userOpHash));

        const Harness = await ethers.getContractFactory("QuantumAuthAccountHarness");
        const dummyEntryPoint = "0x0000000000000000000000000000000000000001";

        const account = await Harness.deploy(
            dummyEntryPoint,
            eoa1.address,
            eoa2.address,
            await verifier.getAddress(),
            tpmKeyId
        );

        const abi = ethers.AbiCoder.defaultAbiCoder();
        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [MODE_NORMAL, "0x", eoa2Sig, tpmSig]);

        const userOp = emptyUserOp(await account.getAddress(), bundle);
        const res = await account.exposedValidateSignature(userOp, userOpHash);
        expect(res).to.equal(0n);
    });

    it("MODE_NORMAL: fails with EOA1 only (no TPM)", async function () {
        const [eoa1, eoa2] = await ethers.getSigners();

        const Verifier = await ethers.getContractFactory("TPMVerifierSecp256k1");
        const verifier = await Verifier.deploy();

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:normal:no-tpm"));
        const tpmKeyId = keyIdFromAddress(eoa1.address);

        const eoa1Sig = await eoa1.signMessage(ethers.getBytes(userOpHash));

        const Harness = await ethers.getContractFactory("QuantumAuthAccountHarness");
        const dummyEntryPoint = "0x0000000000000000000000000000000000000001";

        const account = await Harness.deploy(
            dummyEntryPoint,
            eoa1.address,
            eoa2.address,
            await verifier.getAddress(),
            tpmKeyId
        );

        const abi = ethers.AbiCoder.defaultAbiCoder();
        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [MODE_NORMAL, eoa1Sig, "0x", "0x"]);

        const userOp = emptyUserOp(await account.getAddress(), bundle);
        const res = await account.exposedValidateSignature(userOp, userOpHash);
        expect(res).to.equal(1n);
    });

    it("MODE_RECOVERY: passes with EOA1 + EOA2 (no TPM)", async function () {
        const [eoa1, eoa2] = await ethers.getSigners();

        const Verifier = await ethers.getContractFactory("TPMVerifierSecp256k1");
        const verifier = await Verifier.deploy();
        const tpmKeyId = keyIdFromAddress(eoa1.address);

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:recovery:2eoa"));
        const sig1 = await eoa1.signMessage(ethers.getBytes(userOpHash));
        const sig2 = await eoa2.signMessage(ethers.getBytes(userOpHash));

        const Harness = await ethers.getContractFactory("QuantumAuthAccountHarness");
        const dummyEntryPoint = "0x0000000000000000000000000000000000000001";

        const account = await Harness.deploy(
            dummyEntryPoint,
            eoa1.address,
            eoa2.address,
            await verifier.getAddress(),
            tpmKeyId
        );

        const abi = ethers.AbiCoder.defaultAbiCoder();
        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [MODE_RECOVERY, sig1, sig2, "0x"]);

        const userOp = emptyUserOp(await account.getAddress(), bundle);
        const res = await account.exposedValidateSignature(userOp, userOpHash);
        expect(res).to.equal(0n);
    });

    it("MODE_RECOVERY: fails with only one EOA", async function () {
        const [eoa1, eoa2] = await ethers.getSigners();

        const Verifier = await ethers.getContractFactory("TPMVerifierSecp256k1");
        const verifier = await Verifier.deploy();
        const tpmKeyId = keyIdFromAddress(eoa1.address);

        const userOpHash = ethers.keccak256(ethers.toUtf8Bytes("op:recovery:1eoa"));
        const sig1 = await eoa1.signMessage(ethers.getBytes(userOpHash));

        const Harness = await ethers.getContractFactory("QuantumAuthAccountHarness");
        const dummyEntryPoint = "0x0000000000000000000000000000000000000001";

        const account = await Harness.deploy(
            dummyEntryPoint,
            eoa1.address,
            eoa2.address,
            await verifier.getAddress(),
            tpmKeyId
        );

        const abi = ethers.AbiCoder.defaultAbiCoder();
        const bundle = abi.encode(["uint8", "bytes", "bytes", "bytes"], [MODE_RECOVERY, sig1, "0x", "0x"]);

        const userOp = emptyUserOp(await account.getAddress(), bundle);
        const res = await account.exposedValidateSignature(userOp, userOpHash);
        expect(res).to.equal(1n);
    });
});
