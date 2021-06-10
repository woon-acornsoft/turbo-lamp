package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var C Config

type Config struct {
	mode string
	Db   Postgres `yaml:"postgres"`
	Log  Log      `yaml:"log"`
}

type Log struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type Postgres struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func init() {
	fmt.Println("config.go init")
	mode := os.Getenv("PROJECT_MODE")
	configPath := os.Getenv("PROJECT_CONFIG_PATH")
	filename, _ := filepath.Abs("config/config.yaml")
	if configPath != "" {
		filename, _ = filepath.Abs(configPath + "config.yaml")

	}

	yamlFile, err := ioutil.ReadFile(filename)
	fmt.Println("err", err)
	fmt.Println("yamlFile", yamlFile)
	if err == nil {
		error := yaml.Unmarshal(yamlFile, &C)
		fmt.Println("error", error)
		// fmt.Println("???", m)
	}
	C.mode = mode
}

func GetConfig() *Config {
	return &C
}

func LoadConfig() *Config {

	mode := os.Getenv("PROJECT_MODE")
	postgresYamlPath := os.Getenv("PROJECT_CONFIG_PATH")
	filename, _ := filepath.Abs("config/config.yaml")
	if postgresYamlPath != "" {
		filename, _ = filepath.Abs(postgresYamlPath + "config.yaml")

	}
	fmt.Println("file name", filename)
	// var p Postgres
	var c *Config
	yamlFile, err := ioutil.ReadFile(filename)
	fmt.Println("err", err)
	fmt.Println("yamlFile", yamlFile)
	if err == nil {
		// m := make(map[interface{}]interface{})
		error := yaml.Unmarshal(yamlFile, &c)
		fmt.Println("error", error)
		// fmt.Println("???", m)
	}
	// fmt.Println("postgres yaml >>>", p)
	// fmt.Println("postgres yaml conf >>>", conf.Db.Address)
	// fmt.Println("postgres yaml conf >> password >", conf.Db.Password)
	c.mode = mode
	fmt.Println("return config >", c.Log.Level)
	fmt.Println("return config >", c.Db.Password)
	return c
}
