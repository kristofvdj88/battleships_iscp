// +build ignore

package dashboardcmd

import (
	"fmt"
	"os"

	"github.com/iotaledger/wasp/tools/wasp-cli/dashboard"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/dwf"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/dwf/dwfdashboard"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/fa"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/fa/fadashboard"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/fr"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/fr/frdashboard"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/tr"
	"github.com/iotaledger/wasp/tools/wasp-cli/sc/tr/trdashboard"
	"github.com/spf13/pflag"
)

func InitCommands(commands map[string]func([]string), flags *pflag.FlagSet) {
	commands["dashboard"] = cmd
}

func cmd(args []string) {
	listenAddr := ":10000"
	if len(args) > 0 {
		if len(args) != 1 {
			fmt.Printf("Usage: %s dashboard [listen-address]\n", os.Args[0])
			os.Exit(1)
		}
		listenAddr = args[0]
	}

	scs := make([]dashboard.SCDashboard, 0)
	if fr.Config.IsAvailable() {
		scs = append(scs, frdashboard.Dashboard())
		fmt.Printf("FairRoulette: %s\n", fr.Config.Href())
	} else {
		fmt.Println("FairRoulette not available")
	}
	if fa.Config.IsAvailable() {
		scs = append(scs, fadashboard.Dashboard())
		fmt.Printf("FairAuction: %s\n", fa.Config.Href())
	} else {
		fmt.Println("FairAuction not available")
	}
	if tr.Config.IsAvailable() {
		scs = append(scs, trdashboard.Dashboard())
		fmt.Printf("TokenRegistry: %s\n", tr.Config.Href())
	} else {
		fmt.Println("TokenRegistry not available")
	}
	if dwf.Config.IsAvailable() {
		fmt.Printf("DonateWithFeedback: %s\n", dwf.Config.Href())
		scs = append(scs, dwfdashboard.Dashboard())
	} else {
		fmt.Println("DonateWithFeedback not available")
	}

	dashboard.StartServer(listenAddr, scs)
}
