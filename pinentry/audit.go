package pinentry

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Auditor struct {
	logFile string
	enableLog bool
}

func NewAuditor(logfile string, enableLog bool) (*Auditor, error) {
	if enableLog {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return &Auditor{}, err
		}
		defer f.Close()
		log.SetOutput(f)
	}
	return &Auditor{logfile, enableLog}, nil
}

func (l *Auditor) getFileWriter() io.Writer {
	f, err := os.OpenFile(l.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return f
}

func (l *Auditor) Println(v ...interface{}) {
	if l.enableLog {
		w := l.getFileWriter()
		fmt.Fprintln(w, v...)
	} else {

	}
}

func (l *Auditor) Printf(format string, v ...interface{}) {
	if l.enableLog {
		w := l.getFileWriter()
		fmt.Fprintf(w, fmt.Sprintf(format, v...))
	}
}