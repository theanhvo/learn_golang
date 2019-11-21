package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"reflect"
)

type Yml struct {
	Actions []string
	Person  map[string]map[string]interface{}
}

func main() {
	yml := Yml{}
	yamlFile, err := ioutil.ReadFile("example.yml")
	HandleError("Read yml file error -> ", err)

	e := yaml.Unmarshal(yamlFile, &yml)

	HandleError("Unmarshal error -> ", e)

	for _, v := range yml.Person {
		for _, ps := range v {
			i := reflect.ValueOf(ps)
			fmt.Println(i)
			if i.Kind() == reflect.Map {
				for _, key := range i.MapKeys() {
					strct := i.MapIndex(key)
					fmt.Println(key.Interface(), strct.Interface())
				}
			} else if i.Kind() == reflect.Slice {
				for t := 0; t < i.Len(); t++ {
					fmt.Println(i.Index(t))
				}
			} else if i.Kind() == reflect.String {
				fmt.Println(i)
			}

		}
	}
}

func HandleError(message string, e error) {
	if e != nil {
		log.Fatal(message, e)
	}
}
