package main

import (
	"encoding/json"
	"fmt"
	// "github.com/go-redis/redis"
	"github.com/labstack/echo"
	// "log"
	"net/http"
	// "reflect"
)

type City struct {
	Name    string `json:"name"`
	Authors []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	} `json:"authors"`
}

var Data = `{"name": "hcm", "authors": [{"name": "duc_tri", "age": 18}, {"name": "theanh", "age": 18}]}`

func show(c echo.Context) error {
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// param := c.Param("city")
	// if param != "" {
	// 	valueParam, exc := client.Get("city").Result()
	// 	if exc != nil {
	// 		fmt.Println(exc)
	// 	}
	// 	fmt.Println(valueParam)
	// }

	// json, err := json.Marshal(Author{Name: "Elliot", Age: 25})
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = client.Set("author", json, 0).Err()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// val, err := client.Get("author").Result()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(val)

	// fmt.Println(c.Param("city"))

	var city City
	json.Unmarshal([]byte(Data), &city)
	fmt.Println(city.Authors)
	if city.Name == c.Param("city") {
		for _, val := range city.Authors {
			fmt.Println(val.Name)
		}

	}

	// ci := City{
	// 	Name: city["name"].(string),
	// }
	// fmt.Println(ci.Name)

	return c.String(http.StatusOK, Data)
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
