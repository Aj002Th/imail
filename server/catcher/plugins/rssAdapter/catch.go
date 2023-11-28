package rssAdapter

import (
	"fmt"
	"github.com/Aj002Th/imail/server/catcher"
	"github.com/mmcdole/gofeed"
	"time"
)

type Catcher struct {
	Url      string
	Category string
	Source   string
}

func NewCatcher(url, category, source string) *Catcher {
	return &Catcher{
		Url:      url,
		Category: category,
		Source:   source,
	}
}

func (c *Catcher) Catch() ([]catcher.Content, error) {
	fp := gofeed.NewParser()
	fp.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0"
	feed, err := fp.ParseURL(c.Url)
	if err != nil {
		str := err.Error()
		fmt.Println(str)
		return nil, err
	}

	result := make([]catcher.Content, 0)
	for _, item := range feed.Items {
		// 尝试获取作者名称
		authorName := ""
		authors := item.Authors
		if len(authors) > 0 {
			authorName = authors[0].Name
		}
		// 尝试获取发布时间
		publishedTime := time.Now()
		if item.PublishedParsed != nil {
			publishedTime = *item.PublishedParsed
		}
		result = append(result, catcher.Content{
			Title:       item.Title,
			Time:        publishedTime,
			Description: item.Description,
			Cover:       "",
			Link:        item.Link,
			Author:      authorName,
			Source:      c.Source,
			Category:    c.Category,
		})
	}
	return result, nil
}
