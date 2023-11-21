package manager

import (
	"github.com/Aj002Th/imail/server/catcher"
	"github.com/Aj002Th/imail/server/manager/dal/model"
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
		Description: c.Description,
		Cover:       c.Cover,
		Link:        c.Link,
		Author:      c.Author,
		Source:      c.Source,
		Category:    c.Category,
		Title:       c.Title,
	}
}

func convContentsToMessage(contents []*catcher.Content) string {
	result := ""
	for _, c := range contents {
		result += convContentToMessage(c) + "\n"
	}
	return result
}

func convContentToMessage(c *catcher.Content) string {
	return c.Title + "\n" + c.Description + "\n" + c.Cover + "\n" + c.Link + "\n" + c.Author + "\n" + c.Source + "\n" + c.Category
}
