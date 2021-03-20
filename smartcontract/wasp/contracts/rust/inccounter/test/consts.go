// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package test

import (
	"github.com/iotaledger/wasp/packages/coretypes"
)

const ScName = "inccounter"
const ScHname = coretypes.Hname(0xaf2438e9)

const ParamCounter = "counter"
const ParamNumRepeats = "numRepeats"

const VarCounter = "counter"
const VarInt1 = "int1"
const VarIntArray1 = "intArray1"
const VarNumRepeats = "numRepeats"
const VarString1 = "string1"
const VarStringArray1 = "stringArray1"

const FuncCallIncrement = "callIncrement"
const FuncCallIncrementRecurse5x = "callIncrementRecurse5x"
const FuncIncrement = "increment"
const FuncInit = "init"
const FuncLocalStateInternalCall = "localStateInternalCall"
const FuncLocalStatePost = "localStatePost"
const FuncLocalStateSandboxCall = "localStateSandboxCall"
const FuncPostIncrement = "postIncrement"
const FuncRepeatMany = "repeatMany"
const FuncResultsTest = "resultsTest"
const FuncStateTest = "stateTest"
const FuncWhenMustIncrement = "whenMustIncrement"
const ViewGetCounter = "getCounter"
const ViewResultsCheck = "resultsCheck"
const ViewStateCheck = "stateCheck"

const HFuncCallIncrement = coretypes.Hname(0xeb5dcacd)
const HFuncCallIncrementRecurse5x = coretypes.Hname(0x8749fbff)
const HFuncIncrement = coretypes.Hname(0xd351bd12)
const HFuncInit = coretypes.Hname(0x1f44d644)
const HFuncLocalStateInternalCall = coretypes.Hname(0xecfc5d33)
const HFuncLocalStatePost = coretypes.Hname(0x3fd54d13)
const HFuncLocalStateSandboxCall = coretypes.Hname(0x7bd22c53)
const HFuncPostIncrement = coretypes.Hname(0x81c772f5)
const HFuncRepeatMany = coretypes.Hname(0x4ff450d3)
const HFuncResultsTest = coretypes.Hname(0xd0544634)
const HFuncStateTest = coretypes.Hname(0x41830d59)
const HFuncWhenMustIncrement = coretypes.Hname(0xb4c3e7a6)
const HViewGetCounter = coretypes.Hname(0xb423e607)
const HViewResultsCheck = coretypes.Hname(0xa39ac571)
const HViewStateCheck = coretypes.Hname(0xaafeb10a)
