package crontab

import (
	"github.com/robfig/cron/v3"
)

func StartScheduledTasks(crontab string, cmd func()) {
	c := cron.New()
	c.AddFunc(crontab, cmd)
	c.Start()
}
