package elo

const (
	startingScore = 1000
)

type User struct {
	Id    string
	Score float64
}

type Result struct {
	Winner string
	Loser  string
}

func CalculateRankings(results []*Result) map[string]*User {
	userById := map[string]*User{}
	for _, result := range results {

		winner := getOrCreateUser(userById, result.Winner)
		loser := getOrCreateUser(userById, result.Loser)

		newWinnerScore, newLoserScore := updateRanking(winner.Score, loser.Score)

		winner.Score = newWinnerScore
		loser.Score = newLoserScore
	}
	return userById
}

func getOrCreateUser(currentUsers map[string]*User, userId string) *User {
	user, ok := currentUsers[userId]
	if ok {
		return user
	}
	newUser := &User{
		Id:    userId,
		Score: startingScore,
	}
	currentUsers[userId] = newUser
	return newUser
}
