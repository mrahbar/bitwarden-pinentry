package main

import (
	"encoding/json"
	"flag"
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

	configuration := loadConfig(configpath)
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

func loadConfig(configpath string) bitwarden.Configuration {
	bytes, err := os.ReadFile(configpath)
	if err != nil {
		panic(err)
	}

	configuration := bitwarden.Configuration{}
	err = json.Unmarshal(bytes, &configuration)
	if err != nil {
		panic(err)
	}

	if session, exists := os.LookupEnv("BW_SESSION"); exists {
		configuration.Session = session
	}
	if item, exists := os.LookupEnv("BW_ITEMID"); exists {
		configuration.ItemID = item
	}

	return configuration
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
