# GoChain

[![Lint](https://github.com/BrianLusina/gochain/actions/workflows/lint.yml/badge.svg)](https://github.com/BrianLusina/gochain/actions/workflows/lint.yml)
[![Release](https://github.com/BrianLusina/gochain/actions/workflows/release.yml/badge.svg)](https://github.com/BrianLusina/gochain/actions/workflows/release.yml)
[![Code Scanning](https://github.com/BrianLusina/gochain/actions/workflows/codeql.yml/badge.svg)](https://github.com/BrianLusina/gochain/actions/workflows/codeql.yml)

This is a simple demonstration of a blockchain written in Go. Note that it does not actually integrate into any blockchain networks nor does it handle crypto mining. This shows the simple basic building blocks that are available in a blockchain network and how they interact with each other.

A blockchain is a distributed ledger that consists of a chain of blocks linked together. They provide benefits such as decentralization, transparency, and enhanced security through cryptographic algorithms.

A block is a fundamental blockchain unit that contains a single or a set of transactions or digital records. Each block contains a unique hash generated from its data and the previous block’s hash, forming an immutable chain.

## Requirements

This uses Go 1.22 which can be downloaded from [here](https://go.dev/doc/install), that should/is the only requirement for running this project :).

## Running & testing the project

Running tests can be done using [make](https://www.gnu.org/software/make/):

``` shell
make test
```

Tests can be run with coverage using:

``` shell

make test-coverage
```

Running the project can be done with:

```shell
go run cmd/main.go
```

Which should print out something like:

```plain
Alice's wallet created successfully
Bob's wallet created successfully
Alice to Bob Transaction created successfully
Transaction Verified Successfully
Previous hash: 
Data in Block: Genesis
Hash of block: 64572909aa4583d58dfe87f6b89d667e
IsValidPoW: true

Transactions:
Sender: Coinbase
Receiver: Genesis
Amount: 0.000000
Coinbase: true


Previous hash: 
Data in Block: Block 1
Hash of block: aa4cd1f3eb42649d21bb771c1642a7c5
IsValidPoW: true

Transactions:
Sender: Coinbase
Receiver: Alice
Amount: 10.000000
Coinbase: true

Sender: 27678327667185575677489037260907325674021522134083400034177112902348904042313936369923198279006157504855144283793124760940507422911068399508026627819789954832926149959578991099596026030841131337060064941323257883256686810980076767922128852540542630614743735618116777447469897490400728784727793557928294883630025457291130735842642864398952225098217652381893878421606663961931399513935938770701861422418735155920408098488953629592530576167228293886834370834211520409149571346887772040541172191530342965306896088647320150528765228460032300528221206644960353747335378226302103855814689458258703419858479160676626978834769
Receiver: 22321842186648171711481631361825596969077862815338303655230759636304662999247953642304855735801785457933690007062049794847391912065859220962938889150125127119169492510520150448602744871584407465507053644035162425397503171886919057557268641938939603493404216901046546691356038307302150602819433320420209848660519977889349476102932781954815112909529588354194567031003032737406368130261000058812981139197361508303509783268930472805546962990664244280117025926374918527854502727305894021847305811040250934579647257704417574979921228231069251138366608448318330171156344565564569332926956614575594824125079934139555665679381
Amount: 5.000000
Coinbase: false


Previous hash: 
Data in Block: Block 2
Hash of block: a723832e1a0bc2016ca832e37677b0e8
IsValidPoW: true

Transactions:
Sender: Coinbase
Receiver: Bob
Amount: 10.000000
Coinbase: true

Sender: Bob
Receiver: Charlie
Amount: 2.300000
Coinbase: false
```

## Built With

- [Go](https://go.dev/doc/)
- [Testify](https://github.com/stretchr/testify)

## Contributing

Please read the [contributing guide](./github/CONTRIBUTING.md) to learn how to contribute to this project.

## Versioning

[Semantic versioning](https://semver/) is used to track the version of the project

---

[![forthebadge](https://forthebadge.com/images/featured/featured-built-with-love.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/license-mit.svg)](https://forthebadge.com)
