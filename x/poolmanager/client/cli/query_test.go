package cli_test

import (
	gocontext "context"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/dymensionxyz/dymension/testutil"
	poolmanagerqueryproto "github.com/dymensionxyz/dymension/x/poolmanager/client/queryproto"
)

type QueryTestSuite struct {
	testutil.KeeperTestHelper
	queryClient poolmanagerqueryproto.QueryClient
}

func (s *QueryTestSuite) SetupSuite() {
	s.Setup()
	s.queryClient = poolmanagerqueryproto.NewQueryClient(s.QueryHelper)
	// create a new pool
	s.PrepareBalancerPool()
	s.Commit()
}

func (s *QueryTestSuite) TestQueriesNeverAlterState() {
	testCases := []struct {
		name   string
		query  string
		input  interface{}
		output interface{}
	}{
		{
			"Query num pools",
			"/osmosis.poolmanager.v1beta1.Query/NumPools",
			&poolmanagerqueryproto.NumPoolsRequest{},
			&poolmanagerqueryproto.NumPoolsResponse{},
		},
		{
			"Query estimate swap in",
			"/osmosis.poolmanager.v1beta1.Query/EstimateSwapExactAmountIn",
			&poolmanagerqueryproto.EstimateSwapExactAmountInRequest{},
			&poolmanagerqueryproto.EstimateSwapExactAmountInResponse{},
		},
		{
			"Query estimate swap out",
			"/osmosis.poolmanager.v1beta1.Query/EstimateSwapExactAmountOut",
			&poolmanagerqueryproto.EstimateSwapExactAmountOutRequest{},
			&poolmanagerqueryproto.EstimateSwapExactAmountOutResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			s.SetupSuite()
			err := s.QueryHelper.Invoke(gocontext.Background(), tc.query, tc.input, tc.output)
			s.Require().NoError(err)
			s.StateNotAltered()
		})
	}
}

func TestQueryTestSuite(t *testing.T) {

	// TODO: re-enable this once poolmanager is fully merged.
	t.SkipNow()

	suite.Run(t, new(QueryTestSuite))
}
