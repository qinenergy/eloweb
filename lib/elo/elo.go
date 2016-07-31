package elo

import (
	"math"
)

const (
	//K defines the sensitivity of rankings to results
	k = 32
)

func updateRanking(winnerRank, loserRank float64) (newWinnerRank float64, newLoserRank float64) {
	//Calculate transformed ratings
	winnerRankTransformed := math.Pow(10, winnerRank/400)
	loserRankTransformed := math.Pow(10, loserRank/400)

	//Calculate expected scores
	r1Expected := winnerRankTransformed / (winnerRankTransformed + loserRankTransformed)
	r2Expected := loserRankTransformed / (winnerRankTransformed + loserRankTransformed)

	//Becuase player 1 won
	var s1 float64 = 1
	var s2 float64 = 0

	//Update score with results
	newR1 := winnerRank + k*(s1-r1Expected)
	newR2 := loserRank + k*(s2-r2Expected)

	return newR1, newR2
}
