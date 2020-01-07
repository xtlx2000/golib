package conf

import (
	"github.com/robfig/config"
	"github.com/xtlx2000/golib/log"
)

/* global var */
var Settings *config.Config

/* class */

/* init func */
func init() {
	Settings = GetSettings()
	log.Infof("config file init done.")
}

/* func */
func GetSettings() *config.Config {
	var settings *config.Config = nil
	var err error
	if Settings == nil {
		settings, err = config.ReadDefault("config.cfg")
		if err != nil {
			log.Errorf("config file parse error: %v\n", err)
			return nil
		}
	}
	return settings
}
