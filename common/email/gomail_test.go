package email

import (
	"github.com/Aj002Th/imail/common/config"
	"testing"
)

func TestGomail(t *testing.T) {
	config.Init("../../imail.yaml")
	sender := NewGomail()
	err := sender.SendEmail("updated!", "some url")
	if err != nil {
		panic(err)
	}
}
