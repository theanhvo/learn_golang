package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		dict := make(map[string]string)
		dict["name"] = "tri"
		fmt.Println(dict)
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
