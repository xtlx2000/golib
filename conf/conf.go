package conf

import (
	"github.com/robfig/config"
	"github.com/xtlx2000/golib/log"
)

func SetConfigFile(filename string) error {
	configFile = filename
	var err error
	settings, err = config.ReadDefault(configFile)
	if err != nil {
		log.Errorf("readConfigFileError: %v", err)
		return err
	}
	log.Infof("config file init done. configFile: %v", configFile)
	return nil
}

func GetConfig() *config.Config {
	return settings
}
