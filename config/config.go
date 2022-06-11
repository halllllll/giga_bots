package config

import (
	"log"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Service
}

type Service struct {
	Loilo     string
	L_Gate    string
	Miraiseed string
}

var Cfg ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}

	Cfg = ConfigList{
		Service: Service{
			Loilo:     cfg.Section("ServiceStat").Key("loilo").String(),
			L_Gate:    cfg.Section("ServiceStat").Key("l_gate").String(),
			Miraiseed: cfg.Section("ServiceStat").Key("miraiseed").String(),
		},
	}
}

func GetConfig(sectionName, key string) (string, error) {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return "", err
	}
	return cfg.Section(sectionName).Key(key).String(), nil
}

func UpdateConfig(sectionName, key, value string) error {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		return err
	}
	cfg.Section(sectionName).Key(key).SetValue(value)
	err = cfg.SaveTo("config.ini")
	LoadConfig()
	return err
}
