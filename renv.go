package renv

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var envMode = flag.String("env", "", "env mode")

func ParseCmd(v interface{}) {
	Parse(*envMode, v)
}

func Parse(env string, v interface{}) {
	var fileName string
	if env == "" {
		fileName = "./.env.local.yaml"
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fileName = "./.env.yaml"
			if _, err := os.Stat(fileName); os.IsNotExist(err) {
				fmt.Println(fileName)
				panic("missing env file")
			}
		}
	} else {
		fileName = fmt.Sprintf("./.env.%s.yaml", env)
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			fmt.Println(fileName)
			panic("missing env file")
		}
	}

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(raw, v)
	if err != nil {
		panic(err)
	}
}
