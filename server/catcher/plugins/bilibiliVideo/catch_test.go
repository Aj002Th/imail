package bilibiliVideo

import (
	"fmt"
	"testing"
)

func Test_test(t *testing.T) {
	uid := "1525355"
	catcher := NewCatcher(uid, "开发")
	info, err := catcher.Catch()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(info)
}
