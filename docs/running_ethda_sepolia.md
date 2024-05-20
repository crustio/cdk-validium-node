## How to run cdk node with Ethda on sepolia testnet

1. Clone this GitHub repostory to your local machine:

```bash
git clone https://github.com/crustio/zkevm-contracts.git
```

2. Navigate to the cloned directory and checkout to ethda branch:

```bash
cd zkevm-contracts && git checkout feat/support-ethda-da
```

3. Prepare two Ethereum private key accounts (with sufficient sepolia ETH):

	- \<aggregator-account\>:   aggregator.keystore
	- \<sequencer-account\>:    sequencer.keystore

You can prepare your account with [evmkey-tool](https://github.com/crustio/evmkey):
- `go install -ldflags="-w -s" github.com/crustio/evmkey@v1.0.2`
- `evmkey account new`
- ATTENTION. evmkey is now a simple key util, Please backup your mnemonic phrase yourself


4. Prepare infura API key:
	- visit [infra/register](https://app.infura.io/register) and regist
	- visit [app.infura.io](https://app.infura.io/)，and copy an API Key

5. run command `cp .env.example .env` and rewrite environment `MNEMONIC` and `INFURA_PROJECT_ID`

- `MNEMONIC`: // Mnemonic phrase corresponding to \<sequencer-account\>

- `INFURA_PROJECT_ID`: // your \<infura API key\>

6. update deploy parameters

```bash
cp deployment/v2/deploy_parameters.json.example deployment/v2/deploy_parameters.json

cp deployment/v2/create_rollup_parameters.json.example deployment/v2/create_rollup_parameters.json
```

Edit deploy_parameters.json:

```json
{
  "test": true,
  "timelockAdminAddress": "<sequencer-account-address>", // address of <sequencer-account>
  "minDelayTimelock": 1, // Modify to 1
  "salt": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "initialZkEVMDeployerOwner": "<sequencer-account-address>", // address of <sequencer-account>
  "admin": "<sequencer-account-address>", // address of <sequencer-account>
  "trustedAggregator": "<aggregator-account-address>", // address of <aggregator-account>
  "trustedAggregatorTimeout": 604799,
  "pendingStateTimeout": 604799,
  "emergencyCouncilAddress": "<sequencer-account-address>", // address of <sequencer-account>
  "polTokenAddress": "",  // Autofill, leave blank
  "zkEVMDeployerAddress": "", // Autofill, leave blank
  "deployerPvtKey": "",
  "maxFeePerGas": "",
  "maxPriorityFeePerGas": "",
  "multiplierGas": ""
}
```

Edit create_rollup_parameters:

```json
{
  "realVerifier": false,
  "trustedSequencerURL": "http://localhost:8123", // your cdk domain
  "networkName": "zkevm",
  "description": "0.0.1",
  "trustedSequencer": "<sequencer-account-address>", // address of <sequencer-account>
  "chainID": 1001,
  "adminZkEVM": "<sequencer-account-address>", // address of <sequencer-account>
  "forkID": 8,
  "consensusContract": "PolygonValidiumEtrog", // use PolygonValidiumEtrog
  "dataAvailabilityProtocol": "Ethda", // use Ethda
  "gasTokenAddress": "",
  "deployerPvtKey": "",
  "maxFeePerGas": "",
  "maxPriorityFeePerGas": "",
  "multiplierGas": ""
}
```

7. Create rollup:

```bash
npm install && npm run deploy:testnet:v2:sepolia
```

8. Check output:

```bash
tree deployments/

deployments
└── sepolia_1715752769
    ├── create_rollup_output.json
    ├── deploy_output.json
    ├── deploy_parameters.json
    ├── genesis.json
    └── sepolia.json
```

9. Build the genesis file for the cdk node:
* First, clone the [https://github.com/crustio/cdk-validium-node](https://github.com/crustio/cdk-validium-node) repo and checkout `ethda`.
* Edit the `test/config/test.genesis.config.json` file taking values in the generated output files created previously in the contract repo’s deployments/sepolia_1715752769 folder:

Parameters to change
- l1Config.polygonZkEVMAddress ==> rollupAddress @ `create_rollup_output.json`
- l1Config.polygonRollupManagerAddress ==> polygonRollupManagerAddress @ `deploy_output.json` 
- l1Config.polTokenAddress ==> polTokenAddress @ `deploy_output.json` 
- l1Config.polygonZkEVMGlobalExitRootAddress ==> polygonZkEVMGlobalExitRootAddress @ `deploy_output.json` 
- rollupCreationBlockNumber ==> createRollupBlockNumber @ `create_rollup_output.json `
- rollupManagerCreationBlockNumber ==> deploymentRollupManagerBlockNumber @ `deploy_output.json `
- root ==> root @ `genesis.json `
- genesis ==> genesis @ `genesis.json`

10. Build config file for the cdk node:
* First, replace `test/sequencer.keystore` and `test/aggregator.keystore` with your \<sequencer-account\> and \<aggregator-account\>
* Edit `test/config/test.node.config.toml` configuration of the `SequenceSender`, `Aggregator`, `EthTxManager` and `Etherman` parts in the file:

```bash
[SequenceSender]
…
L2Coinbase = "<sequencer-account-address>", // address of <sequencer-account>
PrivateKey = {Path = "/pk/sequencer.keystore", Password = "<sequencer-account-password>"} // password of <sequencer-account>
	[SequenceSender.StreamClient]
		Server = "zkevm-sequencer:6900" 
…

[Aggregator]
…
SenderAddress = "<aggregator-account-address>" // address of <aggregator-account>

[EthTxManager]
ForcedGas = 0
PrivateKeys = [
	{Path = "/pk/sequencer.keystore", Password = "<sequencer-account-password>"}, //  password of <sequencer-account>
	{Path = "/pk/aggregator.keystore", Password = "<aggregator-account-password>"} // password of <aggregator-account>
]

[Etherman]
URL = "<L1-RPC>" // L1 rpc: sepolia rpc
…

```

11.run cdk node:

```bash
cd test/ && make run-ethda-sepolia
```
