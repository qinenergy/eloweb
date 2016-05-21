package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EloTestSuite struct {
	suite.Suite
}

func TestElo(t *testing.T) {
	eloSuite := &EloTestSuite{}

	suite.Run(t, eloSuite)
}

func (s *EloTestSuite) TestExampleMatch_OnePlayerGainsRating_OtherLoses() {

	player1Rank, player2Rank := Win(2400, 2000)

	require.Condition(s.T(), equalWithTolerance(2403, player1Rank, 0.1), "Player 1 rank not correct was %v, expected %v", player1Rank, 2403)
	require.Condition(s.T(), equalWithTolerance(1997, player2Rank, 0.1), "Player 2 rank not correct was %v, expected %v", player2Rank, 1997)
}

func equalWithTolerance(expected, actual, tolerance float64) assert.Comparison {
	return func() bool { return expected-tolerance < actual && actual < expected+tolerance }
}
