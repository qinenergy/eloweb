package lib

import (
	"math"
)

const (
	//K defines the sensitivity of rankings to results
	k = 32
)

func Win(r1, r2 float64) (float64, float64) {
	//Calculate transformed ratings
	r1Transformed := math.Pow(10, r1/400)
	r2Transformed := math.Pow(10, r2/400)

	//Calculate expected scores
	r1Expected := r1Transformed / (r1Transformed + r2Transformed)
	r2Expected := r2Transformed / (r1Transformed + r2Transformed)

	//Becuase player 1 won
	var s1 float64 = 1
	var s2 float64 = 0

	//Update score with results
	newR1 := r1 + k*(s1-r1Expected)
	newR2 := r2 + k*(s2-r2Expected)

	return newR1, newR2
}
