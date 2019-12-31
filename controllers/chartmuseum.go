package controllers

import (
	"os/exec"

	"chartmuseum-ui/models"

	config "chartmuseum-ui/config"

	"github.com/astaxie/beego/httplib"
	"github.com/sirupsen/logrus"

	"github.com/astaxie/beego/logs"
)

const apiGetCharts = "/api/charts"

func getCharts() map[string][]models.Chart {

	l := logs.GetLogger()

	logs.Info("Getting charts on url: %s", getBaseURL())
	res, err := httplib.Get(getBaseURL()).SetBasicAuth(config.Config.Chartmuseum.Username, config.Config.Chartmuseum.Password).Debug(true).Bytes()

	if err != nil {
		l.Panic(err.Error)
	}

	charts, err := models.NewCharts(res)
	if err != nil {
		errorRes, innerErr := models.NewError(res)
		if innerErr != nil {
			l.Panic(innerErr)
		}
		l.Panicf("Error received from ChartMuseum application: %s\n", errorRes.Message, err)
	}
	return charts
}

func uploadChart(filePath string) {

	cmd := exec.Command("curl", "-L", "--data-binary", "@"+filePath, getBaseURL())
	out, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Fatalf("cmd.Run() failed with %s\n", err)
	}
	logrus.Printf("combined out:\n%s\n", string(out))
}

func deleteChart(name string, version string) {

	logrus.Println("in deleteChart()")
	cmd := exec.Command("curl", "-X", "DELETE", getBaseURL()+"/"+name+"/"+version)
	out, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Fatalf("cmd.Run() failed with %s\n", err)
	}
	logrus.Printf("combined out:\n%s\n", string(out))
}

func getBaseURL() string {
	api := config.Config.Chartmuseum.HostAPI
	if len(api) == 0 {
		api = apiGetCharts
	}
	url := config.Config.Chartmuseum.Host + api
	return url
}
