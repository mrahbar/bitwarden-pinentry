package main

import (
	"flag"
	"github.com/Tkanos/gonfig"
	"github.com/foxcpp/go-assuan/pinentry"
	"github.com/mrahbar/bitwarden-pinentry/bitwarden"
	p "github.com/mrahbar/bitwarden-pinentry/pinentry"
	"os"
	"path"
	"path/filepath"
)

const ConfigFileName = "bitwarden-pinentry.json"
const LogFileName = "bitwarden-pinentry.log"


func main() {
	config := flag.String("config", "", "path to bitwarden-pinentry.json file")
	flag.Parse()

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	configpath := *config
	if *config == "" {
		configpath = path.Join(exPath, ConfigFileName)
	}

	configuration := bitwarden.Configuration{}
	err = gonfig.GetConf(configpath, &configuration)
	if err != nil {
		panic(err)
	}

	logPath := path.Join(exPath, LogFileName)
	if configuration.LogPath != "" {
		logPath = configuration.LogPath
	}

	auditor, err := p.NewAuditor(logPath, configuration.EnableLog)
	if err != nil {
		panic(err)
	}

	auditor.Println("Loaded bitwarden-pinentry.json config")
	startServe(configuration, auditor)
}

func startServe(conf bitwarden.Configuration, auditor *p.Auditor) {
	myClient := p.BitwardenClient{
		Session: conf.Session,
		ItemId:  conf.ItemID,
		Auditor: auditor,
	}
	callbacks := pinentry.Callbacks{
		Confirm: myClient.Confirm,
		GetPIN:  myClient.GetPIN,
		Msg:     myClient.Message,
	}

	auditor.Println("Starting sessions of bitwarden-pinentry")
	pinentry.Serve(callbacks, "Hi Ho from bitwarden-pinentry")
}
