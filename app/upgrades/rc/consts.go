package rc

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	UpgradeName = "rc-TBD"
)

var (
	PYMTokenMetata = banktypes.Metadata{
		Description: "Denom metadata for PYM (upym)",
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    "upym",
				Exponent: 0,
				Aliases:  []string{},
			},
			{
				Denom:    "PYM",
				Exponent: 6,
				Aliases:  []string{},
			},
		},
		Base:    "upym",
		Display: "PYM",
		Name:    "PYM",
		Symbol:  "PYM",
		URI:     "",
		URIHash: "",
	}
)
