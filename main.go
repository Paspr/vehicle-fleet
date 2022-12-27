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

type Brand struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	SeatCapacity   string  `json:"seat_capacity"`
	TankCapacity   float32 `json:"tank_capacity"`
	WeightCapacity float32 `json:"weight_capacity"`
}

type VehicleBrand struct {
	Vehicle Vehicle `json:"vehicle_entities"`
	Brand   Brand   `json:"brand_entities"`
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

func ListBrandHandler(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM brand")
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error with DB"})
	}

	items := make([]Brand, 0)

	if rows != nil {
		defer rows.Close()
		for rows.Next() {

			item := Brand{}
			if err := rows.Scan(&item.ID, &item.Name, &item.SeatCapacity,
				&item.TankCapacity, &item.Type, &item.WeightCapacity); err != nil {
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
	router.GET("/brands", ListBrandHandler)

	router.Run()
}
