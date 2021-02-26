module github.com/FigureTechnologies/provenance-faucet

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.40.1
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.6.0
	github.com/tendermint/tendermint v0.34.3
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
