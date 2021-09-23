package enum

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type RespSt struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Config struct {
	Fund map[string]struct {
		Name string `yaml:"name"`
		Url  string `yaml:"url"`
	} `yaml:"fund"`
}

func (c *Config) GetConf() *Config {

	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	yamlFile, err := ioutil.ReadFile(path + "/config/conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
