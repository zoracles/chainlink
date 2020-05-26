# Zoracles

Join our growing [community](https://discord.gg/9vn7UdY)! 

Zoracles is a decentralized network of oracles that leverage zero-knowledge proofs to transfer data securely.

The current node supports:

- easy connectivity of on-chain contracts to any off-chain computation or API
- multiple methods for scheduling both on-chain and off-chain computation for a user's smart contract
- automatic gas price bumping to prevent stuck transactions, assuring your data is delivered in a timely manner
- push notification of smart contract state changes to off-chain systems, by tracking Ethereum logs
- translation of various off-chain data types into EVM consumable types and transactions
- easy to implement smart contract libraries for connecting smart contracts directly to their preferred oracles
- easy to install node, which runs natively across operating systems, blazingly fast, and with a low memory footprint

## Install

1. [Install Go 1.14](https://golang.org/doc/install#install), and add your GOPATH's [bin directory to your PATH](https://golang.org/doc/code.html#GOPATH)
2. Install [NodeJS](https://nodejs.org/en/download/package-manager/) & [Yarn](https://yarnpkg.com/lang/en/docs/install/)
3. Install [Postgres (>= 9.6)](https://wiki.postgresql.org/wiki/Detailed_installation_guides).
4. Download Zoracles: `git clone https://github.com/zoracles/core && cd core`
5. Build and install Zoracles: `make install`
6. Run the node: `core help`

### Ethereum Node Requirements

In order to run the Zoracles node you must have access to a running Ethereum node with an open websocket connection.
Any Ethereum based network will work once you've [configured](https://github.com/smartcontractkit/chainlink#configure) the chain ID.
Ethereum node versions currently tested and supported:

- [Parity 1.11+](https://github.com/paritytech/parity-ethereum/releases) (due to a [fix with pubsub](https://github.com/paritytech/parity/issues/6590).)
- [Geth 1.8+](https://github.com/ethereum/go-ethereum/releases)

## Run

To start your Zoracles node, simply run:

```bash
core node start
```

By default this will start on port 6688, where it exposes a [REST API](https://github.com/smartcontractkit/chainlink/wiki/REST-API).

Once your node has started, you can view your current jobs with:

```bash
core jobs list
```

View details of a specific job with:

```bash
core jobs show "$JOB_ID"
```


## Configure

You can configure your node's behavior by setting environment variables which can be, along with default values that get used if no corresponding environment variable is found. 


## Contributing

Zoracle's source code is [licensed under the MIT License](./LICENSE), and contributions are welcome.

Please check out our [contributing guidelines](./docs/CONTRIBUTING.md) for more details.

Thank you!
