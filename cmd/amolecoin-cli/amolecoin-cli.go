/*
cli is a command line client for interacting with a amolecoin node and offline wallet management
*/
package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/amolecoin/amolecoin/src/cli"
	"github.com/amolecoin/amolecoin/src/util/logging"
)

func main() {
	logging.SetLevel(logrus.WarnLevel)

	cfg, err := cli.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	amoleCLI, err := cli.NewCLI(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := amoleCLI.Execute(); err != nil {
		os.Exit(1)
	}
}
