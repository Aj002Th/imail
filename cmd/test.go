package cmd

import (
	"bytes"
	"fmt"
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/common/logs"
	"github.com/Aj002Th/imail/server/catcher"
	"github.com/Aj002Th/imail/server/manager"
	"github.com/Aj002Th/imail/server/messager"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"log/slog"
	"os"
	"time"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use: "test",
	Run: testHandle,
}

func testHandle(cmd *cobra.Command, args []string) {
	config.Init(ConfigPath)
	logs.Init()
	if len(args) == 0 {
		slog.Error("args is empty")
		return
	}
	function, ok := testCmdMap[args[0]]
	if !ok {
		slog.Error("function not found")
		return
	}
	function(args[1:])
}

var testCmdMap = map[string]func([]string){
	"sendTemplate":         sendTemplateCmd,
	"sendTemplateWithData": sendTemplateWithDataCmd,
}

func sendTemplateCmd(args []string) {
	slog.Info("run cmd: imail content sendTemplateCmd")

	templatePath := "./server/manager/template/emailHTML.html"
	content, err := os.ReadFile(templatePath)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	contentStr := string(content)
	emailMessager := messager.NewEmailMessager()
	err = emailMessager.Push("test template", contentStr)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func sendTemplateWithDataCmd(args []string) {
	slog.Info("run cmd: imail test sendTemplateWithDataCmd")

	date := time.Now().Format("2006-01-02")
	data := manager.TemplateData{
		Title: date + " 订阅消息",
		Total: 3,
		Now:   date,
		Contents: map[string][]*catcher.Content{
			"source1": []*catcher.Content{
				&catcher.Content{
					Title:       "title1",
					Description: "description1",
					Cover:       "https://i0.hdslb.com/bfs/archive/4236f56d07603b29cf1cfdd1c104e761a3d7dc6a.jpg@320w_200h_1c_!web-space-upload-video.webp",
					Link:        "https://www.bilibili.com/video/BV15N411T7sc/",
					Author:      "author1",
					Source:      "source1",
					Category:    "category1",
				},
				&catcher.Content{
					Title:       "title1.1",
					Description: "description1.1",
					Cover:       "https://i0.hdslb.com/bfs/archive/4236f56d07603b29cf1cfdd1c104e761a3d7dc6a.jpg@320w_200h_1c_!web-space-upload-video.webp",
					Link:        "https://www.bilibili.com/video/BV15N411T7sc/",
					Author:      "author1.1",
					Source:      "source1.1",
					Category:    "category1.1",
				},
			},
			"source2": []*catcher.Content{
				&catcher.Content{
					Title:       "title2",
					Description: "description2",
					Cover:       "https://i0.hdslb.com/bfs/archive/4236f56d07603b29cf1cfdd1c104e761a3d7dc6a.jpg@320w_200h_1c_!web-space-upload-video.webp",
					Link:        "https://www.bilibili.com/video/BV15N411T7sc/",
					Author:      "author2",
					Source:      "source2",
					Category:    "category2",
				},
			},
		},
	}
	t, err := template.ParseFiles("./server/manager/template/emailTemplate.html")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		log.Fatal(err)
	}

	// 输出渲染结果
	result := tpl.String()
	fmt.Println(result)

	emailMessager := messager.NewEmailMessager()
	err = emailMessager.Push("test template with data", result)
}
