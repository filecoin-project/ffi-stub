package generated

// FilAggregate function as declared in filecoin-ffi/filcrypto.h:349
func FilAggregate(flattenedSignaturesPtr []byte, flattenedSignaturesLen uint) *FilAggregateResponse {
	var r FilAggregateResponse
	return &r
}

// FilAggregateSealProofs function as declared in filecoin-ffi/filcrypto.h:352
func FilAggregateSealProofs(registeredProof FilRegisteredSealProof, registeredAggregation FilRegisteredAggregationProof, commRsPtr []Fil32ByteArray, commRsLen uint, seedsPtr []Fil32ByteArray, seedsLen uint, sealCommitResponsesPtr []FilSealCommitPhase2Response, sealCommitResponsesLen uint) *FilAggregateProof {
	var r FilAggregateProof
	return &r
}

// FilClearCache function as declared in filecoin-ffi/filcrypto.h:361
func FilClearCache(sectorSize uint64, cacheDirPath string) *FilClearCacheResponse {
	var r FilClearCacheResponse
	return &r
}

// FilCreateZeroSignature function as declared in filecoin-ffi/filcrypto.h:368
func FilCreateZeroSignature() *FilZeroSignatureResponse {
	var r FilZeroSignatureResponse
	return &r
}

// FilDestroyAggregateProof function as declared in filecoin-ffi/filcrypto.h:374
func FilDestroyAggregateProof(ptr *FilAggregateProof) {
	// NOOP
}

// FilDestroyAggregateResponse function as declared in filecoin-ffi/filcrypto.h:376
func FilDestroyAggregateResponse(ptr *FilAggregateResponse) {
	// NOOP
}

// FilDestroyClearCacheResponse function as declared in filecoin-ffi/filcrypto.h:378
func FilDestroyClearCacheResponse(ptr *FilClearCacheResponse) {
	// NOOP
}

// FilDestroyFauxrepResponse function as declared in filecoin-ffi/filcrypto.h:380
func FilDestroyFauxrepResponse(*FilFauxRepResponse) {
	// NOOP
}

// FilDestroyFinalizeTicketResponse function as declared in filecoin-ffi/filcrypto.h:382
func FilDestroyFinalizeTicketResponse(ptr *FilFinalizeTicketResponse) {
	// NOOP
}

// FilDestroyGenerateDataCommitmentResponse function as declared in filecoin-ffi/filcrypto.h:384
func FilDestroyGenerateDataCommitmentResponse(ptr *FilGenerateDataCommitmentResponse) {
	// NOOP
}

// FilDestroyGenerateFallbackSectorChallengesResponse function as declared in filecoin-ffi/filcrypto.h:386
func FilDestroyGenerateFallbackSectorChallengesResponse(ptr *FilGenerateFallbackSectorChallengesResponse) {
	// NOOP
}

// FilDestroyGeneratePieceCommitmentResponse function as declared in filecoin-ffi/filcrypto.h:388
func FilDestroyGeneratePieceCommitmentResponse(ptr *FilGeneratePieceCommitmentResponse) {
	// NOOP
}

// FilDestroyGenerateSingleVanillaProofResponse function as declared in filecoin-ffi/filcrypto.h:390
func FilDestroyGenerateSingleVanillaProofResponse(ptr *FilGenerateSingleVanillaProofResponse) {
	// NOOP
}

// FilDestroyGenerateWindowPostResponse function as declared in filecoin-ffi/filcrypto.h:392
func FilDestroyGenerateWindowPostResponse(ptr *FilGenerateWindowPoStResponse) {
	// NOOP
}

// FilDestroyGenerateWinningPostResponse function as declared in filecoin-ffi/filcrypto.h:394
func FilDestroyGenerateWinningPostResponse(ptr *FilGenerateWinningPoStResponse) {
	// NOOP
}

// FilDestroyGenerateWinningPostSectorChallenge function as declared in filecoin-ffi/filcrypto.h:396
func FilDestroyGenerateWinningPostSectorChallenge(ptr *FilGenerateWinningPoStSectorChallenge) {
	// NOOP
}

// FilDestroyGpuDeviceResponse function as declared in filecoin-ffi/filcrypto.h:398
func FilDestroyGpuDeviceResponse(ptr *FilGpuDeviceResponse) {
	// NOOP
}

// FilDestroyHashResponse function as declared in filecoin-ffi/filcrypto.h:400
func FilDestroyHashResponse(ptr *FilHashResponse) {
	// NOOP
}

