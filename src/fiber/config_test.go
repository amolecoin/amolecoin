package fiber

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/amolecoin/amolecoin/src/cipher/bip44"
)

// TODO(therealssj): write better tests
func TestNewConfig(t *testing.T) {
	coinConfig, err := NewConfig("test.fiber.toml", "./testdata")
	require.NoError(t, err)
	require.Equal(t, Config{
		Node: NodeConfig{
			GenesisSignatureStr: "eb10468d10054d15f2b6f8946cd46797779aa20a7617ceb4be884189f219bc9a164e56a5b9f7bec392a804ff3740210348d73db77a37adb542a8e08d429ac92700",
			GenesisAddressStr:   "vczdwAwVsyVoGyZCFBoRPyWU3Qq35Ln1XB",
			BlockchainPubkeyStr: "02f37a1c7ebb2716404c1973cd5b3132a313e77b848fa25628bf1520d05365c5dc",
			BlockchainSeckeyStr: "",
			GenesisTimestamp:    1426562704,
			GenesisCoinVolume:   100e12,
			DefaultConnections: []string{
				"118.178.135.93:9982",
				"47.88.33.156:9982",
				"104.237.142.206:9982",
				"176.58.126.224:9982",
				"172.104.85.6:9982",
				"139.162.7.132:9982",
			},
			Port:                           9982,
			PeerListURL:                    "https://raw.githubusercontent.com/M3ND3X/amolenodes/master/nodes.txt",
			WebInterfacePort:               9981,
			UnconfirmedBurnFactor:          10,
			UnconfirmedMaxTransactionSize:  777,
			UnconfirmedMaxDropletPrecision: 3,
			CreateBlockBurnFactor:          9,
			CreateBlockMaxTransactionSize:  1234,
			CreateBlockMaxDropletPrecision: 4,
			MaxBlockTransactionsSize:       1111,
			DisplayName:                    "Testcoin",
			Ticker:                         "TST",
			CoinHoursName:                  "Testcoin Hours",
			CoinHoursNameSingular:          "Testcoin Hour",
			CoinHoursTicker:                "TCH",
			ExplorerURL:                    "https://explorer.testcoin.com",
			VersionURL:                     "https://version.testcoin.com/testcoin/version.txt",
			Bip44Coin:                      bip44.CoinTypeAmolecoin,
		},
		Params: ParamsConfig{
			MaxCoinSupply:           1e8,
			UnlockAddressRate:       5,
			InitialUnlockedCount:    25,
			UnlockTimeInterval:      60 * 60 * 24 * 365,
			UserBurnFactor:          3,
			UserMaxTransactionSize:  999,
			UserMaxDropletPrecision: 2,
		},
	}, coinConfig)
}
