package main

import (
	"fmt"
	"github.com/labstack/echo"
	// "gopkg.in/yaml.v2"
	"bufio"
	"flag"
	"strings"
	// "io/ioutil"
	"log"
	// "io"
	"net/http"
	"os"
	// "encoding/json"
)

func countLeadingSpace(line string) int {
	i := 0
	for _, runeValue := range line {
		if runeValue == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}

func construcKeyLevel(keyLevels [10]string, currentLevel int) string {
	keyForHashedMap := ""
	for index := 0; index <= currentLevel; index++ {
		keyForHashedMap += "|" + keyLevels[index]
	}
	return keyForHashedMap
}

func query(c echo.Context) error {
	protocol := c.FormValue("protocol")
	fmt.Println(protocol)

	if protocol == "" {
		return c.String(http.StatusNotFound, "Not FOUND")
	}
	// read file yaml'
	var keyLevels [10]string
	var k map[string][]string
	k = make(map[string][]string)
	fptr := flag.String("fpath", "format.yaml", "file path")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		currentLine := s.Text()
		if currentLine != "" {
			numberOfLeadingSpace := countLeadingSpace(currentLine)
			currentLevel := numberOfLeadingSpace / 2
			currentLine := strings.TrimSpace(currentLine)
			currentKeyValue := strings.Split(currentLine, ":")
			keyLevels[currentLevel] = currentKeyValue[0]
			if currentKeyValue[1] != "" {
				// construct key
				keyForHashedMap := construcKeyLevel(keyLevels, currentLevel)
				k[keyForHashedMap] = append(k[keyForHashedMap], currentKeyValue[1])

				// append with value of parent key level
				for index := 0; index < currentLevel; index++ {
					parentLevelKeyForHashedMap := construcKeyLevel(keyLevels, index)
					if len(k[parentLevelKeyForHashedMap]) > 0 {

						for index := 0; index < len(k[parentLevelKeyForHashedMap]); index++ {
							k[keyForHashedMap] = append(k[keyForHashedMap], k[parentLevelKeyForHashedMap][index])
						}

					}
				}
			}
			// fmt.Println(keyLevels)
		}
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(k)
	fmt.Println("====")

	fmt.Println("Example")
	fmt.Println("|mobile|iphone -> ")
	for index := 0; index < len(k["|mobile|iphone"]); index++ {
		fmt.Println(k["|mobile|iphone"][index])
	}

	fmt.Println("|tv|panasonic|android_tv_os -> ")
	for index := 0; index < len(k["|tv|panasonic|android_tv_os"]); index++ {
		fmt.Println(k["|tv|panasonic|android_tv_os"][index])
	}

	return c.HTML(http.StatusOK, "Thank you! \n")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.POST("/query", query)

	e.Logger.Fatal(e.Start(":1323"))
}