// FilDestroyInitLogFdResponse function as declared in filecoin-ffi/filcrypto.h:402
func FilDestroyInitLogFdResponse(ptr *FilInitLogFdResponse) {
	// NOOP
}

// FilDestroyPrivateKeyGenerateResponse function as declared in filecoin-ffi/filcrypto.h:404
func FilDestroyPrivateKeyGenerateResponse(ptr *FilPrivateKeyGenerateResponse) {
	// NOOP
}

// FilDestroyPrivateKeyPublicKeyResponse function as declared in filecoin-ffi/filcrypto.h:406
func FilDestroyPrivateKeyPublicKeyResponse(ptr *FilPrivateKeyPublicKeyResponse) {
	// NOOP
}

// FilDestroyPrivateKeySignResponse function as declared in filecoin-ffi/filcrypto.h:408
func FilDestroyPrivateKeySignResponse(ptr *FilPrivateKeySignResponse) {
	// NOOP
}

// FilDestroySealCommitPhase1Response function as declared in filecoin-ffi/filcrypto.h:410
func FilDestroySealCommitPhase1Response(ptr *FilSealCommitPhase1Response) {
	// NOOP
}

// FilDestroySealCommitPhase2Response function as declared in filecoin-ffi/filcrypto.h:412
func FilDestroySealCommitPhase2Response(ptr *FilSealCommitPhase2Response) {
	// NOOP
}

// FilDestroySealPreCommitPhase1Response function as declared in filecoin-ffi/filcrypto.h:414
func FilDestroySealPreCommitPhase1Response(ptr *FilSealPreCommitPhase1Response) {
	// NOOP
}

// FilDestroySealPreCommitPhase2Response function as declared in filecoin-ffi/filcrypto.h:416
func FilDestroySealPreCommitPhase2Response(ptr *FilSealPreCommitPhase2Response) {
	// NOOP
}

// FilDestroyStringResponse function as declared in filecoin-ffi/filcrypto.h:418
func FilDestroyStringResponse(ptr *FilStringResponse) {
	// NOOP
}

// FilDestroyUnsealRangeResponse function as declared in filecoin-ffi/filcrypto.h:420
func FilDestroyUnsealRangeResponse(ptr *FilUnsealRangeResponse) {
	// NOOP
}

// FilDestroyVerifyAggregateSealResponse function as declared in filecoin-ffi/filcrypto.h:426
func FilDestroyVerifyAggregateSealResponse(ptr *FilVerifyAggregateSealProofResponse) {
	// NOOP
}

// FilDestroyVerifySealResponse function as declared in filecoin-ffi/filcrypto.h:432
func FilDestroyVerifySealResponse(ptr *FilVerifySealResponse) {
	// NOOP
}

// FilDestroyVerifyWindowPostResponse function as declared in filecoin-ffi/filcrypto.h:434
func FilDestroyVerifyWindowPostResponse(ptr *FilVerifyWindowPoStResponse) {
	// NOOP
}

// FilDestroyVerifyWinningPostResponse function as declared in filecoin-ffi/filcrypto.h:440
func FilDestroyVerifyWinningPostResponse(ptr *FilVerifyWinningPoStResponse) {
	// NOOP
}

// FilDestroyWriteWithAlignmentResponse function as declared in filecoin-ffi/filcrypto.h:442
func FilDestroyWriteWithAlignmentResponse(ptr *FilWriteWithAlignmentResponse) {
	// NOOP
}

// FilDestroyWriteWithoutAlignmentResponse function as declared in filecoin-ffi/filcrypto.h:444
func FilDestroyWriteWithoutAlignmentResponse(ptr *FilWriteWithoutAlignmentResponse) {
	// NOOP
}

// FilDestroyZeroSignatureResponse function as declared in filecoin-ffi/filcrypto.h:446
func FilDestroyZeroSignatureResponse(ptr *FilZeroSignatureResponse) {
	// NOOP
}

// FilDropSignature function as declared in filecoin-ffi/filcrypto.h:451
func FilDropSignature(sig []byte) {
	// NOOP
}

// FilFauxrep function as declared in filecoin-ffi/filcrypto.h:453
func FilFauxrep(registeredProof FilRegisteredSealProof, cacheDirPath string, sealedSectorPath string) *FilFauxRepResponse {
	var r FilFauxRepResponse
	return &r
}

