package pinetry

import (
	"log"
	"os"
)

type Auditor struct {
	Logger *log.Logger
}

func NewAuditor(logfile string) (*Auditor, error) {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &Auditor{}, err
	}
	defer f.Close()

	logger := log.New(f, "bitwarden-pinentry", log.LstdFlags)
	return &Auditor{logger}, nil
}
