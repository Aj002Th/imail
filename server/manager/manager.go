package manager

import (
	"github.com/Aj002Th/imail/common/config"
	"github.com/Aj002Th/imail/server/catcher"
	"github.com/Aj002Th/imail/server/catcher/plugins/bilibiliVideo"
	"github.com/Aj002Th/imail/server/manager/dal/model"
	"github.com/Aj002Th/imail/server/manager/dal/query"
	"github.com/Aj002Th/imail/server/messager"
	"log/slog"
)

type Manager struct {
	Catchers  []catcher.Catcher
	Messagers []messager.Messager
}

func NewContentManager() *Manager {
	manager := &Manager{}

	// catcher init
	bilibiliVideoConfigs := config.GetBilibiliVideoConfigs()
	if len(bilibiliVideoConfigs) != 0 {
		for _, cfg := range bilibiliVideoConfigs {
			manager.Catchers = append(manager.Catchers,
				bilibiliVideo.NewCatcher(
					cfg.Uid,
					cfg.Category,
				))
		}
	}

	// messager init
	manager.Messagers = append(manager.Messagers, messager.NewEmailMessager())

	return manager
}

func (m *Manager) Run() {
	//crontab.StartScheduledTasks(config.GetCronTab(), func() {m.run()})
	m.run()
}

func (m *Manager) run() {
	contents := make([]catcher.Content, 0)

	// 获取所有爬虫爬取到的数据
	for _, c := range m.Catchers {
		contentBatch, err := c.Catch()
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		contents = append(contents, contentBatch...)
	}

	// 插入数据库, 通过一个唯一索引来进行了去重
	for _, c := range contents {
		_ = query.Content.Create(&model.Content{
			Content: catcher.Content{
				Description: c.Description,
				Cover:       c.Cover,
				Link:        c.Link,
				Author:      c.Author,
				Source:      c.Source,
				Category:    c.Category,
				Title:       c.Title,
			},
			Sended: false,
		})
	}

	// 将所有未发送的数据进行发送
	contentToSend, err := query.Content.Where(query.Content.Sended.Is(false)).Find()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	for _, m := range m.Messagers {
		// todo: 格式整理
		err := m.Push("每日订阅消息", convContentsToMessage(convContentModelsToContents(contentToSend)))
		if err != nil {
			slog.Error(err.Error())
			return
		}
	}
	_, err = query.Content.Where(query.Content.Sended.Is(false)).Update(query.Content.Sended, true)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