// FilFauxrep2 function as declared in filecoin-ffi/filcrypto.h:457
func FilFauxrep2(registeredProof FilRegisteredSealProof, cacheDirPath string, existingPAuxPath string) *FilFauxRepResponse {
	var r FilFauxRepResponse
	return &r
}

// FilGenerateDataCommitment function as declared in filecoin-ffi/filcrypto.h:464
func FilGenerateDataCommitment(registeredProof FilRegisteredSealProof, piecesPtr []FilPublicPieceInfo, piecesLen uint) *FilGenerateDataCommitmentResponse {
	var r FilGenerateDataCommitmentResponse
	return &r
}

// FilGenerateFallbackSectorChallenges function as declared in filecoin-ffi/filcrypto.h:472
func FilGenerateFallbackSectorChallenges(registeredProof FilRegisteredPoStProof, randomness Fil32ByteArray, sectorIdsPtr []uint64, sectorIdsLen uint, proverId Fil32ByteArray) *FilGenerateFallbackSectorChallengesResponse {
	var r FilGenerateFallbackSectorChallengesResponse
	return &r
}

// FilGeneratePieceCommitment function as declared in filecoin-ffi/filcrypto.h:482
func FilGeneratePieceCommitment(registeredProof FilRegisteredSealProof, pieceFdRaw int32, unpaddedPieceSize uint64) *FilGeneratePieceCommitmentResponse {
	var r FilGeneratePieceCommitmentResponse
	return &r
}

// FilGenerateSingleVanillaProof function as declared in filecoin-ffi/filcrypto.h:490
func FilGenerateSingleVanillaProof(replica FilPrivateReplicaInfo, challengesPtr []uint64, challengesLen uint) *FilGenerateSingleVanillaProofResponse {
	var r FilGenerateSingleVanillaProofResponse
	return &r
}

// FilGenerateWindowPost function as declared in filecoin-ffi/filcrypto.h:498
func FilGenerateWindowPost(randomness Fil32ByteArray, replicasPtr []FilPrivateReplicaInfo, replicasLen uint, proverId Fil32ByteArray) *FilGenerateWindowPoStResponse {
	var r FilGenerateWindowPoStResponse
	return &r
}

// FilGenerateWindowPostWithVanilla function as declared in filecoin-ffi/filcrypto.h:507
func FilGenerateWindowPostWithVanilla(registeredProof FilRegisteredPoStProof, randomness Fil32ByteArray, proverId Fil32ByteArray, vanillaProofsPtr []FilVanillaProof, vanillaProofsLen uint) *FilGenerateWindowPoStResponse {
	var r FilGenerateWindowPoStResponse
	return &r
}

// FilGenerateWinningPost function as declared in filecoin-ffi/filcrypto.h:517
func FilGenerateWinningPost(randomness Fil32ByteArray, replicasPtr []FilPrivateReplicaInfo, replicasLen uint, proverId Fil32ByteArray) *FilGenerateWinningPoStResponse {
	var r FilGenerateWinningPoStResponse
	return &r
}

// FilGenerateWinningPostSectorChallenge function as declared in filecoin-ffi/filcrypto.h:526
func FilGenerateWinningPostSectorChallenge(registeredProof FilRegisteredPoStProof, randomness Fil32ByteArray, sectorSetLen uint64, proverId Fil32ByteArray) *FilGenerateWinningPoStSectorChallenge {
	var r FilGenerateWinningPoStSectorChallenge
	return &r
}

// FilGenerateWinningPostWithVanilla function as declared in filecoin-ffi/filcrypto.h:535
func FilGenerateWinningPostWithVanilla(registeredProof FilRegisteredPoStProof, randomness Fil32ByteArray, proverId Fil32ByteArray, vanillaProofsPtr []FilVanillaProof, vanillaProofsLen uint) *FilGenerateWinningPoStResponse {
	var r FilGenerateWinningPoStResponse
	return &r
}

// FilGetGpuDevices function as declared in filecoin-ffi/filcrypto.h:544
func FilGetGpuDevices() *FilGpuDeviceResponse {
	var r FilGpuDeviceResponse
	return &r
}

// FilGetMaxUserBytesPerStagedSector function as declared in filecoin-ffi/filcrypto.h:550
func FilGetMaxUserBytesPerStagedSector(registeredProof FilRegisteredSealProof) uint64 {
	return 0
}

