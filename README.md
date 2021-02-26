# cosmos-faucet

This faucet was developed to use [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) executable binaries only. Hence
it is compatible with pratically any blockchain built with [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) even if
different types of keys are used (as [ethermint](https://github.com/cosmos/ethermint) for example).

The `main` version supports `0.40 provenance` only. 

## Installation

You can build the faucet with:

```bash
$ make build
```

The executable binary will be avaialable in the `./build/` directory. To install it to `$GOPATH/bin`, use:

```bash
$ make install
```

## Usage

### Configuration

You can configure the faucet either using command line flags or environment variables. The following table
shows the available configuration options and respective defaults:

| flag             	| env              	| description                                                      	| default                      	|
|------------------	|------------------	|------------------------------------------------------------------	|------------------------------	|
| log-level        	| LOG_LEVEL        	| the log level to be used (trace, debug, info, warn or error)     	| info                         	|
| port             	| PORT             	| the port in which the server will be listening for HTTP requests 	| 8000                         	|
| key-name         	| KEY_NAME         	| the name of the key to be used by the faucet                     	| faucet                       	|
| mnemonic         	| MNEMONIC         	| a mnemonic to restore an existing key (this is optional)         	|                              	|
| keyring-password 	| KEYRING_PASSWORD 	| the password for accessing the keys keyring                      	|                              	|
| cli-name         	| CLI_NAME         	| the name of the cli executable                                   	| provenanced                  	|
| denom            	| DENOM            	| the denomination of the coin to be distributed by the faucet     	| nhash                        	|
| credit-amount    	| CREDIT_AMOUNT    	| the amount to credit in each request                             	| 10                         	|
| max-credit       	| MAX_CREDIT       	| the maximum credit per account                                   	| 100000000                    	|

### Example

Start the faucet with:
 with environment variables:

```bash
$ export KEYRING_PASSWORD=<password_here>
$ export NODE_ADDR=35.243.142.236
$ export MNEMONIC=<your_mnemonic>
$ faucet
INFO[0000] listening on :42000
```

### Request tokens

You can request tokens by sending a `POST` request to any path on the server, with a key address in a `JSON`
each request disburses `10nhash`, any account can only get max of `10000nhash`:

```bash
$ curl -X POST 'localhost:42000' -d '{"address": "tp1604k546fa8frjs8r7xpgjhjgl2r9lkcjg2eck3"}'
{"status": "ok"}
```


to query the account you gave 10nhash to 

```bash
provenanced -t q bank balances   tp1604k546fa8frjs8r7xpgjhjgl2r9lkcjg2eck3 --node=tcp://35.243.142.236:26657
balances:
- amount: "99949957170"
  denom: nhash
pagination:
  next_key: null
  total: "0"
```
