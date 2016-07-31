package main

import (
	"log"
	"net/http"
	"os"

	"database/sql"
	"errors"
	"fmt"
	"github.com/RaganH/eloweb/lib/elo"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

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

	router.GET("/result", addScore)
	router.GET("/rankings", rankings)

	router.Run(":" + port)
}

func rankings(c *gin.Context) {
	db, err := sql.Open("postgres", "user=postgres dbname=postgres")
	if err != nil {
		log.Printf("Error querying db for results %v\n", err)
		c.Abort()
	}

	rows, err := db.Query("SELECT * FROM Results")
	if err != nil {
		log.Printf("Error querying db for results %v\n", err)
		c.Abort()
	}
	defer db.Close()
	var results []*elo.Result
	for rows.Next() {
		r := &elo.Result{}
		err = rows.Scan(&r.Winner, &r.Loser)
		if err != nil {
			log.Println("Error scanning results rows: %v\n", err)
			c.Abort()
		}
		results = append(results, r)
	}

	fmt.Println(err)
	c.HTML(http.StatusOK, "rankings.tmpl.html", elo.CalculateRankings(results))
}

func addScore(c *gin.Context) {
	winner := c.Query("winner")
	loser := c.Query("loser")

	if winner == "" || loser == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Both usernames not provided"))
		return
	}

	db, err := sql.Open("postgres", "user=postgres dbname=postgres")
	if err != nil {
		log.Printf("Error querying db for results %v\n", err)
		c.Abort()
	}
	defer db.Close()

	_, err = db.Query("INSERT INTO Results VALUES ($1,$2)", winner, loser)
	if err != nil {
		log.Printf("Error querying db for results %v\n", err)
		c.Abort()
	}

	c.Redirect(http.StatusTemporaryRedirect, "/rankings")
}
