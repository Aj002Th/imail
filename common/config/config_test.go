package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	Init("../../imail.yaml")
	list := GetBilibiliVideoTargets()
	fmt.Println(list)
}
