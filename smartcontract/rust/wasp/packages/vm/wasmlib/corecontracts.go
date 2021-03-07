// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmlib

const CoreAccounts = ScHname(0x3c4b5e02)
const CoreAccountsFuncDeposit = ScHname(0xbdc9102d)
const CoreAccountsFuncWithdrawToAddress = ScHname(0x26608cb5)
const CoreAccountsFuncWithdrawToChain = ScHname(0x437bc026)
const CoreAccountsViewAccounts = ScHname(0x3c4b5e02)
const CoreAccountsViewBalance = ScHname(0x84168cb4)
const CoreAccountsViewTotalAssets = ScHname(0xfab0f8d2)

const CoreAccountsParamAgentID = Key("a")

const CoreBlob = ScHname(0xfd91bc63)
const CoreBlobFuncStoreBlob = ScHname(0xddd4c281)
const CoreBlobViewGetBlobField = ScHname(0x1f448130)
const CoreBlobViewGetBlobInfo = ScHname(0xfde4ab46)
const CoreBlobViewListBlobs = ScHname(0x62ca7990)

const CoreBlobParamField = Key("field")
const CoreBlobParamHash = Key("hash")

const CoreEventlog = ScHname(0x661aa7d8)
const CoreEventlogViewGetNumRecords = ScHname(0x2f4b4a8c)
const CoreEventlogViewGetRecords = ScHname(0xd01a8085)

const CoreEventlogParamContractHname = Key("contractHname")
const CoreEventlogParamFromTs = Key("fromTs")
const CoreEventlogParamMaxLastRecords = Key("maxLastRecords")
const CoreEventlogParamToTs = Key("toTs")

const CoreRoot = ScHname(0xcebf5908)
const CoreRootFuncClaimChainOwnership = ScHname(0x03ff0fc0)
const CoreRootFuncDelegateChainOwnership = ScHname(0x93ecb6ad)
const CoreRootFuncDeployContract = ScHname(0x28232c27)
const CoreRootFuncGrantDeployPermission = ScHname(0xf440263a)
const CoreRootFuncRevokeDeployPermission = ScHname(0x850744f1)
const CoreRootFuncSetContractFee = ScHname(0x8421a42b)
const CoreRootFuncSetDefaultFee = ScHname(0x3310ecd0)
const CoreRootViewFindContract = ScHname(0xc145ca00)
const CoreRootViewGetChainInfo = ScHname(0x434477e2)
const CoreRootViewGetFeeInfo = ScHname(0x9fe54b48)

const CoreRootParamChainOwner = Key("$$owner$$")
const CoreRootParamDeployer = Key("$$deployer$$")
const CoreRootParamDescription = Key("$$description$$")
const CoreRootParamHname = Key("$$hname$$")
const CoreRootParamName = Key("$$name$$")
const CoreRootParamOwnerFee = Key("$$ownerfee$$")
const CoreRootParamProgramHash = Key("$$proghash$$")
const CoreRootParamValidatorFee = Key("$$validatorfee$$")
