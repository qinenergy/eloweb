package main

import (
	"log"
	"net/http"
	"os"

	"errors"
	"github.com/RaganH/eloweb/lib/elo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var allResults []*elo.Result

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/result", func(c *gin.Context) { addScore(c) })
	router.GET("/rankings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rankings.tmpl.html", elo.CalculateRankings(allResults))
	})

	router.Run(":" + port)
}

func addScore(c *gin.Context) {
	winner := c.Query("winner")
	loser := c.Query("loser")

	if winner == "" || loser == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Both usernames not provided"))
		return
	}

	allResults = append(allResults, &elo.Result{
		Winner: winner,
		Loser:  loser,
	})

	c.Redirect(http.StatusTemporaryRedirect, "/rankings")

}
