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

	// check sample
	var sampleData map[string]interface{}
	fileNameSample := "./.env.sample.yaml"
	if _, err := os.Stat(fileNameSample); os.IsNotExist(err) {
		// intent blank
	} else {
		raw, err := ioutil.ReadFile(fileNameSample)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(raw, &sampleData)
		if err != nil {
			panic(err)
		}
	}
	var raw []byte
	var err error
	// merge with sample
	if len(sampleData) > 0 {
		var configData map[string]interface{}
		raw, err = ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(raw, &configData)
		if err != nil {
			panic(err)
		}

		mergedData := map[string]interface{}{}
		for key, value := range sampleData {
			mergedData[key] = value
		}
		for key, value := range configData {
			mergedData[key] = value
		}
		raw, err = yaml.Marshal(mergedData)
		if err != nil {
			panic(err)
		}
	} else {
		raw, err = ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
	}

	// parse to v
	err = yaml.Unmarshal(raw, v)
	if err != nil {
		panic(err)
	}
}
