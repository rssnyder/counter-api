package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

const (
	getCounter = `SELECT value FROM counters WHERE uuser=$1 AND counter=$2 LIMIT 1;`
	incCounter = `UPDATE counters SET value = value + 1 WHERE uuser=$1 AND counter=$2 RETURNING value;`
	decCounter = `UPDATE counters SET value = value - 1 WHERE uuser=$1 AND counter=$2 RETURNING value;`
	iniCounter = `INSERT INTO counters(uuser, counter, value) values($1,$2,$3) RETURNING value;`
)

func main() {

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", getEnv("PG_HOST", "postgres"), getEnv("PG_USER", "postgres"), getEnv("PG_PASSWORD", "postgres"), getEnv("PG_DB", "postgres"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	p := ginprometheus.NewPrometheus("gin")
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		url := c.Request.URL.Path
		for _, p := range c.Params {
			if p.Key == "user" {
				url = strings.Replace(url, p.Value, ":user", 1)
			}
			if p.Key == "counter" {
				url = strings.Replace(url, p.Value, ":counter", 1)
			}
		}
		return url
	}
	p.Use(r)

	r.POST("/:user/:counter", func(c *gin.Context) {
		var value string

		user := c.Param("user")
		counter := c.Param("counter")

		err = db.QueryRow(iniCounter, user, counter, 0).Scan(&value)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(500)
		} else {
			c.String(http.StatusOK, "%s\n", value)
		}
	})

	r.GET("/:user/:counter", func(c *gin.Context) {
		var value string

		user := c.Param("user")
		counter := c.Param("counter")

		row := db.QueryRow(getCounter, user, counter)

		switch err := row.Scan(&value); err {
		case sql.ErrNoRows:
			log.Printf("nil user/counter requested: %s/%s\n", user, counter)
			c.AbortWithStatus(204)
		case nil:
			c.String(http.StatusOK, "%s\n", value)
		default:
			log.Println(err)
		}
	})

	r.HEAD("/:user/:counter", func(c *gin.Context) {
		var value string

		user := c.Param("user")
		counter := c.Param("counter")

		err = db.QueryRow(incCounter, user, counter).Scan(&value)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(500)
		} else {
			c.String(http.StatusOK, "")
		}
	})

	r.DELETE("/:user/:counter", func(c *gin.Context) {
		var value string

		user := c.Param("user")
		counter := c.Param("counter")

		err = db.QueryRow(decCounter, user, counter).Scan(&value)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(500)
		} else {
			c.String(http.StatusOK, "")
		}
	})

	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
