package email

import (
	"github.com/Aj002Th/imail/common/config"
	"testing"
)

func TestEmailSender(t *testing.T) {
	config.Init("../../imail.yaml")
	Init()
	err := SendEmail("updated!", "some url")
	if err != nil {
		panic(err)
	}
}
