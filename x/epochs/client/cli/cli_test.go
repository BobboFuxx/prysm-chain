package cli_test

import (
	"testing"

	"github.com/dymensionxyz/dymension/osmoutils/osmocli"
	"github.com/dymensionxyz/dymension/x/epochs/client/cli"
	"github.com/dymensionxyz/dymension/x/epochs/types"
)

func TestGetCmdCurrentEpoch(t *testing.T) {
	desc, _ := cli.GetCmdCurrentEpoch()
	tcs := map[string]osmocli.QueryCliTestCase[*types.QueryCurrentEpochRequest]{
		"basic test": {
			Cmd: "day",
			ExpectedQuery: &types.QueryCurrentEpochRequest{
				Identifier: "day",
			},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}

func TestGetCmdEpochsInfo(t *testing.T) {
	desc, _ := cli.GetCmdEpochInfos()
	tcs := map[string]osmocli.QueryCliTestCase[*types.QueryEpochsInfoRequest]{
		"basic test": {
			Cmd:           "",
			ExpectedQuery: &types.QueryEpochsInfoRequest{},
		},
	}
	osmocli.RunQueryTestCases(t, desc, tcs)
}
