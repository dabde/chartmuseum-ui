package config

import (
	"flag"
	"io/ioutil"

	models "chartmuseum-ui/models"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

var configFile = flag.String("config", "./conf/config.yaml", "Config location")
var Config models.Config

func init() {
	logrus.Debug("Load Config")
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		logrus.Error(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		logrus.Fatal(err)
	}
}
