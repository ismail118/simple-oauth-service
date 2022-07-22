package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"simple-oauth-service/helper"
)

func NewConfig(filename string) *Config {
	config := new(Config)

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(errors.New("couldn't read config file:" + filename))
	}

	err = json.Unmarshal(data, config)
	helper.PanicIfError(err)

	return config
}

type Config struct {
	Server   Server `json:"server"`
	Database struct {
		MySql MySql `json:"mysql"`
		Redis Redis `json:"redis"`
	} `json:"database"`
	Email Email `json:"email"`
}

type Server struct {
	Addr string `json:"addr"`
}

type MySql struct {
	Addr         string `json:"addr"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type Email struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
