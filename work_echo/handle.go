package main

import (
	"fmt"
	"github.com/labstack/echo"
	// "gopkg.in/yaml.v2"
	"bufio"
	"flag"
	// "reflect"
	"strings"
	// "io/ioutil"
	"log"
	// "io"
	"net/http"
	"os"
	// "encoding/json"
)

func CountKeylevel(line string) int {
	i := 0
	for _, val := range line {
		if val == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}
func construcKeyLevel(keyLevel [10]string, currentLevel int) string {
	keyForHashedMap := ""
	for index := 0; index <= currentLevel; index++ {
		keyForHashedMap += "|" + keyLevel[index]
	}
	return keyForHashedMap
}

func query(c echo.Context) error {
	protocol := c.FormValue("protocol")
	fmt.Println(protocol)
	if protocol == "" {
		return c.String(http.StatusNotFound, "Not FOUND")
	}

	var keyLevel [10]string
	key := make(map[string][]string)
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
			numberOfLeadingSpace := CountKeylevel(currentLine)
			currentLevel := numberOfLeadingSpace / 2
			currentLine := strings.TrimSpace(currentLine)
			currentKeyValue := strings.Split(currentLine, ":")
			keyLevel[currentLevel] = currentKeyValue[0]

			if currentKeyValue[1] != "" {
				// construct key
				keyForHashedMap := construcKeyLevel(keyLevel, currentLevel)
				// fmt.Println(keyForHashedMap)
				key[keyForHashedMap] = append(key[keyForHashedMap], currentKeyValue[1])

				// append with value of parent key level
				for index := 0; index < currentLevel; index++ {
					parentLevelKeyForHashedMap := construcKeyLevel(keyLevel, index)
					if len(key[parentLevelKeyForHashedMap]) > 0 {

						for index := 0; index < len(key[parentLevelKeyForHashedMap]); index++ {
							key[keyForHashedMap] = append(key[keyForHashedMap], key[parentLevelKeyForHashedMap][index])
						}

					}
				}
			}

		}
	}

	for index := 0; index < len(key["|mobile|iphone"]); index++ {
		fmt.Println(key["|mobile|iphone"][index])
	}

	fmt.Println("|tv|panasonic|android_tv_os -> ")
	for index := 0; index < len(key["|tv|panasonic|android_tv_os"]); index++ {
		fmt.Println(key["|tv|panasonic|android_tv_os"][index])
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
