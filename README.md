![amolecoin logo](https://avatars1.githubusercontent.com/u/47041087?s=200&v=4)

# Amolecoin

[![Build Status](https://travis-ci.com/amolecoin/amolecoin.svg)](https://travis-ci.com/amolecoin/amolecoin)
[![GoDoc](https://godoc.org/github.com/amolecoin/amolecoin?status.svg)](https://godoc.org/github.com/amolecoin/amolecoin)
[![Go Report Card](https://goreportcard.com/badge/github.com/amolecoin/amolecoin)](https://goreportcard.com/report/github.com/amolecoin/amolecoin)

Amolecoin is a next-generation cryptocurrency.

Amolecoin was written from scratch and designed over four years to realize the
ideal of Bitcoin and represents the apex of cryptocurrency design.
Amolecoin is not designed to add features to Bitcoin,
but rather improves Bitcoin by increasing simplicity,
security and stripping out everything non-essential.

Some people have hyped the Amolecoin Project as leading into "Bitcoin 3.0".
The coin itself is not "Bitcoin 3.0",
but is rather "Bitcoin 1.0". Bitcoin is a prototype crypto-coin.
Amolecoin was designed to be what Bitcoin would look like if it were built from
scratch, to remedy the rough edges in the Bitcoin design.

- no duplicate coin-base outputs
- enforced checks for hash collisions
- simple deterministic wallets
- no transaction malleability
- no signature malleability
- removal of the scripting language
- CoinJoin and normal transactions are indistinguishable
- elimination of edge-cases that prevent independent node implementations
- <=10 second transaction times
- elimination of the need for mining to achieve blockchain consensus

## Links

* [amolecoin.com](https://www.amolecoin.com)
* [Amolecoin Blog](https://www.amolecoin.com/blog)
* [Amolecoin Docs](https://www.amolecoin.com/docs)
* [Amolecoin Blockchain Explorer](https://explorer.amolecoin.com)
* [Amolecoin Development Telegram Channel](https://t.me/amolecoin)
* [Amolecoin Github Wiki](https://github.com/amolecoin/amolecoin/wiki)

## Table of Contents

<!-- MarkdownTOC levels="1,2,3,4,5" autolink="true" bracket="round" -->

- [Changelog](#changelog)
- [Installation](#installation)
	- [Go 1.14+ Installation and Setup](#go-1.14-installation-and-setup)
	- [Go get amolecoin](#go-get-amolecoin)
	- [Run Amolecoin from the command line](#run-amolecoin-from-the-command-line)
	- [Show Amolecoin node options](#show-amolecoin-node-options)
	- [Run Amolecoin with options](#run-amolecoin-with-options)
	- [Docker image](#docker-image)
	- [Building your own images](#building-your-own-images)
	- [Development image](#development-image)
- [API Documentation](#api-documentation)
	- [REST API](#rest-api)
	- [Amolecoin command line interface](#amolecoin-command-line-interface)
- [Integrating Amolecoin with your application](#integrating-amolecoin-with-your-application)
- [Contributing a node to the network](#contributing-a-node-to-the-network)
- [Creating a new coin](#creating-a-new-coin)
- [Daemon CLI Options](#daemon-cli-options)
- [URI Specification](#uri-specification)
- [Wire protocol user agent](#wire-protocol-user-agent)
- [Offline transaction signing](#offline-transaction-signing)
- [Deploy a public Amolecoin API node with HTTPS](#deploy-a-public-amolecoin-api-node-with-https)
- [Development](#development)
	- [Modules](#modules)
	- [Client libraries](#client-libraries)
	- [Running Tests](#running-tests)
	- [Running Integration Tests](#running-integration-tests)
		- [Stable Integration Tests](#stable-integration-tests)
		- [Live Integration Tests](#live-integration-tests)
		- [Debugging Integration Tests](#debugging-integration-tests)
		- [Update golden files in integration testdata](#update-golden-files-in-integration-testdata)
	- [Test coverage](#test-coverage)
		- [Test coverage for the live node](#test-coverage-for-the-live-node)
	- [Formatting](#formatting)
	- [Code Linting](#code-linting)
	- [Profiling](#profiling)
	- [Fuzzing](#fuzzing)
		- [base58](#base58)
		- [encoder](#encoder)
	- [Dependencies](#dependencies)
		- [Rules](#rules)
		- [Management](#management)
	- [Configuration Modes](#configuration-modes)
		- [Development Desktop Client Mode](#development-desktop-client-mode)
		- [Server Daemon Mode](#server-daemon-mode)
		- [Electron Desktop Client Mode](#electron-desktop-client-mode)
		- [Standalone Desktop Client Mode](#standalone-desktop-client-mode)
	- [Wallet GUI Development](#wallet-gui-development)
		- [Translations](#translations)
	- [Releases](#releases)
		- [Update the version](#update-the-version)
		- [Pre-release testing](#pre-release-testing)
		- [Creating release builds](#creating-release-builds)
		- [Release signing](#release-signing)
- [Responsible Disclosure](#responsible-disclosure)

<!-- /MarkdownTOC -->

## Changelog

[CHANGELOG.md](CHANGELOG.md)

## Installation

Amolecoin supports go1.14+.

### Go 1.14+ Installation and Setup

[Golang 1.14+ Installation/Setup](./INSTALLATION.md)

### Go get amolecoin

```sh
$ go get github.com/amolecoin/amolecoin/cmd/...
```

This will download `github.com/amolecoin/amolecoin` to `$GOPATH/src/github.com/amolecoin/amolecoin`.

You can also clone the repo directly with `git clone https://github.com/amolecoin/amolecoin`,
but it must be cloned to this path: `$GOPATH/src/github.com/amolecoin/amolecoin`.

### Run Amolecoin from the command line

```sh
$ cd $GOPATH/src/github.com/amolecoin/amolecoin
$ make run-client
```

### Show Amolecoin node options

```sh
$ cd $GOPATH/src/github.com/amolecoin/amolecoin
$ make run-help
```

### Run Amolecoin with options

Example:

```sh
$ cd $GOPATH/src/github.com/amolecoin/amolecoin
$ make ARGS="--launch-browser=false -data-dir=/custom/path" run
```

### Docker image

This is the quickest way to start using Amolecoin using Docker.

```sh
$ docker volume create amolecoin-data
$ docker volume create amolecoin-wallet
$ docker run -ti --rm \
    -v amolecoin-data:/data/.amolecoin \
    -v amolecoin-wallet:/wallet \
    -p 9982:9982 \
    -p 9981:9981 \
    amolecoin/amolecoin
```

This image has a `amolecoin` user for the amolecoin daemon to run, with UID and GID 10000.
When you mount the volumes, the container will change their owner, so you
must be aware that if you are mounting an existing host folder any content you
have there will be own by 10000.

The container will run with some default options, but you can change them
by just appending flags at the end of the `docker run` command. The following
example will show you the available options.

```sh
$ docker run --rm amolecoin/amolecoin -help
```

Access the dashboard: [http://localhost:9981](http://localhost:9981).

Access the API: [http://localhost:9981/version](http://localhost:9981/version).

### Building your own images

[Building your own images](docker/images/mainnet/README.md).

### Development image

The [amolecoin/amolecoin-cli docker image](docker/images/dev-cli/README.md) is provided in order to make
easy to start developing Amolecoin. It comes with the compiler, linters, debugger
and the vim editor among other tools.

The [amolecoin/amolecoin-dind docker image](docker/images/dev-docker/README.md) comes with docker installed
and all tools available on `amolecoin/amolecoin-cli:develop` docker image.

Also, the [amolecoin/amolecoin-vscode docker image](docker/images/dev-vscode/README.md) is provided
to facilitate the setup of the development process with [Visual Studio Code](https://code.visualstudio.com)
and useful tools included in `amolecoin/amolecoin-cli`.

## API Documentation

### REST API

[REST API](src/api/README.md).

### Amolecoin command line interface

[CLI command API](cmd/amolecoin-cli/README.md).

## Integrating Amolecoin with your application

[Amolecoin Integration Documentation](INTEGRATION.md)

## Contributing a node to the network

Add your node's `ip:port` to the [peers.txt](peers.txt) file.
This file will be periodically uploaded to https://raw.githubusercontent.com/M3ND3X/amolenodes/master/nodes.txt
and used to seed client with peers.

*Note*: Do not add Amolewire nodes to `peers.txt`.
Only add Amolecoin nodes with high uptime and a static IP address (such as a Amolecoin node hosted on a VPS).

## Creating a new coin

See the [newcoin tool README](./cmd/newcoin/README.md)

## Daemon CLI Options

See the [Amolecoin Daemon CLI options](./cmd/amolecoin/README.md)

## URI Specification

Amolecoin URIs obey the same rules as specified in Bitcoin's [BIP21](https://github.com/bitcoin/bips/blob/master/bip-0021.mediawiki).
They use the same fields, except with the addition of an optional `hours` parameter, specifying the coin hours.

Example Amolecoin URIs:

* `amolecoin:2hYbwYudg34AjkJJCRVRcMeqSWHUixjkfwY`
* `amolecoin:2hYbwYudg34AjkJJCRVRcMeqSWHUixjkfwY?amount=123.456&hours=70`
* `amolecoin:2hYbwYudg34AjkJJCRVRcMeqSWHUixjkfwY?amount=123.456&hours=70&label=friend&message=Birthday%20Gift`

Additonally, if no `amolecoin:` prefix is present when parsing, the string may be treated as an address:

* `2hYbwYudg34AjkJJCRVRcMeqSWHUixjkfwY`

However, do not use this URI in QR codes displayed to the user, because the address can't be disambiguated from other Amolefiber coins.

## Wire protocol user agent

[Wire protocol user agent description](https://github.com/amolecoin/amolecoin/wiki/Wire-protocol-user-agent)

## Offline transaction signing

Before doing the offline transaction signing, we need to have the unsigned transaction created. Using the `amolecoin-cli` tool to create an unsigned transaction in hot wallet, and copy the hex encoded transaction to the computer where the cold wallet is installed. Then use the `amolecoin-cli` tool to sign it offline.

### Create an unsigned transaction

The `amolecoin-cli` tool replys on the APIs of the `amolecoin node`, hence we have to start the node before running the tool.

 Go to the project root and run:

```bash
$ ./run-client.sh -launch-browser=false
```

Once the node is started, we could use the following command to create an unsigned transaction.

```bash
$ amolecoin-cli createRawTransactionV2 $WALLET_FILE $RECIPIENT_ADDRESS $AMOUNT --unsign
```

> Note: Don't forget the `--unsign` flag, otherwise it would try to sign the transaction.

<details>
 <summary>View Output</summary>

```json
b700000000e6b869f570e2bfebff1b4d7e7c9e86885dbc34d6de988da6ff998e7acd7e6e14010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000007531184ad0afeebbff2049b855e0921329cb1cb74d769ac57c057c9c8bd2b6810100000000ed5ea2ca4fe9b4560409b50c5bf7cb39b6c5ff6e50690f00000000000000000000000000
```

</details>

Copy and save the generated transaction string. We will sign it with a cold wallet offline in the next section.

### Sign the transaction

The `amolecoin node` needs to have the most recently `DB` so that the user would not lose much coin hours when signing the transaction. We could copy the full synchronized `data.db` from the hot wallet to the computer where the cold wallet is installed. And place it in `$HOME/.amolecoin/data.db`. Then start the node with the network disabled.

```bash
$ ./run-client.sh -launch-browser=false -disable-networking
```

Run the following command to sign the transaction:

```bash
$ amolecoin-cli signTransaction $RAW_TRANSACTION
```

The `$RAW_TRANSACTION` is the transaction string that we generated in the hot wallet.

If the cold wallet is encrypted, you will be prompted to enter the password to sign the transaction.

<details>
 <summary>View Output</summary>

```json
b700000000e6b869f570e2bfebff1b4d7e7c9e86885dbc34d6de988da6ff998e7acd7e6e14010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000007531184ad0afeebbff2049b855e0921329cb1cb74d769ac57c057c9c8bd2b6810100000000ed5ea2ca4fe9b4560409b50c5bf7cb39b6c5ff6e50690f00000000000000000000000000
```

</details>

Once the transaction is signed, we could copy and save the signed transaction string and broadcast it in the hot wallet.

```bash
$ amolecoin-cli broadcastTransaction $SIGNED_RAW_TRANSACTION
```

A transaction id would be returned and you can check it in the [explorer](https://explorer.amolecoin.com).

## Deploy a public Amolecoin API node with HTTPS

We recommend using [caddy server](https://caddyserver.com/) to deploy a public Amolecoin API node on a
Linux server. The public API node should have the `HTTPS` support, which could be handled automatically
by the `caddy server`. But we need to have a domain and create a DNS record to bind the server ip address
to it.

Suppose we're going to deploy a Amolecoin API node on `apitest.amolecoin.com`, and we have already bound
the server's IP to it. Follow the steps below to complete the deployment.

### Install and run a amolecoin api node

```bash
# Create a amolecoin folder so that the files could be isolated
$ mkdir $HOME/amolecoin && cd $HOME/amolecoin
# Download the amolecoin binary file
$ wget https://downloads.amolecoin.com/wallet/amolecoin-0.26.0-gui-standalone-linux-x64.tar.gz
$ tar -zxvf amolecoin-0.26.0-gui-standalone-linux-x64.tar.gz
$ cd amolecoin-0.26.0-gui-standalone-linux-x64
$ ./amolecoin -web-interface-port=9981 -host-whitelist=$DOMAIN_NAME -enable-api-sets="READ,TXN"
```

> Note: we should running the `amolecoin` node with `-host-whitelist` flag, otherwise it would
> throw `403 Forbidden` error.

### Install the caddy server

```bash
# Create a caddy folder
$ mkdir $HOME/caddy && cd $HOME/caddy
# Download the caddy server binary file
$ wget https://github.com/caddyserver/caddy/releases/download/v1.0.4/caddy_v1.0.4_linux_amd64.tar.gz
$ tar -zxvf caddy_v1.0.4_linux_amd64.tar.gz
$ cd caddy_v1.0.4_linux_amd64
```

The `caddy` tool would be exist in the folder, let's create a `Caddyfile` to define the reverse proxy
rules now.

```bash
cat <<EOF >Caddyfile
apitest.amolecoin.com {
   proxy / localhost:9981 {
      transparent
   }
}
EOF
```

Then run the caddy server

```bash
$ ./caddy
```

You will be prompted to enter an email address to receive the notifications from let's Encrypt.
That's all about the deployment, check the https://apitest.amolecoin.com/api/v1/version to see if
the Amolecoin API node is working correctly.

## Development

We have two branches: `master` and `develop`.

`develop` is the default branch and will have the latest code.

`master` will always be equal to the current stable release on the website, and should correspond with the latest release tag.

### Modules

* `api` - REST API interface
* `cipher` - cryptographic library (key generation, addresses, hashes)
* `cipher/base58` - Base58 encoding
* `cipher/encoder` - reflect-based deterministic runtime binary encoder
* `cipher/encrypt` - at-rest data encryption (chacha20poly1305+scrypt)
* `cipher/go-bip39` - BIP-39 seed generation
* `cli` - CLI library
* `coin` - blockchain data structures (blocks, transactions, unspent outputs)
* `daemon` - top-level application manager, combining all components (networking, database, wallets)
* `daemon/gnet` - networking library
* `daemon/pex` - peer management
* `params` - configurable transaction verification parameters
* `readable` - JSON-encodable representations of internal structures
* `amolecoin` - core application initialization and configuration
* `testutil` - testing utility methods
* `transaction` - methods for creating transactions
* `util` - miscellaneous utilities
* `visor` - top-level blockchain database layer
* `visor/blockdb` - low-level blockchain database layer
* `visor/historydb` - low-level blockchain database layer for historical blockchain metadata
* `wallet` - wallet file management

### Client libraries

Amolecoin implements client libraries which export core functionality for usage from
other programming languages.

* [libamolecoin C client library and SWIG interface](https://github.com/amolecoin/libamolecoin)
* [amolecoin-lite: Javascript and mobile bindings](https://github.com/amolecoin/amolecoin-lite)

### Running Tests

```sh
$ make test
```

### Running Integration Tests

There are integration tests for the CLI and HTTP API interfaces. They have two
run modes, "stable" and "live".

The stable integration tests will use a amolecoin daemon
whose blockchain is synced to a specific point and has networking disabled so that the internal
state does not change.

The live integration tests should be run against a synced or syncing node with networking enabled.

#### Stable Integration Tests

```sh
$ make integration-test-stable
```

or

```sh
$ ./ci-scripts/integration-test-stable.sh -v -w
```

The `-w` option, run wallet integrations tests.

The `-v` option, show verbose logs.

#### Live Integration Tests

The live integration tests run against a live runnning amolecoin node, so before running the test, we
need to start a amolecoin node:

```sh
$ ./run-daemon.sh
```

After the amolecoin node is up, run the following command to start the live tests:

```sh
$ make integration-test-live
```

The above command will run all tests except the wallet-related tests. To run wallet tests, we
need to manually specify a wallet file, and it must have at least `2 coins` and `256 coinhours`,
it also must have been loaded by the node.

We can specify the wallet by setting two environment variables:

* `API_WALLET_ID`, which is the filename (without path), that is loaded by the daemon to test against.
  This is the `"id"` field in API requests. It is used by the API integration tests.
  The wallet directory that the daemon uses can be controlled with the `-wallet-dir` option.
* `CLI_WALLET_FILE`, which is the filename (with path), to be used by the CLI integration tests

If the wallet is encrypted, also set `WALLET_PASSWORD`.

Example of running the daemon with settings for integration tests:

```sh
$ export API_WALLET_ID="$valid_wallet_filename"
$ export CLI_WALLET_FILE="$HOME/.amolecoin/wallets/$valid_wallet_filename"
$ export WALLET_PASSWORD="$wallet_password"
$ make run-integration-test-live
```

Then run the tests with the following command:

```sh
$ make integration-test-live-wallet
```

There are two other live integration test modes for CSRF disabled and networking disabled.

To run the CSRF disabled tests:

```sh
$ export API_WALLET_ID="$valid_wallet_filename"
$ export CLI_WALLET_FILE="$HOME/.amolecoin/wallets/$valid_wallet_filename"
$ export WALLET_PASSWORD="$wallet_password"
$ make run-integration-test-live-disable-csrf
```

```sh
$ make integration-test-live-disable-csrf
```

To run the networking disabled tests, which require a live wallet:

```sh
$ export API_WALLET_ID="$valid_wallet_filename"
$ export CLI_WALLET_FILE="$HOME/.amolecoin/wallets/$valid_wallet_filename"
$ export WALLET_PASSWORD="$wallet_password"
$ make run-integration-test-live-disable-networking
```

Then run the tests with the following command:

```sh
$ make integration-test-live-wallet
```

#### Debugging Integration Tests

Run specific test case:

It's annoying and a waste of time to run all tests to see if the test we real care
is working correctly. There's an option: `-r`, which can be used to run specific test case.
For example: if we only want to test `TestStableAddressBalance` and see the result, we can run:

```sh
$ ./ci-scripts/integration-test-stable.sh -v -r TestStableAddressBalance
```

#### Update golden files in integration testdata

Golden files are expected data responses from the CLI or HTTP API saved to disk.
When the tests are run, their output is compared to the golden files.

To update golden files, use the provided `make` command:

```sh
$ make update-golden-files
```

We can also update a specific test case's golden file with the `-r` option.
For example:
```sh
$ ./ci-scripts/integration-test-stable.sh -v -u -r TestStableAddressBalance
```

### Test coverage

Coverage is automatically generated for `make test` and integration tests run against a stable node.
This includes integration test coverage. The coverage output files are placed in `coverage/`.

To merge coverage from all tests into a single HTML file for viewing:

```sh
$ make check
$ make merge-coverage
```

Then open `coverage/all-coverage.html` in the browser.

#### Test coverage for the live node

Some tests can only be run with a live node, for example wallet spending tests.
To generate coverage for this, build and run the amolecoin node in test mode before running the live integration tests.

In one shell:

```sh
$ make run-integration-test-live-cover
```

In another shell:

```sh
$ make integration-test-live
```

After the tests have run, CTRL-C to exit the process from the first shell.
A coverage file will be generated at `coverage/amolecoin-live.coverage.out`.

Merge the coverage with `make merge-coverage` then open the `coverage/all-coverage.html` file to view it,
or generate the HTML coverage in isolation with `go tool cover -html`

### Formatting

All `.go` source files should be formatted `goimports`.  You can do this with:

```sh
$ make format
```

### Code Linting

Install prerequisites:

```sh
$ make install-linters
```

Run linters:

```sh
$ make lint
```

### Profiling

A full CPU profile of the program from start to finish can be obtained by running the node with the `-profile-cpu` flag.
Once the node terminates, a profile file is written to `-profile-cpu-file` (defaults to `cpu.prof`).
This profile can be analyzed with

```sh
$ go tool pprof cpu.prof
```

The HTTP interface for obtaining more profiling data or obtaining data while running can be enabled with `-http-prof`.
The HTTP profiling interface can be controlled with `-http-prof-host` and listens on `localhost:6060` by default.

See https://golang.org/pkg/net/http/pprof/ for guidance on using the HTTP profiler.

Some useful examples include:

```sh
$ go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
$ go tool pprof http://localhost:6060/debug/pprof/heap
```

A web page interface is provided by http/pprof at http://localhost:6060/debug/pprof/.

### Fuzzing

Fuzz tests are run with [go-fuzz](https://github.com/dvyukov/go-fuzz).
[Follow the instructions on the go-fuzz page](https://github.com/dvyukov/go-fuzz) to install it.

Fuzz tests are written for the following packages:

#### base58

To fuzz the `cipher/base58` package,

```sh
$ make fuzz-base58
```

#### encoder

To fuzz the `cipher/encoder` package,

```sh
$ make fuzz-encoder
```

### Dependencies

#### Rules

Dependencies must not require `cgo`.  This means dependencies cannot be wrappers around C libraries.
Requiring `cgo` breaks cross compilation and interferes with repeatable (deterministic) builds.

Critical cryptographic dependencies used by code in package `cipher` are archived inside the `cipher` folder,
rather than in the `vendor` folder.  This prevents a user of the `cipher` package from accidentally using a
different version of the `cipher` dependencies than were developed, which could have catastrophic but hidden problems.

#### Management

Dependencies are managed with [go modules](https://github.com/golang/go/wiki/Modules).

We still use the `vendor` folder to store our dependencies in case any of the them are 
removed from the internet in the future. 

> When the main module contains a top-level vendor directory and its go.mod file specifies go 1.14 or higher, 
> the go command now defaults to -mod=vendor for operations that accept that flag.


### Configuration Modes
There are 4 configuration modes in which you can run a amolecoin node:
- Development Desktop Daemon
- Server Daemon
- Electron Desktop Client
- Standalone Desktop Client

#### Development Desktop Client Mode
This mode is configured via `run-client.sh`
```bash
$ ./run-client.sh
```

#### Server Daemon Mode
The default settings for a amolecoin node are chosen for `Server Daemon`, which is typically run from source.
This mode is usually preferred to be run with security options, though `-disable-csrf` is normal for server daemon mode, it is left enabled by default.

```bash
$ ./run-daemon.sh
```

To disable CSRF:

```bash
$ ./run-daemon.sh -disable-csrf
```

#### Electron Desktop Client Mode
This mode configures itself via electron-main.js

#### Standalone Desktop Client Mode
This mode is configured by compiling with `STANDALONE_CLIENT` build tag.
The configuration is handled in `cmd/amolecoin/amolecoin.go`

### Wallet GUI Development

The compiled wallet source should be checked in to the repo, so that others do not need to install node to run the software.

Instructions for doing this:

[Wallet GUI Development README](src/gui/static/README.md)

#### Translations

You can find information about how to work with translation files in the [Translations README](./src/gui/static/src/assets/i18n/README.md).

### Releases

#### Update the version

0. If the `master` branch has commits that are not in `develop` (e.g. due to a hotfix applied to `master`), merge `master` into `develop`
0. Make sure the translations are up to date. See the [i18n README](./src/gui/static/src/assets/i18n/README.md) for instructions on how to update translations and how to check if they are up to date.
0. Compile the `src/gui/static/dist/` to make sure that it is up to date (see [Wallet GUI Development README](src/gui/static/README.md))
0. Update version strings to the new version in the following files: `electron/package-lock.json`, `electron/package.json`, `electron/amolecoin/current-amolecoin.json`, `src/cli/cli.go`, `src/gui/static/src/current-amolecoin.json`, `src/cli/integration/testdata/status*.golden`, `template/coin.template`, `README.md` files .
0. If changes require a new database verification on the next upgrade, update `src/amolecoin/amolecoin.go`'s `DBVerifyCheckpointVersion` value
0. Update `CHANGELOG.md`: move the "unreleased" changes to the version and add the date
0. Update the files in https://github.com/amolecoin/repo-info by following the [metadata update procedure](https://github.com/amolecoin/repo-info/#updating-amolecoin-repository-metadate),
0. Merge these changes to `develop`
0. Follow the steps in [pre-release testing](#pre-release-testing)
0. Make a PR merging `develop` into `master`
0. Review the PR and merge it
0. Tag the `master` branch with the version number. Version tags start with `v`, e.g. `v0.20.0`.
    Sign the tag. If you have your GPG key in github, creating a release on the Github website will automatically tag the release.
    It can be tagged from the command line with `git tag -as v0.20.0 $COMMIT_ID`, but Github will not recognize it as a "release".
0. Make sure that the client runs properly from the `master` branch
0. Release builds are created and uploaded by travis. To do it manually, checkout the `master` branch and follow the [create release builds](electron/README.md) instructions.

If there are problems discovered after merging to `master`, start over, and increment the 3rd version number.
For example, `v0.20.0` becomes `v0.20.1`, for minor fixes.

#### Pre-release testing

Performs these actions before releasing:

* `make check`
* `make integration-test-live`
* `make integration-test-live-disable-networking` (requires node run with `-disable-networking`)
* `make integration-test-live-disable-csrf` (requires node run with `-disable-csrf`)
* `make intergration-test-live-wallet` (see [live integration tests](#live-integration-tests)) 6 times: with an unencrypted and encrypted wallet for each wallet type: `deterministic`, `bip44` and `collection`
* `go run cmd/amolecoin-cli/amolecoin-cli.go checkdb` against a fully synced database
* `go run cmd/amolecoin-cli/amolecoin-cli.go checkDBDecoding` against a fully synced database
* On all OSes, make sure that the client runs properly from the command line (`./run-client.sh` and `./run-daemon.sh`)
* Build the releases and make sure that the Electron client runs properly on Windows, Linux and macOS.
    * Use a clean data directory with no wallets or database to sync from scratch and verify the wallet setup wizard.
    * Load a test wallet with nonzero balance from seed to confirm wallet loading works
    * Send coins to another wallet to confirm spending works
    * Restart the client, confirm that it reloads properly
* For both the Android and iOS mobile wallets, configure the node url to be https://staging.node.amolecoin.com
  and test all operations to ensure it will work with the new node version.

#### Creating release builds

[Create Release builds](electron/README.md).

#### Release signing

Releases are signed with this PGP key:

`0x913BBD5206B19620`

The fingerprint for this key is:

```
pub   ed25519 2019-09-17 [SC] [expires: 2023-09-16]
      98F934F04F9334B81DFA3398913BBD5206B19620
uid           [ultimate] iketheadore amolecoin <luxairlake@protonmail.com>
sub   cv25519 2019-09-17 [E] [expires: 2023-09-16]
```

Keybase.io account: https://keybase.io/iketheadore

Follow the [Tor Project's instructions for verifying signatures](https://www.torproject.org/docs/verifying-signatures.html.en).

If you can't or don't want to import the keys from a keyserver, the signing key is available in the repo: [iketheadore.asc](iketheadore.asc).

Releases and their signatures can be found on the [releases page](https://github.com/amolecoin/amolecoin/releases).

Instructions for generating a PGP key, publishing it, signing the tags and binaries:
https://gist.github.com/iketheadore/6485585ce2d22231c2cb3cbc77e1d7b7

## Responsible Disclosure

Security flaws in amolecoin source or infrastructure can be sent to security@amolecoin.com.
Bounties are available for accepted critical bug reports.

PGP Key for signing:

```
-----BEGIN PGP PUBLIC KEY BLOCK-----

mDMEXYCYPxYJKwYBBAHaRw8BAQdAeDPi3n9xLv5xGsxbcbwZjfV4h772W+GPZ3Mz
RS17STm0L2lrZXRoZWFkb3JlIHNreWNvaW4gPGx1eGFpcmxha2VAcHJvdG9ubWFp
bC5jb20+iJYEExYIAD4WIQSY+TTwT5M0uB36M5iRO71SBrGWIAUCXYCYPwIbAwUJ
B4TOAAULCQgHAgYVCgkICwIEFgIDAQIeAQIXgAAKCRCRO71SBrGWID0NAP0VRiNA
2Kq2uakPMqV29HY39DVhc9QgxJfMIwXWtFxKAwEAn0NqGRV/iKXNf+qxqAtMWa5X
F2S36hkEfDHO5W44DwC4OARdgJg/EgorBgEEAZdVAQUBAQdAeiEz/tUmCgOA67Rq
ANmHmX2vrdZp/SfJ9KOI2ANCCm8DAQgHiH4EGBYIACYWIQSY+TTwT5M0uB36M5iR
O71SBrGWIAUCXYCYPwIbDAUJB4TOAAAKCRCRO71SBrGWIJOJAQDTaqxpcLtAw5kH
Hp2jWvUnLudIONeqeUTCmkLJhcNv1wD+PFJZWMKD1btIG4pkXRW9YoA7M7t5by5O
x5I+LywZNww=
=p6Gq
-----END PGP PUBLIC KEY BLOCK-----
```


Key ID: [0x913BBD5206B19620](https://pgp.mit.edu/pks/lookup?search=0x913BBD5206B19620&op=index)

The fingerprint for this key is:

```
pub   ed25519 2019-09-17 [SC] [expires: 2023-09-16]
      98F934F04F9334B81DFA3398913BBD5206B19620
uid           [ultimate] iketheadore amolecoin <luxairlake@protonmail.com>
sub   cv25519 2019-09-17 [E] [expires: 2023-09-16]
```

Keybase.io account: https://keybase.io/iketheadore
