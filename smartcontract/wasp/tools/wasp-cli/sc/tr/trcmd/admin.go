// +build ignore

package trcmd

import (
	"fmt"
	"os"

	"github.com/iotaledger/wasp/tools/wasp-cli/sc/tr"
	"github.com/iotaledger/wasp/tools/wasp-cli/wallet"
)

func adminCmd(args []string) {
	if len(args) < 1 {
		adminUsage()
	}

	switch args[0] {
	case "deploy":
		log.Check(tr.Config.Deploy(wallet.Load().SignatureScheme()))

	default:
		adminUsage()
	}
}

func adminUsage() {
	fmt.Printf("Usage: %s tr admin [deploy]\n", os.Args[0])
	os.Exit(1)
}
