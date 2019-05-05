package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/Tkanos/gonfig"
	"github.com/alexflint/go-arg"
	"github.com/foxcpp/go-assuan/pinentry"
	"github.com/mrahbar/bitwarden-pinentry/bitwarden"
	p "github.com/mrahbar/bitwarden-pinentry/pinentry"
)

const ConfigFileName = "bitwarden-pinentry.json"
const LogFileName = "bitwarden-pinentry.log"

func main() {
	var args struct {
		Display string
		Config  string `help:"path to bitwarden-pinentry.json file"`
	}

	err := arg.Parse(&args)
	if err != nil {
		panic(err)
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	configpath := args.Config
	if args.Config == "" {
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
