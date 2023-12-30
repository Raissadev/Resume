package jobs

import (
	w "api/jobs/weak"
	"api/src/config/logger"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron"
)

func Init() error {
	c := cron.New()

	// Compress
	err := c.AddFunc("0 0 0 * * 0", func() {
		folderPath := os.Getenv("LOG_OUTPUT")
		zipName := fmt.Sprintf("%s/BACK-%s.log.zip", os.Getenv("LOG_OUTPUT_BKP"), time.Now().Format("2006-01-02"))
		w.Compress(folderPath, zipName)
	})

	if err != nil {
		logger.Error("Error adding function to cron job:", err)
		return err
	}

	return nil
}

func InitJobs() {
	if err := Init(); err != nil {
		log.Println(err)
	}
}
