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
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

type Staging struct {
	Name    string `json:"name"`
	Content string `json:"content"`
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

func showYml(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Go MySQL Tutorial")

		db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/session")

		if err != nil {
			panic(err.Error())
		}

		defer db.Close()

		results, err := db.Query("SELECT * FROM staging")
		fmt.Println(results)
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var staging Staging
			// for each row, scan the result into our staging composite object
			err = results.Scan(&staging.Name, &staging.Content)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			log.Printf(staging.Name)
		}
		// defer get.Close()
		return c.String(http.StatusOK, "Hello World!")
	})
	//Tạo kết nối với Redis
	e.GET("/api/:city", show)
	// e.GET("/users/:id", getUser)
	e.GET("/api/:app", showYml)
	e.Logger.Fatal(e.Start(":1323"))
}
