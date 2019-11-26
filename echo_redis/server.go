package main

import (
	"encoding/json"
	"fmt"
	// "github.com/go-redis/redis"
	"github.com/labstack/echo"
	"log"
	// "github.com/fatih/structs"
	"net/http"
	// "reflect"
)

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

var Data = `{"name": "hcm", "authors": [
    {"name": "duc_tri", "age": 18, "studies": [{"languages": ["eng", "zhe"]}]},
    {"name": "theanh", "age": 20, "studies": [{"languages": ["eng", "fre"]}]}
]}`

func show(c echo.Context) error {
	var city City
	json.Unmarshal([]byte(Data), &city)
	fmt.Println(city.Authors)
	pCity := &city
	fmt.Println(pCity)
	for i := 0; i < len(pCity.Authors); i++ {
		if pCity.Authors[i].Age == 18 {
			pCity.Authors[i].OnDuty = true
		}
	}
	var jsonData []byte
	jsonData, _ = json.Marshal(pCity)
	fmt.Println(string(jsonData))
	return c.String(http.StatusOK, string(jsonData))
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	//Tạo kết nối với Redis
	e.GET("/api/:city", show)
	// e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}
