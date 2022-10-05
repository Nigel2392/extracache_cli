package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	extraconnector "github.com/Nigel2392/extraconnector"
)

type Config struct {
	SERVER *extraconnector.Server `json:"server"`
}

func (c *Config) LoadConfig() {
	file, err := ioutil.ReadFile("client.json")
	if err != nil {
		if os.IsNotExist(err) {
			c.SaveConfig()
			return
		} else {
			panic(err)
		}
	}
	err = json.Unmarshal(file, &c.SERVER)
	if err != nil {
		panic(err)
	}
	c.SERVER.Current_Channel = 0
}

func (c *Config) SaveConfig() error {
	c.SERVER = &extraconnector.Server{
		IP:              "127.0.0.1",
		PORT:            3239,
		Current_Channel: 0,
	}
	SERVER = c.SERVER
	file, err := json.MarshalIndent(c.SERVER, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("client.json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}
