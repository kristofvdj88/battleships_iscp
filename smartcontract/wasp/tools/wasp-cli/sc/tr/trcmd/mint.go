// +build ignore

package trcmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/iotaledger/wasp/packages/vm/examples/tokenregistry/trclient"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/tr"
	"github.com/iotaledger/wasp/tools/wasp-cli/wallet"
)

func mintCmd(args []string) {
	if len(args) != 2 {
		fmt.Printf("Usage: %s tr mint <description> <amount>\n", os.Args[0])
		os.Exit(1)
	}

	description := args[0]

	amount, err := strconv.Atoi(args[1])
	log.Check(err)

	client := tr.Client()
	tx, err := client.MintAndRegister(trclient.MintAndRegisterParams{
		Supply:      int64(amount),
		MintTarget:  wallet.Load().Address(),
		Description: description,
	})
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	fmt.Printf("Minted %d tokens of color %s into address %s.\n"+
		"Metadata of the supply: '%s'\n"+
		"Metadata was sent to TokenRegistry SC at %s\n",
		amount, tx.ID().String(), client.OwnerAddress().String(), description, tr.Config.Address().String())
}
