package config

import (
	"fmt"
	"os"

	logg "log"

	log "github.com/sirupsen/logrus"
)

var (
	Log *log.Logger
)

func init() {
	var filename string = "src/log/logfile.log"

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	Formatter := new(log.TextFormatter)

	Formatter.FullTimestamp = true

	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	if err != nil {
		fmt.Println(err)
	} else {
		log.SetOutput(f)
	}

	log.Info(os.Stdout, logg.Llongfile, logg.LstdFlags, logg.Lshortfile)

	defer f.Close()
}
