// +build ignore

package dwf

import (
	"github.com/iotaledger/wasp/packages/vm/examples/donatewithfeedback/dwfclient"
	"github.com/iotaledger/wasp/packages/vm/examples/donatewithfeedback/dwfimpl"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc"
	"github.com/iotaledger/wasp/tools/wasp-cli/wallet"
)

var Config = &sc.Config{
	ShortName:   "dwf",
	Name:        "DonateWithFeedback",
	ProgramHash: dwfimpl.ProgramHash,
}

func Client() *dwfclient.DWFClient {
	return dwfclient.NewClient(Config.MakeClient(wallet.Load().SignatureScheme()))
}
