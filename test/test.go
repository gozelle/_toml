package main

import (
	"encoding/json"
	"fmt"
	"github.com/gozelle/_toml"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DefaultConfig struct {
	TargetDSN string `toml:"target_dsn"`
}

type ExtraConfig struct {
	DefaultConfig
	TargetType int `toml:"target_type"`
}

func main() {
	pwd, _ := os.Getwd()
	c, err := os.Open(filepath.Join(pwd, "./test/test.toml"))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = c.Close()
	}()
	
	b, err := ioutil.ReadAll(c)
	if err != nil {
		panic(err)
	}
	defaultConfig := DefaultConfig{}
	err = _toml.Unmarshal(b, &defaultConfig)
	if err != nil {
		panic(err)
	}
	d1, _ := json.MarshalIndent(defaultConfig, "", "\t")
	fmt.Println(string(d1))
	
	extraConfig := ExtraConfig{}
	err = _toml.Unmarshal(b, &extraConfig)
	if err != nil {
		panic(err)
	}
	d2, _ := json.MarshalIndent(extraConfig, "", "\t")
	fmt.Println(string(d2))
}
