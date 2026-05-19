// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.27;

import {MultisigBuilder} from "zeus-templates/templates/MultisigBuilder.sol";
import "../Env.sol";

import {ImageAllowlist} from "../../../src/ImageAllowlist.sol";
import {IImageAllowlist} from "../../../src/interfaces/IImageAllowlist.sol";

/**
 * Purpose:
 *      Register the cos-125-lts CS image's PCR measurements on the
 *      ImageAllowlist contract so eigenx-kms's policy gate
 *      (PolicyChecker.checkPCRAllowlist) accepts attestations from VMs
 *      running the new base image.
 *
 * Context:
 *      Layr-Labs/go-tpm-tools#32 bumps the launcher's CS-image base
 *      from cos-tdx-113-18244-582-80 to cos-125-19216-395-7. Bumping
 *      the base image changes every PCR the TPM measures during boot,
 *      so the allowlist must be updated in lockstep — otherwise every
 *      VM on the new image fails KMS attestation and self-terminates.
 *
 *      Run order:
 *        1. Build the new CS image (via go-tpm-tools' deploy-builder
 *           workflow) and capture its PCR measurements via the
 *           pcr-capture binary.
 *        2. Replace the TODO_CAPTURE_PCRS_HERE placeholders below
 *           with the captured values. PCRs MUST be sorted by index
 *           ascending (the contract reverts with NotSorted() otherwise).
 *        3. Run this script via the compute-ops multisig to enqueue
 *           the addImage transaction.
 *        4. Confirm `ImageAllowlist.isImageAllowed(INTEL_TDX, pcrs)`
 *           returns true for the new image's PCRs before flipping any
 *           production traffic.
 *
 * Operational notes:
 *      - This adds an entry. It does not remove the cos-tdx-113 entry
 *        — keep the older image allowed for the duration of the rollout
 *        so existing VMs continue to attest successfully. Once all
 *        production VMs are off cos-tdx-113, queue a follow-up
 *        removeImage call.
 *      - The Image.version + description fields are advisory only; the
 *        contract keys on hash(pcrs). Use them to make on-chain history
 *        readable.
 *      - If the PCR set differs between TDX and SEV-SNP variants of the
 *        cos-125 image, run separate addImage calls per Platform.
 *        Today (M125) Google publishes a unified cos-125-lts family
 *        that supports both, so a single TDX entry is sufficient unless
 *        we deploy SEV-SNP workloads.
 */
contract AddCos125Image is MultisigBuilder {
    using Env for *;

    function _runAsMultisig() internal override prank(Env.computeOpsMultisig()) {
        IImageAllowlist.PCR[] memory pcrs = _cos125PCRs();
        IImageAllowlist.Image memory image = IImageAllowlist.Image({
            pcrs: pcrs,
            version: "cos-125-19216-395-7",
            description: "Container-Optimized OS 125 LTS Refresh (May 8 2026), TDX-capable"
        });

        Env.proxy.imageAllowlist().addImage(IImageAllowlist.Platform.INTEL_TDX, image);
    }

    /// @dev Replace the PCR values below with the measurements captured
    ///      from a build of go-tpm-tools' deploy-builder workflow against
    ///      cos-125-19216-395-7. The pcr-capture binary writes a JSON
    ///      attestation envelope to gs://${PROVENANCE_BUCKET}/<image>/attestation.json
    ///      with the per-image PCR values.
    ///
    ///      Indices MUST be in strictly ascending order (the contract
    ///      reverts with NotSorted() otherwise). The exact set of PCR
    ///      indices to register is determined by the eigenx-kms
    ///      PolicyChecker, which forwards every PCR present in the
    ///      attestation's TPMClaims.PCRs map. As of writing, that's
    ///      indices 0, 1, 2, 3, 4, 5, 6, 7, 9, 14 — verify against
    ///      pcrMapToContractPCRs in eigenx-kms/pkg/policy/policy.go
    ///      before merging this PR.
    function _cos125PCRs() private pure returns (IImageAllowlist.PCR[] memory) {
        IImageAllowlist.PCR[] memory pcrs = new IImageAllowlist.PCR[](10);
        pcrs[0] = IImageAllowlist.PCR({index: 0, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[1] = IImageAllowlist.PCR({index: 1, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[2] = IImageAllowlist.PCR({index: 2, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[3] = IImageAllowlist.PCR({index: 3, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[4] = IImageAllowlist.PCR({index: 4, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[5] = IImageAllowlist.PCR({index: 5, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[6] = IImageAllowlist.PCR({index: 6, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[7] = IImageAllowlist.PCR({index: 7, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[8] = IImageAllowlist.PCR({index: 9, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        pcrs[9] = IImageAllowlist.PCR({index: 14, value: bytes32(0)}); // TODO_CAPTURE_PCRS_HERE
        return pcrs;
    }

    function testScript() public virtual {
        execute();
        _validateState();
    }

    function _validateState() internal view {
        IImageAllowlist.PCR[] memory pcrs = _cos125PCRs();
        // Once PCRs are filled in, this check should pass on a forked
        // mainnet test. With placeholder zeros it will pass trivially
        // (zero hash is a valid bytes32 key), so the PR review must not
        // rely on this assertion alone — manual verification of the
        // captured PCR values against the on-chain entry post-broadcast
        // is the real check.
        assertTrue(
            Env.proxy.imageAllowlist().isImageAllowed(IImageAllowlist.Platform.INTEL_TDX, pcrs),
            "cos-125 image not registered after addImage"
        );
    }
}
