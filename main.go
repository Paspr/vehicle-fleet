package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Vehicle struct {
	ID      int     `json:"id"`
	Cost    float32 `json:"cost"`
	YOM     int     `json:"yom"`
	Mileage float32 `json:"mileage"`
}

var db *sql.DB
var err error

func ListVehicleHandler(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM vehicle")
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
	}

	items := make([]Vehicle, 0)

	if rows != nil {
		defer rows.Close()
		for rows.Next() {

			item := Vehicle{}
			if err := rows.Scan(&item.ID, &item.Cost, &item.YOM, &item.Mileage); err != nil {
				fmt.Println(err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
			}
			items = append(items, item)
		}
	}

	c.JSON(http.StatusOK, items)
}

func main() {

	db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/vehicle?sslmode=disable")

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	router := gin.Default()

	router.StaticFile("/", "static/index.html")
	router.GET("/vehicles", ListVehicleHandler)

	router.Run()
}
