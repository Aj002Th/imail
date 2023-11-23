package crontab

import (
	"github.com/robfig/cron/v3"
)

func StartScheduledTasks(crontab string, cmd func()) error {
	c := cron.New()
	_, err := c.AddFunc(crontab, cmd)
	if err != nil {
		return err
	}
	c.Start()
	return nil
}
