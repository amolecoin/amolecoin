/*
amolecoin daemon
*/
package main

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

import (
	"flag"
	_ "net/http/pprof"
	"os"

	"github.com/amolecoin/amolecoin/src/fiber"
	"github.com/amolecoin/amolecoin/src/readable"
	"github.com/amolecoin/amolecoin/src/amolecoin"
	"github.com/amolecoin/amolecoin/src/util/logging"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.27.0"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// CoinName name of coin
	CoinName = "amolecoin"

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "75a277ef48b5b822073b39b23ede295cf204a4651832f2c75eb6dec40ad84baa564d460c850b9e19dae1d10ca1a524e7862ec14560b801f5335500426e97328801"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "vczdwAwVsyVoGyZCFBoRPyWU3Qq35Ln1XB"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "02f37a1c7ebb2716404c1973cd5b3132a313e77b848fa25628bf1520d05365c5dc"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = "9f911248dbce148e64eb5801c36638af671a2ee3322383517137b2c6a997622a"

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1426562704
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 100000000000000

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
		"64.227.22.96:9982",
		"104.248.121.138:9982",
		"64.227.28.104:9982",
	}

	nodeConfig = amolecoin.NewNodeConfig(ConfigMode, fiber.NodeConfig{
		CoinName:            CoinName,
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "https://raw.githubusercontent.com/M3ND3X/amolenodes/master/nodes.txt",
		Port:                9982,
		WebInterfacePort:    9981,
		DataDirectory:       "$HOME/.amolecoin",

		UnconfirmedBurnFactor:          10,
		UnconfirmedMaxTransactionSize:  32768,
		UnconfirmedMaxDropletPrecision: 3,
		CreateBlockBurnFactor:          10,
		CreateBlockMaxTransactionSize:  32768,
		CreateBlockMaxDropletPrecision: 3,
		MaxBlockTransactionsSize:       32768,

		DisplayName:           "Amolecoin",
		Ticker:                "AMC",
		CoinHoursName:         "Coin Hours",
		CoinHoursNameSingular: "Coin Hour",
		CoinHoursTicker:       "ACH",
		ExplorerURL:           "https://explorer.amolecoin.com",
		VersionURL:            "https://version.amolecoin.com/amolecoin/version.txt",
		Bip44Coin:             8000,
	})

	parseFlags = true
)

func init() {
	nodeConfig.RegisterFlags()
}

func main() {
	if parseFlags {
		flag.Parse()
	}

	// create a new fiber coin instance
	coin := amolecoin.NewCoin(amolecoin.Config{
		Node: nodeConfig,
		Build: readable.BuildInfo{
			Version: Version,
			Commit:  Commit,
			Branch:  Branch,
		},
	}, logger)

	// parse config values
	if err := coin.ParseConfig(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// run fiber coin node
	if err := coin.Run(); err != nil {
		os.Exit(1)
	}
}
