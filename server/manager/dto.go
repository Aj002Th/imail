package manager

import (
	"bytes"
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/common/text"
	"github.com/Aj002Th/imail/server/catcher"
	"github.com/Aj002Th/imail/server/manager/dal/model"
	"html/template"
	"log/slog"
	"time"
)

func convContentModelsToContents(contents []*model.Content) []*catcher.Content {
	result := make([]*catcher.Content, 0)
	for _, c := range contents {
		result = append(result, convContentModelToContent(c))
	}
	return result
}

func convContentModelToContent(c *model.Content) *catcher.Content {
	return &catcher.Content{
		Title:       c.Title,
		Time:        c.Time,
		Description: text.LengthLimit(c.Description, config.GetDescriptionLengthLimit()),
		Cover:       c.Cover,
		Link:        c.Link,
		Author:      c.Author,
		Source:      c.Source,
		Category:    c.Category,
	}
}

type TemplateData struct {
	Title    string
	Total    int
	Now      string
	Contents map[string][]*catcher.Content
}

func convContentsToMessage(contents []*catcher.Content) string {
	// 1. 按照 category 进行分类
	categoryMap := make(map[string][]*catcher.Content)
	for _, c := range contents {
		categoryMap[c.Category] = append(categoryMap[c.Category], c)
	}

	// 2. 组装数据, 注入 template
	date := time.Now().Format("2006-01-02")
	data := TemplateData{
		Title:    date + " 订阅消息",
		Total:    len(contents),
		Now:      date,
		Contents: categoryMap,
	}
	t, err := template.ParseFiles("./server/manager/template/emailTemplate.html")
	if err != nil {
		slog.Error(err.Error())
		return ""
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		slog.Error(err.Error())
		return ""
	}

	// 输出渲染结果
	result := tpl.String()
	return result
}
