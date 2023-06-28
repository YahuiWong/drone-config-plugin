package main

import (
	"drone-config-plugin/src/plugin"
	"net/http"

	"github.com/drone/drone-go/plugin/config"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type spec struct {
	Debug                    bool   `envconfig:"PLUGIN_DEBUG"`
	Address                  string `envconfig:"PLUGIN_ADDRESS" default:":3000"`
	Secret                   string `envconfig:"PLUGIN_SECRET"`
	Token                    string `envconfig:"TOKEN"`
	ServerType               string `envconfig:"SERVERTYPE"`
	DroneConfigNamespaceTemp string `envconfig:"DRONE_CONFIG_NAMESPACE_TEMP"`
	DroneConfigRepoNameTemp  string `envconfig:"DRONE_CONFIG_REPONAME_TEMP"`
	DroneConfigBranchTemp    string `envconfig:"DRONE_CONFIG_BRANCH_TEMP" default:"master"`
	DroneConfigPathTemp      string `envconfig:"DRONE_CONFIG_YAMLPATH_TEMP" default:".drone.yml"`
}

func main() {
	spec := new(spec)
	err := envconfig.Process("", spec)
	if err != nil {
		logrus.Fatal(err)
	}

	if spec.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if spec.Secret == "" {
		logrus.Fatalln("missing secret key")
	}
	if spec.Token == "" {
		logrus.Warnln("missing token")
	}
	if spec.ServerType == "" {
		logrus.Warnln("missing servertype gitea | github")
	}
	if spec.DroneConfigNamespaceTemp == "" {
		logrus.Warnln("missing repository owner")
	}
	if spec.DroneConfigRepoNameTemp == "" {
		logrus.Warnln("missing repository name")
	}
	if spec.Address == "" {
		spec.Address = ":3000"
	}

	handler := config.Handler(
		plugin.New(
			spec.DroneConfigNamespaceTemp,
			spec.DroneConfigRepoNameTemp,
			spec.DroneConfigPathTemp,
			spec.DroneConfigBranchTemp,
			spec.ServerType,
			spec.Token,
		),
		spec.Secret,
		logrus.StandardLogger(),
	)

	logrus.Infof("server listening on address %s", spec.Address)
	logrus.Infof("debug:%s namespace:%s reponame:%s branch:%s path:%s", spec.Debug, spec.DroneConfigNamespaceTemp, spec.DroneConfigRepoNameTemp, spec.DroneConfigBranchTemp, spec.DroneConfigPathTemp)
	http.Handle("/", handler)
	logrus.Fatal(http.ListenAndServe(spec.Address, nil))
}
