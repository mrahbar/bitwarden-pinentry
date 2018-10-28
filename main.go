package main

import (
	"fmt"
	"github.com/Tkanos/gonfig"
	"github.com/foxcpp/go-assuan/pinentry"
	"github.com/mrahbar/bitwarden-pinentry/bitwarden"
	p "github.com/mrahbar/bitwarden-pinentry/pinetry"
	"os"
	"path"
	"path/filepath"
)

const ConfigFileName = "bitwarden-pinentry.json"


func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)

	configuration := bitwarden.Configuration{}
	err = gonfig.GetConf(path.Join(exPath, ConfigFileName), &configuration)
	if err != nil {
		panic(err)
	}

	startServe(configuration)
}

func startServe(conf bitwarden.Configuration) {
	myClient := p.BitwardenClient{
		Session: conf.Session,
		ItemId:  conf.ItemID,
	}
	callbacks := pinentry.Callbacks{
		Confirm: myClient.Confirm,
		GetPIN:  myClient.GetPIN,
		Msg:     myClient.Message,
	}
	pinentry.Serve(callbacks, "Hi Ho from bitwarden-pinentry")
}
