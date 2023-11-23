package crontab

import (
	"fmt"
	"github.com/Aj002Th/imail/common/config"
	"testing"
	"time"
)

func TestEmailSender(t *testing.T) {
	config.Init("../../imail.yaml")
	StartScheduledTasks("*/1 * * * *", func() {
		fmt.Println("pass one minute!")
	})
	time.Sleep(time.Minute * 5)
}
