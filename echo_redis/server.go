package main

import (
	"encoding/json"
	"fmt"
	// "github.com/go-redis/redis"
	"github.com/labstack/echo"
	// "log"
	// "github.com/fatih/structs"
	"net/http"
	// "reflect"
	// "database/sql"
	"github.com/go-redis/redis/v7"
)

type RedisData struct {
}

type City struct {
	Name    string `json:"name"`
	Authors []struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Studies []struct {
			Languages []string `json:"languages"`
		} `json:"studies"`
		OnDuty bool
	} `json:"authors"`
}

type Staging struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

var Data = `{"name": "hcm", "authors": [
    {"name": "duc_tri", "age": 18, "studies": [{"languages": ["eng", "zhe"]}]},
    {"name": "theanh", "age": 20, "studies": [{"languages": ["eng", "fre"]}]}
]}`

func show(c echo.Context) error {
	tp := c.FormValue("city")
	var city City
	json.Unmarshal([]byte(Data), &city)
	pCity := &city
	pCity.Name = tp
	for i := 0; i < len(pCity.Authors); i++ {
		if pCity.Authors[i].Age == 18 {
			pCity.Authors[i].OnDuty = true
		}
	}
	var jsonData []byte
	jsonData, _ = json.Marshal(pCity)
	// err := client.Set("key", jsonData, 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	return c.String(http.StatusOK, string(jsonData))
}

func showYml(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		getData, err := client.Get("students").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", getData)
		return c.String(http.StatusOK, string(getData))
	})
	//Tạo kết nối với Redis
	e.POST("/api/city/create", show)
	e.Logger.Fatal(e.Start(":1323"))
}