// FilGetPostCircuitIdentifier function as declared in filecoin-ffi/filcrypto.h:556
func FilGetPostCircuitIdentifier(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetPostParamsCid function as declared in filecoin-ffi/filcrypto.h:562
func FilGetPostParamsCid(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetPostParamsPath function as declared in filecoin-ffi/filcrypto.h:569
func FilGetPostParamsPath(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetPostVerifyingKeyCid function as declared in filecoin-ffi/filcrypto.h:575
func FilGetPostVerifyingKeyCid(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetPostVerifyingKeyPath function as declared in filecoin-ffi/filcrypto.h:582
func FilGetPostVerifyingKeyPath(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetPostVersion function as declared in filecoin-ffi/filcrypto.h:588
func FilGetPostVersion(registeredProof FilRegisteredPoStProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetSealCircuitIdentifier function as declared in filecoin-ffi/filcrypto.h:594
func FilGetSealCircuitIdentifier(registeredProof FilRegisteredSealProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetSealParamsCid function as declared in filecoin-ffi/filcrypto.h:600
func FilGetSealParamsCid(registeredProof FilRegisteredSealProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetSealParamsPath function as declared in filecoin-ffi/filcrypto.h:607
func FilGetSealParamsPath(registeredProof FilRegisteredSealProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetSealVerifyingKeyCid function as declared in filecoin-ffi/filcrypto.h:613
func FilGetSealVerifyingKeyCid(registeredProof FilRegisteredSealProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetSealVerifyingKeyPath function as declared in filecoin-ffi/filcrypto.h:620
func FilGetSealVerifyingKeyPath(registeredProof FilRegisteredSealProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilGetSealVersion function as declared in filecoin-ffi/filcrypto.h:626
func FilGetSealVersion(registeredProof FilRegisteredSealProof) *FilStringResponse {
	var r FilStringResponse
	return &r
}

// FilHash function as declared in filecoin-ffi/filcrypto.h:636
func FilHash(messagePtr []byte, messageLen uint) *FilHashResponse {
	var r FilHashResponse
	return &r
}

// FilHashVerify function as declared in filecoin-ffi/filcrypto.h:650
func FilHashVerify(signaturePtr []byte, flattenedMessagesPtr []byte, flattenedMessagesLen uint, messageSizesPtr []uint, messageSizesLen uint, flattenedPublicKeysPtr []byte, flattenedPublicKeysLen uint) int32 {
	return 0
}

// FilInitLogFd function as declared in filecoin-ffi/filcrypto.h:667
func FilInitLogFd(logFd int32) *FilInitLogFdResponse {
	var r FilInitLogFdResponse
	return &r
}

// FilPrivateKeyGenerate function as declared in filecoin-ffi/filcrypto.h:672
func FilPrivateKeyGenerate() *FilPrivateKeyGenerateResponse {
	var r FilPrivateKeyGenerateResponse
	return &r
}

// FilPrivateKeyGenerateWithSeed function as declared in filecoin-ffi/filcrypto.h:685
func FilPrivateKeyGenerateWithSeed(rawSeed Fil32ByteArray) *FilPrivateKeyGenerateResponse {
	var r FilPrivateKeyGenerateResponse
	return &r
}

// FilPrivateKeyPublicKey function as declared in filecoin-ffi/filcrypto.h:696
func FilPrivateKeyPublicKey(rawPrivateKeyPtr []byte) *FilPrivateKeyPublicKeyResponse {
	var r FilPrivateKeyPublicKeyResponse
	return &r
}

// FilPrivateKeySign function as declared in filecoin-ffi/filcrypto.h:709
func FilPrivateKeySign(rawPrivateKeyPtr []byte, messagePtr []byte, messageLen uint) *FilPrivateKeySignResponse {
	var r FilPrivateKeySignResponse
	return &r
}

// FilSealCommitPhase1 function as declared in filecoin-ffi/filcrypto.h:717
func FilSealCommitPhase1(registeredProof FilRegisteredSealProof, commR Fil32ByteArray, commD Fil32ByteArray, cacheDirPath string, replicaPath string, sectorId uint64, proverId Fil32ByteArray, ticket Fil32ByteArray, seed Fil32ByteArray, piecesPtr []FilPublicPieceInfo, piecesLen uint) *FilSealCommitPhase1Response {
	var r FilSealCommitPhase1Response
	return &r
}

// FilSealCommitPhase2 function as declared in filecoin-ffi/filcrypto.h:729
func FilSealCommitPhase2(sealCommitPhase1OutputPtr []byte, sealCommitPhase1OutputLen uint, sectorId uint64, proverId Fil32ByteArray) *FilSealCommitPhase2Response {
	var r FilSealCommitPhase2Response
	return &r
}

// FilSealPreCommitPhase1 function as declared in filecoin-ffi/filcrypto.h:738
func FilSealPreCommitPhase1(registeredProof FilRegisteredSealProof, cacheDirPath string, stagedSectorPath string, sealedSectorPath string, sectorId uint64, proverId Fil32ByteArray, ticket Fil32ByteArray, piecesPtr []FilPublicPieceInfo, piecesLen uint) *FilSealPreCommitPhase1Response {
	var r FilSealPreCommitPhase1Response
	return &r
}

// FilSealPreCommitPhase2 function as declared in filecoin-ffi/filcrypto.h:752
func FilSealPreCommitPhase2(sealPreCommitPhase1OutputPtr []byte, sealPreCommitPhase1OutputLen uint, cacheDirPath string, sealedSectorPath string) *FilSealPreCommitPhase2Response {
	var r FilSealPreCommitPhase2Response
	return &r
}

// FilUnsealRange function as declared in filecoin-ffi/filcrypto.h:760
func FilUnsealRange(registeredProof FilRegisteredSealProof, cacheDirPath string, sealedSectorFdRaw int32, unsealOutputFdRaw int32, sectorId uint64, proverId Fil32ByteArray, ticket Fil32ByteArray, commD Fil32ByteArray, unpaddedByteIndex uint64, unpaddedBytesAmount uint64) *FilUnsealRangeResponse {
	var r FilUnsealRangeResponse
	return &r
}

// FilVerify function as declared in filecoin-ffi/filcrypto.h:782
func FilVerify(signaturePtr []byte, flattenedDigestsPtr []byte, flattenedDigestsLen uint, flattenedPublicKeysPtr []byte, flattenedPublicKeysLen uint) int32 {
	return 0
}

// FilVerifyAggregateSealProof function as declared in filecoin-ffi/filcrypto.h:792
func FilVerifyAggregateSealProof(registeredProof FilRegisteredSealProof, registeredAggregation FilRegisteredAggregationProof, proverId Fil32ByteArray, proofPtr []byte, proofLen uint, commitInputsPtr []FilAggregationInputs, commitInputsLen uint) *FilVerifyAggregateSealProofResponse {
	var r FilVerifyAggregateSealProofResponse
	return &r
}

// FilVerifySeal function as declared in filecoin-ffi/filcrypto.h:804
func FilVerifySeal(registeredProof FilRegisteredSealProof, commR Fil32ByteArray, commD Fil32ByteArray, proverId Fil32ByteArray, ticket Fil32ByteArray, seed Fil32ByteArray, sectorId uint64, proofPtr []byte, proofLen uint) *FilVerifySealResponse {
	var r FilVerifySealResponse
	return &r
}

// FilVerifyWindowPost function as declared in filecoin-ffi/filcrypto.h:817
func FilVerifyWindowPost(randomness Fil32ByteArray, replicasPtr []FilPublicReplicaInfo, replicasLen uint, proofsPtr []FilPoStProof, proofsLen uint, proverId Fil32ByteArray) *FilVerifyWindowPoStResponse {
	var r FilVerifyWindowPoStResponse
	return &r
}

// FilVerifyWinningPost function as declared in filecoin-ffi/filcrypto.h:827
func FilVerifyWinningPost(randomness Fil32ByteArray, replicasPtr []FilPublicReplicaInfo, replicasLen uint, proofsPtr []FilPoStProof, proofsLen uint, proverId Fil32ByteArray) *FilVerifyWinningPoStResponse {
	var r FilVerifyWinningPoStResponse
	return &r
}

// FilWriteWithAlignment function as declared in filecoin-ffi/filcrypto.h:838
func FilWriteWithAlignment(registeredProof FilRegisteredSealProof, srcFd int32, srcSize uint64, dstFd int32, existingPieceSizesPtr []uint64, existingPieceSizesLen uint) *FilWriteWithAlignmentResponse {
	var r FilWriteWithAlignmentResponse
	return &r
}

// FilWriteWithoutAlignment function as declared in filecoin-ffi/filcrypto.h:849
func FilWriteWithoutAlignment(registeredProof FilRegisteredSealProof, srcFd int32, srcSize uint64, dstFd int32) *FilWriteWithoutAlignmentResponse {
	var r FilWriteWithoutAlignmentResponse
	return &r
}
